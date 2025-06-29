package text

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
	"time"
)

func TestMemoryManagerWithFile(t *testing.T) {
	manager := NewManager(4096, 10) // 4KB块，预分配10块
	go manager.releaseDaemon()
	defer manager.Close() // 确保清理资源

	// 1. 自动查找日志文件路径
	testFileName := "../data/test_paragraph.md"
	// 2. 打开测试文件
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatal("文件打开失败:", err)
	}
	defer file.Close()

	// 获取文件大小
	fileInfo, err := file.Stat()
	if err != nil {
		t.Fatal("获取文件信息失败:", err)
	}
	fileSize := fileInfo.Size()
	fmt.Printf("文件大小: %d 字节\n", fileSize)

	currentOffset := int64(0)
	lineCount := 0
	maxIterations := 10000000 // 增加最大迭代次数以处理大文件
	iterationCount := 0
	const maxLineLen = 1024 * 1024 // 1MB
	const forceChunk = 1024        // 1KB

	var remainingData []byte
	dataOffset := currentOffset

	for currentOffset < fileSize {
		iterationCount++
		if iterationCount > maxIterations {
			t.Fatalf("达到最大迭代次数，可能存在无限循环，当前已处理%d行，remainingData长度=%d", lineCount, len(remainingData))
		}

		// 计算本次读取的大小
		remainingBytes := fileSize - currentOffset

		// 如果已经读取完所有数据，退出循环
		if remainingBytes <= 0 {
			break
		}

		block := manager.AllocateBlock(currentOffset)
		n, err := file.ReadAt(block.Data, currentOffset)

		// 处理读取错误
		if err != nil && err != io.EOF {
			manager.ReleaseBlock(block)
			t.Fatal("读取失败:", err)
		}

		// 如果没有读取到任何数据，退出循环
		if n == 0 {
			manager.ReleaseBlock(block)
			break
		}

		// 调整数据大小
		block.Data = block.Data[:n]

		dataToProcess := block.Data
		if len(remainingData) > 0 {
			dataToProcess = append(remainingData, dataToProcess...)
			dataOffset -= int64(len(remainingData))
			remainingData = nil
		} else {
			dataOffset = currentOffset
		}

		consumed := 0
		for consumed < len(dataToProcess) {
			idx := bytes.IndexByte(dataToProcess[consumed:], '\n')
			if idx == -1 {
				// 超长行保护
				if len(dataToProcess[consumed:]) > maxLineLen {
					line := dataToProcess[consumed : consumed+maxLineLen]
					lineCount++
					fmt.Printf("超长行分段输出 %d: %s\n", lineCount, string(line))
					manager.MarkConsumed(dataOffset+int64(consumed), maxLineLen)
					manager.ReleaseByOffset(dataOffset+int64(consumed), maxLineLen)
					consumed += maxLineLen
					continue
				}
				// 如果没有任何进展，强制消费1KB，避免死循环
				if len(dataToProcess[consumed:]) > 0 && consumed == 0 {
					chunk := dataToProcess[consumed:]
					chunkLen := len(chunk)
					if chunkLen > forceChunk {
						chunkLen = forceChunk
					}
					line := chunk[:chunkLen]
					lineCount++
					fmt.Printf("强制分段输出 %d: %s\n", lineCount, string(line))
					manager.MarkConsumed(dataOffset+int64(consumed), chunkLen)
					manager.ReleaseByOffset(dataOffset+int64(consumed), chunkLen)
					consumed += chunkLen
					continue
				}
				remainingData = dataToProcess[consumed:]
				break
			}
			lineEnd := consumed + idx + 1
			line := dataToProcess[consumed:lineEnd]
			lineCount++
			fmt.Printf("已处理 %d 行日志: %s", lineCount, string(line))
			manager.MarkConsumed(dataOffset+int64(consumed), lineEnd-consumed)
			manager.ReleaseByOffset(dataOffset+int64(consumed), lineEnd-consumed)
			consumed = lineEnd
		}
		// 如果本轮没有任何消费，强制消费1KB，避免死循环
		if consumed == 0 && len(dataToProcess) > 0 {
			chunkLen := len(dataToProcess)
			if chunkLen > forceChunk {
				chunkLen = forceChunk
			}
			line := dataToProcess[:chunkLen]
			lineCount++
			fmt.Printf("死循环保护分段输出 %d: %s\n", lineCount, string(line))
			manager.MarkConsumed(dataOffset, chunkLen)
			manager.ReleaseByOffset(dataOffset, chunkLen)
			consumed += chunkLen
			// 不break，继续下轮
		}

		// 添加调试信息
		if iterationCount%100 == 0 {
			fmt.Printf("调试: 迭代次数=%d, 已处理行数=%d, 当前块大小=%d, 已消费=%d, 剩余数据长度=%d, 当前偏移=%d, 文件大小=%d\n",
				iterationCount, lineCount, n, consumed, len(remainingData), currentOffset, fileSize)
		}

		if len(remainingData) == 0 {
			dataOffset = currentOffset + int64(n)
		}
		currentOffset += int64(n)
		manager.ReleaseBlock(block)

		// 如果已经到达文件末尾，退出循环
		if err == io.EOF {
			break
		}
	}
	if len(remainingData) > 0 {
		lineCount++
		fmt.Printf("已处理 %d 行日志: %s\n", lineCount, string(remainingData))
	}

	time.Sleep(100 * time.Millisecond)
	active, partial, releasable := manager.Stats()
	fmt.Printf("内存统计: Active=%d, Partial=%d, Releasable=%d\n", active, partial, releasable)
	fmt.Printf("总共处理了 %d 行日志\n", lineCount)
	//os.Remove(testFileName)
}
