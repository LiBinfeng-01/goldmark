package text

import (
	"bytes"
	"container/list"
	"io"
	"regexp"
	"slices"
	"unicode/utf8"

	"github.com/yuin/goldmark/util"
)

type BufferWindow struct {
	buffer       []byte
	BufferOffset int
}

func (b *BufferWindow) BufferSize() int {
	return len(b.buffer)
}

func (b *BufferWindow) GetRange(start, stop int) []byte {
	return b.buffer[start-b.BufferOffset : stop-b.BufferOffset]
}

type streamReader struct {
	line       int
	peekedLine []byte
	pos        Segment
	head       int
	lineOffset int
	// added for ChunkReader
	fileSize      int
	input         io.ReaderAt
	bufferWindow  *BufferWindow
	historyWindow list.List
	bufferSize    int
	consumeOffset int
	chunkSize     int
	readEOF       bool
}

// NewStreamReader return a new Reader that can read UTF-8 bytes .
func NewStreamReader(input io.ReaderAt, fileSize int64, bufferSize int, chunkSize int) Reader {
	buffer := make([]byte, bufferSize)
	r := &streamReader{
		fileSize:      int(fileSize),
		input:         input,
		bufferWindow:  &BufferWindow{buffer: buffer, BufferOffset: 0},
		bufferSize:    bufferSize,
		chunkSize:     chunkSize,
		historyWindow: list.List{},
	}
	r.historyWindow.PushBack(r.bufferWindow)
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
	return r.bufferWindow.buffer
}

func (r *streamReader) Value(seg Segment) []byte {
	result := []byte{}
	if seg.Start >= 0 && seg.Start < r.fileSize {
		result = r.bufferWindow.GetRange(seg.Start, seg.Stop)
	}
	return result
}

func (r *streamReader) Peek() byte {
	if r.pos.Start >= 0 && r.pos.Start < r.fileSize {
		return r.bufferWindow.buffer[r.pos.Start-r.bufferWindow.BufferOffset]
	}
	return EOF
}

func (r *streamReader) PeekLine() ([]byte, Segment) {
	if r.pos.Start >= 0 && r.pos.Start < r.fileSize {
		if r.peekedLine == nil {
			r.peekedLine = r.bufferWindow.GetRange(r.pos.Start, r.pos.Stop)
		}
		return r.peekedLine, r.pos
	}
	return nil, r.pos
}

// io.RuneReader interface.
func (r *streamReader) ReadRune() (rune, int, error) {
	return readRuneReader(r)
}

