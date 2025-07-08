package chunker

import (
	"context"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"io"
)

type chunkReader struct {
	ctx        context.Context
	input      io.ReaderAt
	fileSize   int
	chunkSize  int
	bufferSize int
	hasInit    bool
	chunkChan  chan *parser.Chunk
}

func NewChunkReader(ctx context.Context, input io.ReaderAt, fileSize int, chunkSize int, bufferSize int) *chunkReader {
	return &chunkReader{
		ctx:        ctx,
		input:      input,
		fileSize:   fileSize,
		chunkSize:  chunkSize,
		bufferSize: bufferSize,
		hasInit:    false,
	}
}

func (c *chunkReader) NextChunk() (*parser.Chunk, error) {
	if c.ctx.Done() != nil {
		return &parser.Chunk{}, io.EOF
	}
	if !c.hasInit {
		c.hasInit = true
		reader, _ := text.NewStreamReader(c.input, int64(c.fileSize), c.bufferSize)
		root := ast.NewDocument()
		c.chunkChan = make(chan *parser.Chunk, 0)
		ctx := parser.NewContextForChunk(parser.Block{root, nil, 0, 0, []byte{}}, c.chunkChan, c.chunkSize)
		ctx.SetCurrentLevel(1)
		ctx.SetOpenedBlocks(nil)
		markdown := goldmark.New(
			goldmark.WithExtensions(
				extension.Table,
				extension.GFM,
			),
		)
		markdown.Parser().InitParser(reader)
		go func() {
			markdown.Parser().PushChunks(reader, parser.WithContext(ctx))
		}()
	}
	chunk, ok := <-c.chunkChan
	if !ok {
		return &parser.Chunk{}, io.EOF
	}
	return chunk, nil
}
