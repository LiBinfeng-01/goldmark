package text

import (
	"bytes"
	"regexp"
	"unicode/utf8"

	"github.com/yuin/goldmark/util"
)

type streamReader1 struct {
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
func NewStreamReader1(buffer chan []byte) Reader {
	r := &streamReader1{
		buffer:       buffer,
		sourceLength: 0,
	}
	r.ResetPosition()
	return r
}

func (r *streamReader1) FindClosure(opener, closer byte, options FindClosureOptions) (*Segments, bool) {
	return findClosureReader(r, opener, closer, options)
}

func (r *streamReader1) ResetPosition() {
	r.line = -1
	r.head = 0
	r.lineOffset = -1
	r.AdvanceLine()
}

func (r *streamReader1) Source() []byte {
	return r.source
}

func (r *streamReader1) Value(seg Segment) []byte {
	return seg.Value(r.source)
}

func (r *streamReader1) Peek() byte {
	if r.pos.Start >= 0 && r.pos.Start < r.sourceLength {
		if r.pos.Padding != 0 {
			return space[0]
		}
		return r.source[r.pos.Start]
	}
	return EOF
}

func (r *streamReader1) PeekLine() ([]byte, Segment) {
	if r.pos.Start >= 0 && r.pos.Start < r.sourceLength {
		if r.peekedLine == nil {
			r.peekedLine = r.pos.Value(r.Source())
		}
		return r.peekedLine, r.pos
	}
	return nil, r.pos
}

// io.RuneReader interface.
func (r *streamReader1) ReadRune() (rune, int, error) {
	return readRuneReader(r)
}

func (r *streamReader1) LineOffset() int {
	if r.lineOffset < 0 {
		v := 0
		for i := r.head; i < r.pos.Start; i++ {
			if r.source[i] == '\t' {
				v += util.TabWidth(v)
			} else {
				v++
			}
		}
		r.lineOffset = v - r.pos.Padding
	}
	return r.lineOffset
}

func (r *streamReader1) PrecendingCharacter() rune {
	if r.pos.Start <= 0 {
		if r.pos.Padding != 0 {
			return rune(' ')
		}
		return rune('\n')
	}
	i := r.pos.Start - 1
	for ; i >= 0; i-- {
		if utf8.RuneStart(r.source[i]) {
			break
		}
	}
	rn, _ := utf8.DecodeRune(r.source[i:])
	return rn
}

func (r *streamReader1) Advance(n int) {
	r.lineOffset = -1
	if n < len(r.peekedLine) && r.pos.Padding == 0 {
		r.pos.Start += n
		r.peekedLine = nil
		return
	}
	r.peekedLine = nil
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

func (r *streamReader1) AdvanceAndSetPadding(n, padding int) {
	r.Advance(n)
	if padding > r.pos.Padding {
		r.SetPadding(padding)
	}
}

func (r *streamReader1) AdvanceToEOL() {
	if r.pos.Start >= r.sourceLength {
		return
	}

	r.lineOffset = -1
	i := -1
	if r.peekedLine != nil {
		r.pos.Start += len(r.peekedLine) - r.pos.Padding - 1
		if r.source[r.pos.Start] == '\n' {
			i = 0
		}
	}
	if i == -1 {
		i = bytes.IndexByte(r.source[r.pos.Start:], '\n')
	}
	r.peekedLine = nil
	if i != -1 {
		r.pos.Start += i
	} else {
		r.pos.Start = r.sourceLength
	}
	r.pos.Padding = 0
}

func (r *streamReader1) AdvanceLine() {
	newLine := <-r.buffer
	r.source = append(r.source, newLine...)
	r.sourceLength = len(r.source)
	r.lineOffset = -1
	r.peekedLine = nil
	r.pos.Start = r.pos.Stop
	r.head = r.pos.Start
	if r.pos.Start < 0 || r.pos.Start >= r.sourceLength {
		return
	}
	r.pos.Stop = r.sourceLength
	i := 0
	if r.source[r.pos.Start] != '\n' {
		i = bytes.IndexByte(r.source[r.pos.Start:], '\n')
	}
	if i != -1 {
		r.pos.Stop = r.pos.Start + i + 1
	}
	r.line++
	r.pos.Padding = 0
}

func (r *streamReader1) Position() (int, Segment) {
	return r.line, r.pos
}

func (r *streamReader1) SetPosition(line int, pos Segment) {
	r.lineOffset = -1
	r.line = line
	r.pos = pos
}

func (r *streamReader1) SetPadding(v int) {
	r.pos.Padding = v
}

func (r *streamReader1) SkipSpaces() (Segment, int, bool) {
	return skipSpacesReader(r)
}

func (r *streamReader1) SkipBlankLines() (Segment, int, bool) {
	return skipBlankLinesReader(r)
}

func (r *streamReader1) Match(reg *regexp.Regexp) bool {
	return matchReader(r, reg)
}

func (r *streamReader1) FindSubMatch(reg *regexp.Regexp) [][]byte {
	return findSubMatchReader(r, reg)
}
