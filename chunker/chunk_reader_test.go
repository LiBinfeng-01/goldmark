package chunker_test

import (
	"context"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/chunker"
	east "github.com/yuin/goldmark/extension/ast"
	"github.com/yuin/goldmark/parser"
	"io"
	"os"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestMarkdownChunker(t *testing.T) {
	ctx := context.Background()
	convey.Convey("test markdown chunker", t, convey.FailureHalts, func() {
		convey.Convey("test chunking paragraph", func() {
			file, err := os.Open("testdata/test_paragraph.md")
			convey.So(err, convey.ShouldBeNil)
			defer file.Close()
			info, err := file.Stat()
			convey.So(err, convey.ShouldBeNil)
			size := info.Size()
			chunkReader := chunker.NewChunkReader(ctx, file, int(size), 400, 4096)
			expectResult := []parser.Chunk{
				{[]byte("# heading1\n"), ast.KindHeading, 11, 1, 0, 0, 0},
				{[]byte("aaaa\nbbbb\ncccccccccc.\ndddddddddd.\n"), ast.KindParagraph, 34, 2, 1, 0, 11},
				{[]byte("## heading2\n"), ast.KindHeading, 12, 3, 1, 0, 45},
			}

			var result []parser.Chunk
			for {
				chunk, err := chunkReader.NextChunk()
				if err == io.EOF {
					break
				}
				result = append(result, chunk)
			}
			convey.So(result, convey.ShouldEqual, expectResult)
		})
		convey.Convey("test chunking paragraph with smaller chunk size", func() {
			file, err := os.Open("testdata/test_paragraph.md")
			convey.So(err, convey.ShouldBeNil)
			defer file.Close()
			info, err := file.Stat()
			convey.So(err, convey.ShouldBeNil)
			size := info.Size()
			chunkReader := chunker.NewChunkReader(ctx, file, int(size), 20, 4096)
			expectResult := []parser.Chunk{
				{[]byte("# heading1\n"), ast.KindHeading, 11, 1, 0, 0, 0},
				{[]byte("aaaa\nbbbb\n"), ast.KindParagraph, 10, 2, 1, 2, 11},
				{[]byte("cccccccccc.\n"), ast.KindParagraph, 12, 3, 1, 2, 21},
				{[]byte("dddddddddd.\n"), ast.KindParagraph, 12, 4, 1, 2, 33},
				{[]byte("## heading2\n"), ast.KindHeading, 12, 5, 1, 0, 45},
			}

			var result []parser.Chunk
			for {
				chunk, err := chunkReader.NextChunk()
				if err == io.EOF {
					break
				}
				result = append(result, chunk)
			}
			convey.So(result, convey.ShouldEqual, expectResult)
		})
		convey.Convey("test chunking table", func() {
			file, err := os.Open("testdata/test_table.md")
			convey.So(err, convey.ShouldBeNil)
			defer file.Close()
			info, err := file.Stat()
			convey.So(err, convey.ShouldBeNil)
			size := info.Size()
			chunkReader := chunker.NewChunkReader(ctx, file, int(size), 200, 4096)
			expectResult := []parser.Chunk{
				{[]byte("# heading1\n"), ast.KindHeading, 11, 1, 0, 0, 0},
				{[]byte("\n| Name | Type | Required | Description |\n|------|------|----------|-------------|\n| 1    | 1    | 1        | 1           |\n| 2    | 2    | 2        | 2           |\n"),
					east.KindTable, 165, 2, 1, 2, 11},
				{[]byte("| 3    | 3    | 3        | 3           |\n"), east.KindTable, 41, 3, 1, 2, 176},
				{[]byte("\n| Name | Type | Required | Description |\n|------|------|----------|-------------|\n| 1    | 1    | 1        | 1           |\n| 2    | 2    | 2        | 2           |\n"),
					east.KindTable, 165, 4, 1, 2, 217},
				{[]byte("## distroy line\n"), ast.KindHeading, 16, 5, 1, 0, 382},
				{[]byte("| 3    | 3    | 3        | 3           |"), ast.KindParagraph, 40, 6, 5, 0, 398},
			}

			var result []parser.Chunk
			for {
				chunk, err := chunkReader.NextChunk()
				if err == io.EOF {
					break
				}
				result = append(result, chunk)
			}
			convey.So(result, convey.ShouldEqual, expectResult)
		})
		convey.Convey("test chunking smaller table", func() {
			file, err := os.Open("testdata/test_smaller_table.md")
			convey.So(err, convey.ShouldBeNil)
			defer file.Close()
			info, err := file.Stat()
			convey.So(err, convey.ShouldBeNil)
			size := info.Size()
			chunkReader := chunker.NewChunkReader(ctx, file, int(size), 100, 4096)
			expectResult := []parser.Chunk{
				{[]byte("# heading1\n"), ast.KindHeading, 11, 1, 0, 0, 0},
				{[]byte("\n| Name | Type | Required | Description |\n|------|------|----------|-------------|\n"),
					east.KindTable, 83, 2, 1, 2, 11},
				{[]byte("| 1    | 1    | 1        | 1           |"),
					east.KindTable, 40, 3, 1, 2, 94},
			}

			var result []parser.Chunk
			for {
				chunk, err := chunkReader.NextChunk()
				if err == io.EOF {
					break
				}
				result = append(result, chunk)
			}
			convey.So(result, convey.ShouldEqual, expectResult)
		})
		convey.Convey("test chunking smaller table smaller chunk", func() {
			file, err := os.Open("testdata/test_smaller_table.md")
			convey.So(err, convey.ShouldBeNil)
			defer file.Close()
			info, err := file.Stat()
			convey.So(err, convey.ShouldBeNil)
			size := info.Size()
			chunkReader := chunker.NewChunkReader(ctx, file, int(size), 80, 4096)
			expectResult := []parser.Chunk{
				{[]byte("# heading1\n"), ast.KindHeading, 11, 1, 0, 0, 0},
				{[]byte("\n| Name | Type | Required | Description |\n"),
					ast.KindParagraph, 42, 2, 1, 2, 11},
				{[]byte("|------|------|----------|-------------|\n"),
					ast.KindParagraph, 41, 3, 1, 2, 53},
				{[]byte("| 1    | 1    | 1        | 1           |"),
					ast.KindParagraph, 40, 4, 1, 2, 94},
			}

			var result []parser.Chunk
			for {
				chunk, err := chunkReader.NextChunk()
				if err == io.EOF {
					break
				}
				result = append(result, chunk)
			}
			convey.So(result, convey.ShouldEqual, expectResult)
		})
		convey.Convey("test code", func() {
			file, err := os.Open("testdata/test_code.md")
			convey.So(err, convey.ShouldBeNil)
			defer file.Close()
			info, err := file.Stat()
			convey.So(err, convey.ShouldBeNil)
			size := info.Size()
			chunkReader := chunker.NewChunkReader(ctx, file, int(size), 20, 4096)
			expectResult := []parser.Chunk{
				{[]byte("# heading1\n"), ast.KindHeading, 11, 1, 0, 0, 0},
				{[]byte("\n```go\naaaa\n"), ast.KindFencedCodeBlock, 12, 2, 1, 2, 11},
				{[]byte("bbbbbbbbbb.\n"), ast.KindFencedCodeBlock, 12, 3, 1, 2, 23},
				{[]byte("cccccccccc.\nddd\n```\n"), ast.KindFencedCodeBlock, 20, 4, 1, 2, 35},
				{[]byte("# heading2\n"), ast.KindHeading, 11, 5, 0, 0, 55},
			}

			var result []parser.Chunk
			for {
				chunk, err := chunkReader.NextChunk()
				if err == io.EOF {
					break
				}
				result = append(result, chunk)
			}
			convey.So(result, convey.ShouldEqual, expectResult)
		})
		convey.Convey("test list", func() {
			file, err := os.Open("testdata/test_list.md")
			convey.So(err, convey.ShouldBeNil)
			defer file.Close()
			info, err := file.Stat()
			convey.So(err, convey.ShouldBeNil)
			size := info.Size()
			chunkReader := chunker.NewChunkReader(ctx, file, int(size), 20, 4096)
			expectResult := []parser.Chunk{
				{[]byte("- aaaa\n- bbbb\n"), ast.KindListItem, 7, 1, 0, 1, 0},
				{[]byte("- cccccccccc.\n"), ast.KindListItem, 14, 2, 0, 1, 14},
				{[]byte("- dddddddddd."), ast.KindList, 13, 3, 0, 1, 28},
			}

			var result []parser.Chunk
			for {
				chunk, err := chunkReader.NextChunk()
				if err == io.EOF {
					break
				}
				result = append(result, chunk)
			}
			convey.So(result, convey.ShouldEqual, expectResult)
		})
		convey.Convey("test size out of one memory block", func() {
			file, err := os.Open("testdata/test_table.md")
			convey.So(err, convey.ShouldBeNil)
			defer file.Close()
			info, err := file.Stat()
			convey.So(err, convey.ShouldBeNil)
			size := info.Size()
			chunkReader := chunker.NewChunkReader(ctx, file, int(size), 200, 300)
			expectResult := []parser.Chunk{
				{[]byte("# heading1\n"), ast.KindHeading, 11, 1, 0, 0, 0},
				{[]byte("\n| Name | Type | Required | Description |\n|------|------|----------|-------------|\n| 1    | 1    | 1        | 1           |\n| 2    | 2    | 2        | 2           |\n"),
					east.KindTable, 165, 2, 1, 2, 11},
				{[]byte("| 3    | 3    | 3        | 3           |\n"), east.KindTable, 41, 3, 1, 2, 176},
				{[]byte("\n| Name | Type | Required | Description |\n|------|------|----------|-------------|\n| 1    | 1    | 1        | 1           |\n| 2    | 2    | 2        | 2           |\n"),
					east.KindTable, 165, 4, 1, 2, 217},
				{[]byte("## distroy line\n"), ast.KindHeading, 16, 5, 1, 0, 382},
				{[]byte("| 3    | 3    | 3        | 3           |"), ast.KindParagraph, 40, 6, 5, 0, 398},
			}

			var result []parser.Chunk
			for {
				chunk, err := chunkReader.NextChunk()
				if err == io.EOF {
					break
				}
				result = append(result, chunk)
			}
			convey.So(result, convey.ShouldEqual, expectResult)
		})
	})
}
