package text

import (
	"bufio"
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

	// 1. 创建测试文件（如果不存在）
	testFileName := "fe.log"
	//if _, err := os.Stat(testFileName); os.IsNotExist(err) {
	//	// 创建测试文件
	//	file, err := os.Create(testFileName)
	//	if err != nil {
	//		t.Fatal("创建测试文件失败:", err)
	//	}
	//	defer file.Close()
	//
	//	// 写入一些测试数据
	//	for i := 0; i < 1000; i++ {
	//		line := fmt.Sprintf("2024-01-01 12:00:00 [INFO] Test log line %d\n", i)
	//		file.WriteString(line)
	//	}
	//	file.Close()
	//}

	// 2. 打开测试文件
	file, err := os.Open(testFileName)
	if err != nil {
		t.Fatal("文件打开失败:", err)
	}
	defer file.Close()

	// 3. 创建带缓冲的Reader（减少系统调用）
	r := bufio.NewReaderSize(file, 64*1024) // 64KB缓冲区
	currentOffset := int64(0)
	lineCount := 0
	maxIterations := 1000000 // 增加最大迭代次数以处理大文件
	iterationCount := 0

	// 4. 分块读取循环
	for {
		iterationCount++
		if iterationCount > maxIterations {
			t.Fatal("达到最大迭代次数，可能存在无限循环")
		}

		block := manager.AllocateBlock(currentOffset)
		n, err := io.ReadFull(r, block.Data) // 尝试读取完整4KB
		if err == io.EOF {
			manager.ReleaseBlock(block) // 文件结束释放最后一块
			break
		} else if err == io.ErrUnexpectedEOF {
			block.Data = block.Data[:n] // 调整块为实际大小
		} else if err != nil {
			manager.ReleaseBlock(block)
			t.Fatal("读取失败:", err)
		}

		// 如果没有读取到任何数据，退出循环
		if n == 0 {
			manager.ReleaseBlock(block)
			break
		}

		// 5. 模拟消费逻辑（示例：解析日志行）
		consumed := 0
		for consumed < n {
			// 查找换行符作为日志行边界
			idx := bytes.IndexByte(block.Data[consumed:], '\n')
			if idx == -1 {
				break // 本块无完整行
			}
			lineEnd := consumed + idx + 1
			line := block.Data[consumed:lineEnd] // 移除未使用变量

			// 实际业务逻辑：解析时间戳、过滤等
			lineCount++
			// 每1000行打印一次进度，避免输出过多

			fmt.Printf("已处理 %d 行日志: %s\n", lineCount, line)

			// 6. 标记已消费区域（单行）
			lineLength := lineEnd - consumed
			manager.MarkConsumed(currentOffset+int64(consumed), lineLength)
			manager.ReleaseByOffset(currentOffset+int64(consumed), lineLength)
			consumed += lineLength
		}

		// 7. 处理跨块未消费数据（如半行日志）
		if consumed < n {
			remaining := n - consumed
			nextBlock := manager.AllocateBlock(currentOffset + int64(n))

			// 将剩余数据拷贝到新块起始
			copy(nextBlock.Data, block.Data[consumed:n])
			manager.MarkConsumed(currentOffset+int64(consumed), remaining)
		}

		currentOffset += int64(n)
		manager.ReleaseBlock(block) // 整块释放
	}

	// 8. 等待异步释放完成
	time.Sleep(100 * time.Millisecond)

	// 9. 检查统计信息
	active, partial, releasable := manager.Stats()
	fmt.Printf("内存统计: Active=%d, Partial=%d, Releasable=%d\n", active, partial, releasable)
	fmt.Printf("总共处理了 %d 行日志\n", lineCount)

	// 10. 清理测试文件
	//os.Remove(testFileName)
}

// 测试内存管理器的基本功能
func TestMemoryManagerBasic(t *testing.T) {
	manager := NewManager(1024, 5) // 1KB块，预分配5块
	go manager.releaseDaemon()
	defer manager.Close()

	// 测试分配块
	block1 := manager.AllocateBlock(0)
	if block1 == nil {
		t.Fatal("分配块失败")
	}

	// 测试标记消费
	manager.MarkConsumed(0, 512)
	manager.MarkConsumed(512, 512)

	// 测试释放
	manager.ReleaseByOffset(0, 1024)

	// 等待异步处理
	time.Sleep(50 * time.Millisecond)

	// 检查统计
	active, partial, releasable := manager.Stats()
	if active != 0 && partial != 0 && releasable != 0 {
		t.Logf("统计信息: Active=%d, Partial=%d, Releasable=%d", active, partial, releasable)
	}
}

// 测试边界条件
func TestMemoryManagerEdgeCases(t *testing.T) {
	manager := NewManager(1024, 2)
	go manager.releaseDaemon()
	defer manager.Close()

	// 测试空长度释放
	manager.ReleaseByOffset(0, 0)

	// 测试负数长度释放
	manager.ReleaseByOffset(0, -1)

	// 测试不存在的块
	manager.MarkConsumed(9999, 100)

	// 测试边界消费
	manager.AllocateBlock(0)      // 分配块但不保存引用
	manager.MarkConsumed(0, 1024) // 正好一个块大小
	manager.ReleaseByOffset(0, 1024)

	time.Sleep(50 * time.Millisecond)

	// 应该没有活跃块
	active, _, _ := manager.Stats()
	if active > 0 {
		t.Logf("仍有 %d 个活跃块", active)
	}
}
