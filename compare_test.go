package goldmark_test

import (
	"fmt"
	"testing"

	. "github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/text"
)

func TestCompareDebug(t *testing.T) {
	markdownContent := `
	# 1111
	## 2222
	3333.
	
	### 4444
	
	| 表头1 | 表头2 | 表头3 |
	|---|---|---|
	| 值1 | 值2 | 值3 |
	| 值4 | 值5 | 值6 |
	
	- 5555
	- 6666
	
	## 7777
	> 8888
	
	~~~go
	9999
	~~~
	`
	source := []byte(markdownContent)
	reader := text.NewReader(source)
	//const bufferSize = 1000                   // 缓冲区容量（队列长度）
	//dataChan := make(chan []byte, bufferSize) // 带缓冲通道作为队列
	//
	//// if current channel is open, do not open it again
	//go func() {
	//	defer close(dataChan) // 关闭通道通知消费者退出
	//
	//	// 1. 打开文件
	//	file, err := os.Open("test_data.md")
	//	if err != nil {
	//		fmt.Println("打开文件失败:", err)
	//		return
	//	}
	//	defer file.Close() // 确保关闭文件句柄
	//
	//	// 2. 创建带缓冲的读取器
	//	reader1 := bufio.NewReader(file)
	//
	//	// 3. 循环读取每一行
	//	for {
	//		block, err := reader1.ReadBytes('\n') // 以换行符为分隔符
	//		if err != nil {
	//			// 4. 处理文件结束或错误
	//			if err == io.EOF {
	//				if block != nil { // 处理最后一行无换行符的情况
	//					fmt.Print(block)
	//				}
	//				break
	//			}
	//			fmt.Println("读取错误:", err)
	//			return
	//		}
	//		// 5. 处理每行内容（含结尾的\n）
	//		dataChan <- block // 写入队列（缓冲区满时阻塞）
	//		fmt.Printf("生产者写入: %s (队列长度: %d/%d)\n",
	//			block, len(dataChan), bufferSize)
	//	}
	//}()
	//
	//reader := text.NewStreamReader(dataChan)
	//ctx := parser.NewContext()
	//for {
	//	block := markdown.Parser().NextBlock(reader, parser.WithContext(ctx))
	//	block.Node.Dump(source, 0) // 直接输出或进行其他处理
	//}
	markdown := New(
		WithExtensions(
			extension.Table,
			extension.GFM,
		),
	)

	doc := markdown.Parser().Parse(reader)
	fmt.Println(doc)

	//doc.Dump(source, 0)
	//
	//ast.Walk(doc, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
	//	if !entering {
	//		return ast.WalkContinue, nil
	//	}
	//	t.Logf("Node kind: %s", n.Kind())
	//	if n.Kind() == ast.KindText {
	//		t.Logf("  Text: %s", n.Text(source))
	//	}
	//	return ast.WalkContinue, nil
	//})
}
