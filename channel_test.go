package goldmark_test

import (
	"bufio"
	"fmt"
	"github.com/yuin/goldmark/text"
	"io"
	"os"
	"testing"
)

func TestChannelDebug(t *testing.T) {
	const bufferSize = 1000                   // 缓冲区容量（队列长度）
	dataChan := make(chan []byte, bufferSize) // 带缓冲通道作为队列

	// if current channel is open, do not open it again
	go func() {
		defer close(dataChan) // 关闭通道通知消费者退出

		// 1. 打开文件
		file, err := os.Open("test_data.md")
		if err != nil {
			fmt.Println("打开文件失败:", err)
			return
		}
		defer file.Close() // 确保关闭文件句柄

		// 2. 创建带缓冲的读取器
		reader1 := bufio.NewReader(file)

		// 3. 循环读取每一行
		for {
			block, err := reader1.ReadBytes('\n') // 以换行符为分隔符
			if err != nil {
				// 4. 处理文件结束或错误
				if err == io.EOF {
					if block != nil { // 处理最后一行无换行符的情况
						fmt.Print(block)
					}
					break
				}
				fmt.Println("读取错误:", err)
				return
			}
			// 5. 处理每行内容（含结尾的\n）
			dataChan <- block // 写入队列（缓冲区满时阻塞）
			fmt.Printf("生产者写入: %s (队列长度: %d/%d)\n",
				block, len(dataChan), bufferSize)
		}
	}()

	reader := text.NewStreamReader(dataChan)

	//markdown := New(
	//	WithExtensions(
	//		extension.Table,
	//		extension.GFM,
	//	),
	//)
	//ctx := parser.NewContext()

	sizeLimit := 20
	buffer := make([]byte, 0)
	size := 0

	for { // 通道关闭后自动退出循环
		line, _ := reader.PeekLine()
		reader.AdvanceLine()
		if len(line) == 0 {
			break
		}
		if (size + len(line)) > sizeLimit {
			fmt.Printf("------------buffer--------------\n%s\n", buffer)
			buffer = make([]byte, 0)
			buffer = append(buffer, line...)
			size = len(line)
		} else {
			buffer = append(buffer, line...)
			size += len(line)
		}
		//fmt.Printf("消费者读取: %s with size %d\n", line, len(line))
		// 显式释放大对象内存（可选）
		//data = nil // 加速GC回收
		//block := markdown.Parser().NextBlock(reader, parser.WithContext(ctx))
		//t.Logf("Node kind: %s", block.Node.Kind())
	}

	fmt.Println("任务完成")
}
