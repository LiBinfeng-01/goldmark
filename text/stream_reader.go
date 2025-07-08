package text

import (
	"bytes"
	"github.com/yuin/goldmark/util"
	"io"
	"regexp"
	"unicode/utf8"
)

type streamReader struct {
	line       int
	peekedLine []byte
	pos        Segment
	head       int
	lineOffset int
	// added for ChunkReader
	fileSize      int
	input         io.ReaderAt
	buffer        []byte
	bufferSize    int
	bufferOffset  int
	consumeOffset int
	readEOF       bool
}

// NewReader return a new Reader that can read UTF-8 bytes .
func NewStreamReader(input io.ReaderAt, fileSize int64, bufferSize int) (Reader, error) {
	buffer := make([]byte, bufferSize)
	r := &streamReader{
		fileSize:     int(fileSize),
		input:        input,
		buffer:       buffer,
		bufferOffset: 0,
		bufferSize:   bufferSize,
	}
	r.ResetPosition()
	return r, nil
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
	return r.buffer
}

func (r *streamReader) Value(seg Segment) []byte {
	result := []byte{}
	if seg.Start >= 0 && seg.Start < r.fileSize {
		result = r.buffer[seg.Start-r.bufferOffset : seg.Stop-r.bufferOffset]
	}
	return result
}

func (r *streamReader) Peek() byte {
	if r.pos.Start >= 0 && r.pos.Start < r.fileSize {
		return r.buffer[r.pos.Start-r.bufferSize]
	}
	return EOF
}

func (r *streamReader) PeekLine() ([]byte, Segment) {
	if r.pos.Start >= 0 && r.pos.Start < r.fileSize {
		if r.peekedLine == nil {
			r.peekedLine = r.buffer[r.pos.Start-r.bufferOffset : r.pos.Stop-r.bufferOffset]
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
			data := r.buffer[i-r.bufferOffset]
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
			return rune(' ')
		}
		return rune('\n')
	}
	i := r.pos.Start - 1
	for ; i >= 0; i-- {
		data := r.buffer[i-r.bufferOffset]
		if utf8.RuneStart(data) {
			break
		}
	}
	rn, _ := utf8.DecodeRune(r.buffer[i-r.bufferOffset:])
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
	for ; n > 0 && r.pos.Start < l; n-- {
		if r.pos.Padding != 0 {
			r.pos.Padding--
			continue
		}
		data := r.buffer[r.pos.Start-r.bufferSize]
		if data == '\n' {
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
	if r.pos.Start >= r.fileSize {
		return
	}

	r.lineOffset = -1
	i := -1
	if r.peekedLine != nil {
		r.pos.Start += len(r.peekedLine) - r.pos.Padding - 1
		i = 0
	}
	if i == -1 {
		i = bytes.IndexByte(r.buffer[r.pos.Start-r.bufferOffset:], '\n')
		// todo:need to deal with no change line in this block but in next block
	}
	r.peekedLine = nil
	if i != -1 {
		r.pos.Start += i
	} else {
		r.pos.Start = r.fileSize
	}
	r.pos.Padding = 0
}

var targetPuncts = []rune{'。', '.', '？', '?', '!', ';', ':', '，', ',', ' '}

// 按优先级查找最后一个标点的字节下标（从后往前扫描）
func findLastPunctuationIndex(data []byte) int {
	for i := len(data) - 1; i >= 0; {
		// 跳过非UTF-8首字节（如中文的第2/3字节）
		if !utf8.RuneStart(data[i]) {
			i--
			continue
		}

		// 解码当前字符
		r, size := utf8.DecodeRune(data[i:])
		// 按优先级匹配标点
		for _, punct := range targetPuncts {
			if r == punct {
				return i // 返回标点首字节索引
			}
		}
		i -= size // 跳至前一个字符首字节
	}
	return -1 // 未找到
}

func (r *streamReader) punctuationCuttingLine() {
	idx := findLastPunctuationIndex(r.buffer)
	if idx == -1 {
		// it may reach the end of the file
		stopIdx := r.bufferOffset + r.bufferSize
		if r.fileSize < stopIdx {
			r.pos.Stop = r.fileSize
		} else {
			r.pos.Stop = stopIdx
		}
	} else {
		r.pos.Stop = r.bufferOffset + idx + 1
	}
	r.line++
}

func (r *streamReader) AdvanceLine() {
	r.lineOffset = -1
	r.peekedLine = nil
	r.pos.Start = r.pos.Stop
	r.head = r.pos.Start
	r.pos.Padding = 0

	if r.pos.Start == 0 && r.bufferOffset == 0 {
		r.buffer = make([]byte, r.bufferSize)
		n, err := r.input.ReadAt(r.buffer, int64(r.pos.Start))
		if err == io.EOF {
			r.readEOF = true
		}
		r.buffer = r.buffer[:n]
	}
	if r.pos.Start-r.bufferOffset > len(r.buffer) {
		r.pos.Stop = r.fileSize
		r.line++
		return
	}

	idx := bytes.IndexByte(r.buffer[r.pos.Start-r.bufferOffset:], '\n')
	if idx == -1 {
		// if it has reach to the end
		if r.readEOF {
			r.pos.Stop = r.fileSize
			r.line++
			return
		}
		// it means the last time it does not consume and go back here directly,
		// so we need to force it to pop one line, but not moving to the next buffer
		if r.consumeOffset == r.bufferOffset {
			r.punctuationCuttingLine()
			return
		}
		// copy rest of old buffer to nextBuffer
		nextBuffer := make([]byte, r.bufferSize)
		lastReserveIndex := r.consumeOffset - r.bufferOffset
		copy(nextBuffer[:r.bufferSize-lastReserveIndex], r.buffer[lastReserveIndex:r.bufferSize])

		_, err := r.input.ReadAt(nextBuffer[r.bufferSize-lastReserveIndex:], int64(r.bufferSize+r.bufferOffset))
		if err == io.EOF {
			r.readEOF = true
		}
		r.bufferOffset = r.consumeOffset
		r.buffer = nextBuffer
		nextIdx := bytes.IndexByte(r.buffer[r.pos.Start-r.bufferOffset:], '\n')
		if nextIdx == -1 {
			r.punctuationCuttingLine()
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

func (r *streamReader) GetRange(start int, stop int) []byte {
	result := []byte{}
	if start >= 0 && stop <= r.fileSize {
		result = r.buffer[start-r.bufferOffset : stop-r.bufferOffset]
	}
	return result
}

func (r *streamReader) Buffer() []byte {
	return r.buffer
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