func (r *streamReader) LineOffset() int {
	if r.lineOffset < 0 {
		v := 0
		for i := r.head; i < r.pos.Start; i++ {
			data := r.bufferWindow.buffer[i-r.bufferWindow.BufferOffset]
			if data == '\t' {
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
			return ' '
		}
		return '\n'
	}
	i := r.pos.Start - r.bufferWindow.BufferOffset - 1
	for ; i >= 0; i-- {
		data := r.bufferWindow.buffer[i]
		if utf8.RuneStart(data) {
			break
		}
	}
	rn, _ := utf8.DecodeRune(r.bufferWindow.buffer[i:])
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
	l := r.fileSize
	index := r.pos.Start - r.bufferWindow.BufferOffset
	for ; n > 0 && index < l-r.bufferWindow.BufferOffset; n-- {
		if r.pos.Padding != 0 {
			r.pos.Padding--
			continue
		}
		data := r.bufferWindow.buffer[index]
		if data == '\n' {
			r.AdvanceLine()
			continue
		}
		index++
	}
	r.pos.Start = index + r.bufferWindow.BufferOffset
}

func (r *streamReader) AdvanceAndSetPadding(n, padding int) {
	r.Advance(n)
	if padding > r.pos.Padding {
		r.SetPadding(padding)
	}
}

func (r *streamReader) AdvanceToEOL() {
	if r.pos.Start >= r.fileSize {
		return
	}

	r.lineOffset = -1
	i := -1
	if r.peekedLine != nil {
		r.pos.Start += len(r.peekedLine) - r.pos.Padding - 1
		i = 0
	}
	r.peekedLine = nil
	if i != -1 {
		r.pos.Start += i
	} else {
		r.pos.Start = r.fileSize
	}
	r.pos.Padding = 0
}

// we should prioritize punctuations.
var targetPuncts = []rune{'。', '.', '？', '?', '!', ';', ':', '，', ',', ' '}

func findLastPunctuationIndex(data []byte) int {
	for i := len(data) - 1; i >= 0; {
		r, size := utf8.DecodeLastRune(data[:i+1])
		if slices.Contains(targetPuncts, r) {
			return i
		}
		i -= size
	}
	return -1
}

func (r *streamReader) punctuationCuttingLine(slice []byte) {
	idx := findLastPunctuationIndex(slice)
	if idx == -1 {
		// it may reach the end of the file
		stopIdx := r.pos.Start + len(slice)
		r.pos.Stop = min(stopIdx, r.fileSize)
	} else {
		r.pos.Stop = r.pos.Start + idx + 1
	}
	r.line++
}

func (r *streamReader) readFromFile(buffer []byte, start int) {
	n, err := r.input.ReadAt(buffer, int64(start))
	if err == io.EOF {
		r.readEOF = true
	}
	buffer = buffer[:n]
}

func (r *streamReader) moveToNewWindowBuffer(needCopy bool) {
	nextBuffer := make([]byte, r.bufferSize)
	nextBufferWindow := &BufferWindow{}
	if !needCopy {
		r.readFromFile(nextBuffer, r.bufferSize+r.bufferWindow.BufferOffset)
		nextBufferWindow = &BufferWindow{buffer: nextBuffer, BufferOffset: r.bufferSize + r.bufferWindow.BufferOffset}
	} else {
		lastReserveIndex := r.pos.Start - r.bufferWindow.BufferOffset
		copy(nextBuffer[:r.bufferSize-lastReserveIndex], r.bufferWindow.buffer[lastReserveIndex:r.bufferSize])
		r.readFromFile(nextBuffer[r.bufferSize-lastReserveIndex:], r.bufferSize+r.bufferWindow.BufferOffset)
		nextBufferWindow = &BufferWindow{buffer: nextBuffer, BufferOffset: r.pos.Start}
	}
	r.bufferWindow = nextBufferWindow
	r.historyWindow.PushBack(nextBufferWindow)
}

func (r *streamReader) AdvanceLine() {
	r.lineOffset = -1
	r.peekedLine = nil
	r.pos.Start = r.pos.Stop
	r.head = r.pos.Start
	r.pos.Padding = 0
	if r.pos.Start < 0 || r.pos.Start >= r.fileSize {
		return
	}

	if r.pos.Start == 0 {
		r.readFromFile(r.bufferWindow.buffer, r.pos.Start)
	}
	startIdx := r.pos.Start - r.bufferWindow.BufferOffset
	if startIdx > len(r.bufferWindow.buffer) {
		if r.pos.Start > r.fileSize {
			r.pos.Stop = r.fileSize
			r.line++
			return
		}
		// make new buffer and move to next buffer
		r.moveToNewWindowBuffer(false)
	}

	bufferStopIdx := startIdx + r.chunkSize - 1
	if bufferStopIdx > len(r.bufferWindow.buffer) {
		bufferStopIdx = len(r.bufferWindow.buffer)
	}
	idx := bytes.IndexByte(r.bufferWindow.buffer[startIdx:bufferStopIdx], '\n')
	if idx == -1 {
		// it means the last time it does not consume and go back here directly,
		// so we need to force it to pop one line, but not moving to the next buffer
		if r.readEOF || bufferStopIdx-startIdx == r.chunkSize-1 {
			r.punctuationCuttingLine(r.bufferWindow.buffer[startIdx:bufferStopIdx])
			return
		}
		// copy rest of old buffer to nextBuffer and move to next buffer
		r.moveToNewWindowBuffer(true)
		nextIdx := bytes.IndexByte(r.bufferWindow.buffer, '\n')
		if nextIdx == -1 {
			r.punctuationCuttingLine(r.bufferWindow.buffer[:r.chunkSize])
			return
		}
		r.pos.Stop = r.pos.Start + nextIdx + 1
	} else {
		r.pos.Stop = r.pos.Start + idx + 1
	}
	r.line++
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

// when start and stop range cross two different window it should get different part from different window.
func (r *streamReader) GetRange(start, stop int) []byte {
	for head := r.historyWindow.Front(); head != nil; head = head.Next() {
		if window, ok := head.Value.(*BufferWindow); ok {
			if window != nil && start >= window.BufferOffset && stop <= window.BufferOffset+window.BufferSize() {
				return window.GetRange(start, stop)
			}
		}
	}
	// it means range may cross different history window.
	for head := r.historyWindow.Front(); head != nil; head = head.Next() {
		if window, ok := head.Value.(*BufferWindow); ok {
			if window != nil && start >= window.BufferOffset && start < window.BufferOffset+window.BufferSize() {
				result := make([]byte, 0, stop-start)
				middle := window.BufferOffset + window.BufferSize()
				result = append(result, r.GetRange(start, middle)...)
				result = append(result, r.GetRange(middle, stop)...)
				return result
			}
		}
	}
	return []byte{}
}

// when start and stop range cross two different window it should get different part from different window.
func (r *streamReader) ReleaseHistoryWindow(outdatedIndex int) {
	for head := r.historyWindow.Front(); head != nil; head = head.Next() {
		if window, ok := head.Value.(*BufferWindow); ok {
			if window.BufferOffset+window.BufferSize() < outdatedIndex {
				r.historyWindow.Remove(head)
			} else {
				return
			}
		}
	}
}

func (r *streamReader) BufferWindow() *BufferWindow {
	return r.bufferWindow
}

func (r *streamReader) ConsumeOffset() int {
	return r.consumeOffset
}

func (r *streamReader) SetConsumeOffset(consumeOffset int) {
	r.consumeOffset = consumeOffset
}

func (r *streamReader) HasAllConsumed() bool {
	return r.consumeOffset >= r.fileSize
}
