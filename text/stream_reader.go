package text

import (
	"bytes"
	"github.com/yuin/goldmark/util"
	"io"
	"os"
	"regexp"
	"unicode/utf8"
)

type streamReader struct {
	line          int
	peekedLine    []byte
	pos           Segment
	head          int
	lineOffset    int
	currentOffset int64
	remainingData []byte
	fileSize      int64
	file          *os.File
	maxIterations int64
	iterations    int64
	manager       *MemoryManager
}

// NewReader return a new Reader that can read UTF-8 bytes .
func NewStreamReader(source []byte) (Reader, error) {
	manager := NewManager(4096, 10) // 4KB块，预分配10块
	go manager.releaseDaemon()
	//defer manager.Close() // 确保清理资源

	// 1. 自动查找日志文件路径
	testFileName := "../data/test_paragraph.md"
	// 2. 打开测试文件
	file, err := os.Open(testFileName)
	if err != nil {
		return nil, err
	}
	//defer file.Close()

	// 获取文件大小
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	fileSize := fileInfo.Size()

	currentOffset := int64(0)
	maxIterations := 10000000      // 增加最大迭代次数以处理大文件
	const maxLineLen = 1024 * 1024 // 1MB
	const forceChunk = 1024        // 1KB

	block := manager.AllocateBlock(currentOffset)
	_, err = file.ReadAt(block.Data, currentOffset)

	// 处理读取错误
	if err != nil && err != io.EOF {
		manager.ReleaseBlock(block)
		return nil, err
	}
	r := &streamReader{
		currentOffset: 0,
		remainingData: make([]byte, maxLineLen),
		fileSize:      fileSize,
		file:          file,
		maxIterations: int64(maxIterations),
		iterations:    0,
		manager:       manager,
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

// io.RuneReader interface.
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

func (r *streamReader) AdvanceLine() error {
	r.lineOffset = -1
	r.peekedLine = nil
	r.pos.Start = r.pos.Stop
	r.head = r.pos.Start
	if r.currentOffset < r.fileSize {
		// 计算本次读取的大小
		remainingBytes := r.fileSize - r.currentOffset
		// 如果已经读取完所有数据，退出循环
		if remainingBytes <= 0 {
			return io.EOF
		}
		block := r.manager.GetBlockByOffset(r.currentOffset)
		idx := bytes.IndexByte(block.Data[(int64(r.pos.Start)-block.Start):], '\n')
		if idx == -1 {
			block := r.manager.AllocateBlock(r.currentOffset)
			n, err := r.file.ReadAt(block.Data, r.currentOffset)

			// 处理读取错误
			if err != nil && err != io.EOF {
				return err
			}

			// 如果没有读取到任何数据，退出循环
			if n == 0 {
				r.manager.ReleaseBlock(block)
				return nil
			}
		}
		r.pos.Stop = r.pos.Start + idx + 1
	}

	if r.pos.Start < 0 || r.pos.Start >= r.sourceLength {
		return nil
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
	return nil
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

func (r *streamReader) ReleaseProcessedData() {
	// No-op for regular streamReader - it doesn't manage memory
}
