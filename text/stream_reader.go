package text

import (
	"bytes"
	"regexp"
	"unicode/utf8"

	"github.com/yuin/goldmark/util"
)

// A StreamReader is a Reader that can read from a buffer source while maintaining position relationships.
type StreamReader interface {
	Reader
	// Append appends new data to the buffer source.
	Append(data []byte)
	// Clear clears the buffer source.
	Clear()
	// HasIncompleteLine returns true if the last line in the buffer is incomplete.
	HasIncompleteLine() bool
	// Consume removes data from the buffer between start and end positions.
	// Returns the number of bytes consumed.
	Consume(start, end int) int
	// BufferSize returns the current size of the buffer.
	BufferSize() int
	// GetBuffer returns the current buffer content.
	GetBuffer() []byte
}

type streamReader struct {
	source       []byte
	sourceLength int
	line         int
	peekedLine   []byte
	pos          Segment
	head         int
	lineOffset   int
	// incompleteLine tracks if the last line in the buffer is incomplete
	incompleteLine bool
	// consumedBytes tracks the total number of bytes consumed
	consumedBytes int
}

// NewStreamReader returns a new StreamReader.
func NewStreamReader() StreamReader {
	r := &streamReader{
		source:       make([]byte, 0),
		sourceLength: 0,
	}
	r.ResetPosition()
	return r
}

func (r *streamReader) Append(data []byte) {
	// If we have an incomplete line, we need to handle it specially
	if r.incompleteLine && len(r.source) > 0 {
		// Find the last newline in the existing source
		lastNewline := bytes.LastIndexByte(r.source, '\n')
		if lastNewline == -1 {
			// No newline found, the entire existing source is part of the incomplete line
			r.source = append(r.source, data...)
		} else {
			// We have a complete line, append normally
			r.source = append(r.source, data...)
		}
	} else {
		// No incomplete line, just append
		r.source = append(r.source, data...)
	}
	r.sourceLength = len(r.source)

	// Check if the last line is complete
	r.incompleteLine = !bytes.HasSuffix(r.source, []byte("\n"))
}

func (r *streamReader) Clear() {
	r.source = make([]byte, 0)
	r.sourceLength = 0
	r.incompleteLine = false
	r.consumedBytes = 0
	r.ResetPosition()
}

func (r *streamReader) HasIncompleteLine() bool {
	return r.incompleteLine
}

func (r *streamReader) Consume(start, end int) int {
	if start < 0 || end <= start || end > r.sourceLength {
		return 0
	}

	// Calculate how many bytes to consume
	bytesToConsume := end - start

	// If we have an incomplete line, we can't consume past the last newline
	if r.incompleteLine {
		lastNewline := bytes.LastIndexByte(r.source[start:end], '\n')
		if lastNewline != -1 {
			bytesToConsume = lastNewline + 1
		} else {
			// No newline found, can't consume anything
			return 0
		}
	}

	// Remove consumed bytes from the buffer
	r.source = append(r.source[:start], r.source[start+bytesToConsume:]...)
	r.sourceLength = len(r.source)

	// Adjust positions
	if r.pos.Start > start {
		r.pos.Start -= bytesToConsume
	}
	if r.pos.Stop > start {
		r.pos.Stop -= bytesToConsume
	}
	if r.head > start {
		r.head -= bytesToConsume
	}

	// Update consumed bytes counter
	r.consumedBytes += bytesToConsume

	return bytesToConsume
}

func (r *streamReader) BufferSize() int {
	return r.sourceLength
}

func (r *streamReader) GetBuffer() []byte {
	return r.source
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

func (r *streamReader) PrecendingCharacter() rune {
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

func (r *streamReader) Advance(n int) {
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

func (r *streamReader) AdvanceLine() {
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
	} else if !r.incompleteLine {
		// If we don't have an incomplete line, treat the rest as a complete line
		r.pos.Stop = r.sourceLength
	}
	r.line++
	r.pos.Padding = 0
}

func (r *streamReader) Position() (int, Segment) {
	// Adjust the segment to account for consumed bytes
	adjustedPos := r.pos
	adjustedPos.Start += r.consumedBytes
	adjustedPos.Stop += r.consumedBytes
	return r.line, adjustedPos
}

func (r *streamReader) SetPosition(line int, pos Segment) {
	r.lineOffset = -1
	r.line = line
	// Adjust the position to account for consumed bytes
	r.pos = pos
	r.pos.Start -= r.consumedBytes
	r.pos.Stop -= r.consumedBytes
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
