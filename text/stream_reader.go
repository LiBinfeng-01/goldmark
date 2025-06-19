package text

import (
	"regexp"
	"unicode/utf8"

	"github.com/yuin/goldmark/util"
)

type streamReader struct {
	buffer       chan []byte
	source       []byte
	sourceLength int
	line         int
	peekedLine   []byte
	pos          Segment
	head         int
	lineOffset   int
}

// NewStreamReader returns a new StreamReader.
func NewStreamReader(buffer chan []byte) Reader {
	r := &streamReader{
		buffer:       buffer,
		sourceLength: 0,
	}
	r.ResetPosition()
	return r
}

func (r *streamReader) FindClosure(opener, closer byte, options FindClosureOptions) (*Segments, bool) {
	return findClosureReader(r, opener, closer, options)
}

func (r *streamReader) ResetPosition() {
	r.line = -1
	r.head = 0
	r.lineOffset = -1
	r.AdvanceLine()
}

func (r *streamReader) Source() []byte {
	return r.source
}

func (r *streamReader) Value(seg Segment) []byte {
	return seg.Value(r.source)
}

func (r *streamReader) Peek() byte {
	if r.pos.Start >= 0 && r.pos.Start < r.sourceLength {
		if r.pos.Padding != 0 {
			return space[0]
		}
		return r.source[r.pos.Start]
	}
	return EOF
}

func (r *streamReader) PeekLine() ([]byte, Segment) {
	if r.pos.Start >= 0 && r.pos.Start < r.sourceLength {
		if r.peekedLine == nil {
			r.peekedLine = r.pos.Value(r.Source())
		}
		return r.peekedLine, r.pos
	}
	return nil, r.pos
}

func (r *streamReader) ReadRune() (rune, int, error) {
	return readRuneReader(r)
}

func (r *streamReader) LineOffset() int {
	if r.lineOffset < 0 {
		v := 0
		for i := r.head; i < r.pos.Start; i++ {
			if r.peekedLine[i] == '\t' {
				v += util.TabWidth(v)
			} else {
				v++
			}
		}
		r.lineOffset = v - r.pos.Padding
	}
	return r.lineOffset
}

func (r *streamReader) PrecendingCharacter() rune {
	if r.pos.Start <= 0 {
		if r.pos.Padding != 0 {
			return rune(' ')
		}
		return rune('\n')
	}
	i := r.pos.Start - 1
	for ; i >= 0; i-- {
		if utf8.RuneStart(r.peekedLine[i]) {
			break
		}
	}
	rn, _ := utf8.DecodeRune(r.peekedLine[i:])
	return rn
}

func (r *streamReader) Advance(n int) {
	r.lineOffset = -1
	if n < len(r.peekedLine) && r.pos.Padding == 0 {
		r.pos.Start += n
		return
	}
	l := r.sourceLength
	for ; n > 0 && r.pos.Start < l; n-- {
		if r.pos.Padding != 0 {
			r.pos.Padding--
			continue
		}
		if r.source[r.pos.Start] == '\n' {
			r.AdvanceLine()
			continue
		}
		r.pos.Start++
	}
}

func (r *streamReader) AdvanceAndSetPadding(n, padding int) {
	r.Advance(n)
	if padding > r.pos.Padding {
		r.SetPadding(padding)
	}
}

func (r *streamReader) AdvanceToEOL() {
	if r.pos.Start >= r.sourceLength {
		return
	}

	r.lineOffset = -1
	r.pos.Start = r.sourceLength
	r.pos.Padding = 0
}

func (r *streamReader) AdvanceLine() {
	r.lineOffset = -1
	r.peekedLine = <-r.buffer
	//fmt.Printf("消费者读取: %s with size %d\n", r.peekedLine, len(r.peekedLine))
	r.line++
	r.head = 0
	r.pos.Padding = 0
	r.pos.Start = 0
	r.pos.Stop = len(r.peekedLine)
	r.source = r.peekedLine
	r.sourceLength = len(r.peekedLine)
}

func (r *streamReader) Position() (int, Segment) {
	return r.line, r.pos
}

func (r *streamReader) SetPosition(line int, pos Segment) {
	r.lineOffset = -1
	r.line = line
	r.pos = pos
}

func (r *streamReader) SetPadding(v int) {
	r.pos.Padding = v
}

func (r *streamReader) SkipSpaces() (Segment, int, bool) {
	return skipSpacesReader(r)
}

func (r *streamReader) SkipBlankLines() (Segment, int, bool) {
	return skipBlankLinesReader(r)
}

func (r *streamReader) Match(reg *regexp.Regexp) bool {
	return matchReader(r, reg)
}

func (r *streamReader) FindSubMatch(reg *regexp.Regexp) [][]byte {
	return findSubMatchReader(r, reg)
}
