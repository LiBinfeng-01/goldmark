package chunker

import (
	"context"
	"errors"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/text"
	"io"

	"github.com/yuin/goldmark/parser"
)

type chunkReader struct {
	input          io.ReaderAt
	fileSize       int
	chunkSize      int
	bufferSize     int
	startPushChunk bool
	chunkChan      chan *parser.Chunk
	lastChunk      *parser.Chunk
}

func NewChunkReader(input io.ReaderAt, fileSize int, chunkSize int, bufferSize int) *chunkReader {
	return &chunkReader{
		input:          input,
		fileSize:       fileSize,
		chunkSize:      chunkSize,
		bufferSize:     bufferSize,
		startPushChunk: false,
		chunkChan:      make(chan *parser.Chunk),
	}
}

var ErrChunkFinishUnexpected = errors.New("chunker finished before it consume all data")

func (c *chunkReader) pushChunksIntoChannel(ctx context.Context) {
	reader := text.NewStreamReader(c.input, int64(c.fileSize), c.bufferSize, c.chunkSize)
	root := ast.NewDocument()
	ctxForChunk := parser.NewContextForChunk(parser.Block{Node: root, Parser: nil, Length: 0, Start: 0,
		BufferWindow: nil}, c.chunkChan, c.chunkSize)
	ctxForChunk.SetCurrentLevel(1)
	ctxForChunk.SetOpenedBlocks(nil)
	markdown := goldmark.New(
		goldmark.WithExtensions(
			extension.Table,
			extension.GFM,
		),
	)
	markdown.Parser().InitParser()
	go markdown.Parser().PushChunks(ctx, reader, parser.WithContext(ctxForChunk))
}

func (c *chunkReader) NextChunk(ctx context.Context) (*parser.Chunk, error) {
	if !c.startPushChunk {
		c.startPushChunk = true
		// pushing chunks into channel.
		c.pushChunksIntoChannel(ctx)
	}
	select {
	case <-ctx.Done():
	default:
		chunk, ok := <-c.chunkChan
		if ok {
			c.lastChunk = chunk
			return chunk, nil
		}
	}
	// pointer to last chunk pop from channel was pin, if it reached to the end of file, process finish.
	if c.lastChunk.ChunkStart+c.lastChunk.Length == c.fileSize {
		return nil, io.EOF
	} else {
		// something go wrong with chunker, or outside calling to stop it.
		return nil, ErrChunkFinishUnexpected
	}
}
