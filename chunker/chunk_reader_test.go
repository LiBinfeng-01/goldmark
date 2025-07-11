package chunker_test

import (
	"context"
	"errors"
	"io"
	"os"
	"runtime"
	"testing"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/chunker"
	east "github.com/yuin/goldmark/extension/ast"
	"github.com/yuin/goldmark/parser"

	"github.com/smartystreets/goconvey/convey"
)

func TestMarkdownChunker(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Skipping test on windows")
	}
	ctx := context.Background()
	convey.Convey("test markdown chunker", t, convey.FailureHalts, func() {
		convey.Convey("test chunking paragraph", func() {
			file, err := os.Open("testdata/test_paragraph.md")
			convey.So(err, convey.ShouldBeNil)
			defer func() { convey.So(file.Close(), convey.ShouldBeNil) }()
			info, err := file.Stat()
			convey.So(err, convey.ShouldBeNil)
			size := info.Size()
			chunkReader := chunker.NewChunkReader(file, int(size), 40, 4096)
			expectResult := []parser.Chunk{
				{Data: []byte("# heading1\n"), ChunkType: ast.KindHeading, Length: 11, SeqId: 1, ParentSeqId: 0, FirstSeqId: 0, ChunkStart: 0},
				{Data: []byte("aaaa\nbbbb\ncccccccccc.\ndddddddddd.\n"), ChunkType: ast.KindParagraph, Length: 34, SeqId: 2, ParentSeqId: 1, FirstSeqId: 0, ChunkStart: 11},
				{Data: []byte("## heading2\n"), ChunkType: ast.KindHeading, Length: 12, SeqId: 3, ParentSeqId: 1, FirstSeqId: 0, ChunkStart: 45},
			}

			var result []parser.Chunk
			for {
				chunk, err := chunkReader.NextChunk(ctx)
				if errors.Is(err, io.EOF) {
					break
				}
				result = append(result, *chunk)
			}
			convey.So(result, convey.ShouldEqual, expectResult)
		})
		convey.Convey("test chunking paragraph with smaller chunk size", func() {
			file, err := os.Open("testdata/test_paragraph.md")
			convey.So(err, convey.ShouldBeNil)
			defer func() { convey.So(file.Close(), convey.ShouldBeNil) }()
			info, err := file.Stat()
			convey.So(err, convey.ShouldBeNil)
			size := info.Size()
			chunkReader := chunker.NewChunkReader(file, int(size), 20, 4096)
			expectResult := []parser.Chunk{
				{Data: []byte("# heading1\n"), ChunkType: ast.KindHeading, Length: 11, SeqId: 1, ParentSeqId: 0, FirstSeqId: 0, ChunkStart: 0},
				{Data: []byte("aaaa\nbbbb\n"), ChunkType: ast.KindParagraph, Length: 10, SeqId: 2, ParentSeqId: 1, FirstSeqId: 2, ChunkStart: 11},
				{Data: []byte("cccccccccc.\n"), ChunkType: ast.KindParagraph, Length: 12, SeqId: 3, ParentSeqId: 1, FirstSeqId: 2, ChunkStart: 21},
				{Data: []byte("dddddddddd.\n"), ChunkType: ast.KindParagraph, Length: 12, SeqId: 4, ParentSeqId: 1, FirstSeqId: 2, ChunkStart: 33},
				{Data: []byte("## heading2\n"), ChunkType: ast.KindHeading, Length: 12, SeqId: 5, ParentSeqId: 1, FirstSeqId: 0, ChunkStart: 45},
			}

			var result []parser.Chunk
			for {
				chunk, err := chunkReader.NextChunk(ctx)
				if errors.Is(err, io.EOF) {
					break
				}
				result = append(result, *chunk)
			}
			convey.So(result, convey.ShouldEqual, expectResult)
		})
		convey.Convey("test chunking table", func() {
			file, err := os.Open("testdata/test_table.md")
			convey.So(err, convey.ShouldBeNil)
			defer func() { convey.So(file.Close(), convey.ShouldBeNil) }()
			info, err := file.Stat()
			convey.So(err, convey.ShouldBeNil)
			size := info.Size()
			chunkReader := chunker.NewChunkReader(file, int(size), 200, 4096)
			expectResult := []parser.Chunk{
				{Data: []byte("# heading1\n"), ChunkType: ast.KindHeading, Length: 11, SeqId: 1, ParentSeqId: 0, FirstSeqId: 0, ChunkStart: 0},
				{Data: []byte("\n| Name | Type | Required | Description |\n|------|------|----------|-------------|\n| 1    | 1    | 1        | 1           |\n| 2    | 2    | 2        | 2           |\n"),
					ChunkType: east.KindTable, Length: 165, SeqId: 2, ParentSeqId: 1, FirstSeqId: 2, ChunkStart: 11},
				{Data: []byte("| 3    | 3    | 3        | 3           |\n"), ChunkType: east.KindTable, Length: 41, SeqId: 3, ParentSeqId: 1, FirstSeqId: 2, ChunkStart: 176},
				{Data: []byte("\n| Name | Type | Required | Description |\n|------|------|----------|-------------|\n| 1    | 1    | 1        | 1           |\n| 2    | 2    | 2        | 2           |\n"),
					ChunkType: east.KindTable, Length: 165, SeqId: 4, ParentSeqId: 1, FirstSeqId: 2, ChunkStart: 217},
				{Data: []byte("## distroy line\n"), ChunkType: ast.KindHeading, Length: 16, SeqId: 5, ParentSeqId: 1, FirstSeqId: 0, ChunkStart: 382},
				{Data: []byte("| 3    | 3    | 3        | 3           |"), ChunkType: ast.KindParagraph, Length: 40, SeqId: 6, ParentSeqId: 5, FirstSeqId: 0, ChunkStart: 398},
			}

			var result []parser.Chunk
			for {
				chunk, err := chunkReader.NextChunk(ctx)
				if errors.Is(err, io.EOF) {
					break
				}
				result = append(result, *chunk)
			}
			convey.So(result, convey.ShouldEqual, expectResult)
		})
		convey.Convey("test chunking smaller table", func() {
			file, err := os.Open("testdata/test_smaller_table.md")
			convey.So(err, convey.ShouldBeNil)
			defer func() { convey.So(file.Close(), convey.ShouldBeNil) }()
			info, err := file.Stat()
			convey.So(err, convey.ShouldBeNil)
			size := info.Size()
			chunkReader := chunker.NewChunkReader(file, int(size), 100, 4096)
			expectResult := []parser.Chunk{
				{Data: []byte("# heading1\n"), ChunkType: ast.KindHeading, Length: 11, SeqId: 1, ParentSeqId: 0, FirstSeqId: 0, ChunkStart: 0},
				{Data: []byte("\n| Name | Type | Required | Description |\n|------|------|----------|-------------|\n"),
					ChunkType: east.KindTable, Length: 83, SeqId: 2, ParentSeqId: 1, FirstSeqId: 2, ChunkStart: 11},
				{Data: []byte("| 1    | 1    | 1        | 1           |"),
					ChunkType: east.KindTable, Length: 40, SeqId: 3, ParentSeqId: 1, FirstSeqId: 2, ChunkStart: 94},
			}

			var result []parser.Chunk
			for {
				chunk, err := chunkReader.NextChunk(ctx)
				if errors.Is(err, io.EOF) {
					break
				}
				result = append(result, *chunk)
			}
			convey.So(result, convey.ShouldEqual, expectResult)
		})
		convey.Convey("test chunking smaller table smaller chunk", func() {
			file, err := os.Open("testdata/test_smaller_table.md")
			convey.So(err, convey.ShouldBeNil)
			defer func() { convey.So(file.Close(), convey.ShouldBeNil) }()
			info, err := file.Stat()
			convey.So(err, convey.ShouldBeNil)
			size := info.Size()
			chunkReader := chunker.NewChunkReader(file, int(size), 45, 4096)
			expectResult := []parser.Chunk{
				{Data: []byte("# heading1\n"), ChunkType: ast.KindHeading, Length: 11, SeqId: 1, ParentSeqId: 0, FirstSeqId: 0, ChunkStart: 0},
				{Data: []byte("\n| Name | Type | Required | Description |\n"),
					ChunkType: ast.KindParagraph, Length: 42, SeqId: 2, ParentSeqId: 1, FirstSeqId: 2, ChunkStart: 11},
				{Data: []byte("|------|------|----------|-------------|\n"),
					ChunkType: ast.KindParagraph, Length: 41, SeqId: 3, ParentSeqId: 1, FirstSeqId: 2, ChunkStart: 53},
				{Data: []byte("| 1    | 1    | 1        | 1           |"),
					ChunkType: ast.KindParagraph, Length: 40, SeqId: 4, ParentSeqId: 1, FirstSeqId: 2, ChunkStart: 94},
			}

			var result []parser.Chunk
			for {
				chunk, err := chunkReader.NextChunk(ctx)
				if errors.Is(err, io.EOF) {
					break
				}
				result = append(result, *chunk)
			}
			convey.So(result, convey.ShouldEqual, expectResult)
		})
		convey.Convey("test code", func() {
			file, err := os.Open("testdata/test_code.md")
			convey.So(err, convey.ShouldBeNil)
			defer func() { convey.So(file.Close(), convey.ShouldBeNil) }()
			info, err := file.Stat()
			convey.So(err, convey.ShouldBeNil)
			size := info.Size()
			chunkReader := chunker.NewChunkReader(file, int(size), 20, 4096)
			expectResult := []parser.Chunk{
				{Data: []byte("# heading1\n"), ChunkType: ast.KindHeading, Length: 11, SeqId: 1, ParentSeqId: 0, FirstSeqId: 0, ChunkStart: 0},
				{Data: []byte("\n```go\naaaa\n"), ChunkType: ast.KindFencedCodeBlock, Length: 12, SeqId: 2, ParentSeqId: 1, FirstSeqId: 2, ChunkStart: 11},
				{Data: []byte("bbbbbbbbbb.\n"), ChunkType: ast.KindFencedCodeBlock, Length: 12, SeqId: 3, ParentSeqId: 1, FirstSeqId: 2, ChunkStart: 23},
				{Data: []byte("cccccccccc.\nddd\n```\n"), ChunkType: ast.KindFencedCodeBlock, Length: 20, SeqId: 4, ParentSeqId: 1, FirstSeqId: 2, ChunkStart: 35},
				{Data: []byte("# heading2\n"), ChunkType: ast.KindHeading, Length: 11, SeqId: 5, ParentSeqId: 0, FirstSeqId: 0, ChunkStart: 55},
			}

			var result []parser.Chunk
			for {
				chunk, err := chunkReader.NextChunk(ctx)
				if errors.Is(err, io.EOF) {
					break
				}
				result = append(result, *chunk)
			}
			convey.So(result, convey.ShouldEqual, expectResult)
		})
		convey.Convey("test list", func() {
			file, err := os.Open("testdata/test_list.md")
			convey.So(err, convey.ShouldBeNil)
			defer func() { convey.So(file.Close(), convey.ShouldBeNil) }()
			info, err := file.Stat()
			convey.So(err, convey.ShouldBeNil)
			size := info.Size()
			chunkReader := chunker.NewChunkReader(file, int(size), 20, 4096)
			expectResult := []parser.Chunk{
				{Data: []byte("- aaaa\n- bbbb\n"), ChunkType: ast.KindListItem, Length: 7, SeqId: 1, ParentSeqId: 0, FirstSeqId: 1, ChunkStart: 0},
				{Data: []byte("- cccccccccc.\n"), ChunkType: ast.KindListItem, Length: 14, SeqId: 2, ParentSeqId: 0, FirstSeqId: 1, ChunkStart: 14},
				{Data: []byte("- dddddddddd."), ChunkType: ast.KindList, Length: 13, SeqId: 3, ParentSeqId: 0, FirstSeqId: 1, ChunkStart: 28},
			}

			var result []parser.Chunk
			for {
				chunk, err := chunkReader.NextChunk(ctx)
				if errors.Is(err, io.EOF) {
					break
				}
				result = append(result, *chunk)
			}
			convey.So(result, convey.ShouldEqual, expectResult)
		})
		convey.Convey("test size out of one memory block", func() {
			file, err := os.Open("testdata/test_table.md")
			convey.So(err, convey.ShouldBeNil)
			defer func() { convey.So(file.Close(), convey.ShouldBeNil) }()
			info, err := file.Stat()
			convey.So(err, convey.ShouldBeNil)
			size := info.Size()
			chunkReader := chunker.NewChunkReader(file, int(size), 200, 300)
			expectResult := []parser.Chunk{
				{Data: []byte("# heading1\n"), ChunkType: ast.KindHeading, Length: 11, SeqId: 1, ParentSeqId: 0, FirstSeqId: 0, ChunkStart: 0},
				{Data: []byte("\n| Name | Type | Required | Description |\n|------|------|----------|-------------|\n| 1    | 1    | 1        | 1           |\n| 2    | 2    | 2        | 2           |\n"),
					ChunkType: east.KindTable, Length: 165, SeqId: 2, ParentSeqId: 1, FirstSeqId: 2, ChunkStart: 11},
				{Data: []byte("| 3    | 3    | 3        | 3           |\n"), ChunkType: east.KindTable, Length: 41, SeqId: 3, ParentSeqId: 1, FirstSeqId: 2, ChunkStart: 176},
				{Data: []byte("\n| Name | Type | Required | Description |\n|------|------|----------|-------------|\n| 1    | 1    | 1        | 1           |\n| 2    | 2    | 2        | 2           |\n"),
					ChunkType: east.KindTable, Length: 165, SeqId: 4, ParentSeqId: 1, FirstSeqId: 2, ChunkStart: 217},
				{Data: []byte("## distroy line\n"), ChunkType: ast.KindHeading, Length: 16, SeqId: 5, ParentSeqId: 1, FirstSeqId: 0, ChunkStart: 382},
				{Data: []byte("| 3    | 3    | 3        | 3           |"), ChunkType: ast.KindParagraph, Length: 40, SeqId: 6, ParentSeqId: 5, FirstSeqId: 0, ChunkStart: 398},
			}

			var result []parser.Chunk
			for {
				chunk, err := chunkReader.NextChunk(ctx)
				if errors.Is(err, io.EOF) {
					break
				}
				result = append(result, *chunk)
			}
			convey.So(result, convey.ShouldEqual, expectResult)
		})
		convey.Convey("test punctuation", func() {
			file, err := os.Open("testdata/test_punctuation.md")
			convey.So(err, convey.ShouldBeNil)
			defer func() { convey.So(file.Close(), convey.ShouldBeNil) }()
			info, err := file.Stat()
			convey.So(err, convey.ShouldBeNil)
			size := info.Size()
			chunkReader := chunker.NewChunkReader(file, int(size), 10, 20)
			expectResult := []parser.Chunk{
				{Data: []byte("12345,"), ChunkType: ast.KindParagraph, Length: 6, SeqId: 1, ParentSeqId: 0, FirstSeqId: 1, ChunkStart: 0},
				{Data: []byte("678.9123!"), ChunkType: ast.KindParagraph, Length: 9, SeqId: 2, ParentSeqId: 0, FirstSeqId: 1, ChunkStart: 6},
				{Data: []byte("456789:"), ChunkType: ast.KindParagraph, Length: 7, SeqId: 3, ParentSeqId: 0, FirstSeqId: 1, ChunkStart: 15},
				{Data: []byte("1234?"), ChunkType: ast.KindParagraph, Length: 5, SeqId: 4, ParentSeqId: 0, FirstSeqId: 1, ChunkStart: 22},
				{Data: []byte("56789;"), ChunkType: ast.KindParagraph, Length: 6, SeqId: 5, ParentSeqId: 0, FirstSeqId: 1, ChunkStart: 27},
				{Data: []byte("123456789"), ChunkType: ast.KindParagraph, Length: 9, SeqId: 6, ParentSeqId: 0, FirstSeqId: 1, ChunkStart: 33},
				{Data: []byte("1234"), ChunkType: ast.KindParagraph, Length: 4, SeqId: 7, ParentSeqId: 0, FirstSeqId: 1, ChunkStart: 42},
			}

			var result []parser.Chunk
			for {
				chunk, err := chunkReader.NextChunk(ctx)
				if errors.Is(err, io.EOF) {
					break
				}
				result = append(result, *chunk)
			}
			convey.So(result, convey.ShouldEqual, expectResult)
		})
	})
}
