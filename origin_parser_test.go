package goldmark_test

import (
	"fmt"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"io"
	"os"
	"testing"
)

func TestOriginalParser(t *testing.T) {
	fileName := "data/test_list.md"
	content, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	reader := text.NewReader(content)
	root := ast.NewDocument()
	blockChan := make(chan parser.Block, 10)
	ctx := parser.NewContextForChunk(parser.Block{root, nil, 0, -1, 0, -1}, blockChan, 15)
	ctx.Level2Node()[0] = root
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
		defer close(blockChan) // 关闭通道通知消费者退出
		markdown.Parser().PushChunks(reader, parser.WithContext(ctx))
	}()
	file, _ := os.Open(fileName)
	defer file.Close()

	for {
		block := <-blockChan
		if block.Node == nil {
			break
		}
		file.Seek(int64(block.Start), io.SeekStart)
		buf := make([]byte, block.Length)
		file.Read(buf)
		fmt.Printf("start:%d size:%d seqid:%d firstseqid:%d parent: %d kind:%s content:%s \n",
			block.Start, block.Length, block.SeqId, block.FirstSeqId, ctx.SeqID2Parent()[block.SeqId], block.Node.Kind(), buf)
	}
}
