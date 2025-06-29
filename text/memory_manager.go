package text

import "sync"

type MemoryBlock struct {
	Data     []byte          // 数据存储区
	Start    int64           // 在全局内容中的起始偏移
	Usage    int             // 已消费长度（累计）
	Status   BlockStatus     // 状态: Active/Partial/Releasable
}

type BlockStatus int

const (
	BlockActive      BlockStatus = iota // 数据未消费
	BlockPartial                        // 部分消费（需保留未消费部分）
	BlockReleasable                     // 全部消费完成，可释放
)

type MemoryManager struct {
	buffSize    int                     // 单块内存大小
	blockPool   sync.Pool               // 内存块复用池
	activeMap   map[int64]*MemoryBlock  // 当前活跃块: StartOffset → Block
	releaseChan chan int64              // 释放信号通道（触发异步释放）
	mu          sync.RWMutex            // 并发安全锁
	done        chan struct{}           // 关闭信号
}

// 初始化内存池（预分配减少运行时开销）
func NewManager(buffSize int, preAlloc int) *MemoryManager {
	pool := sync.Pool{
		New: func() interface{} {
			return &MemoryBlock{Data: make([]byte, buffSize)}
		},
	}
	// 预热内存池
	for i := 0; i < preAlloc; i++ {
		pool.Put(pool.New())
	}
	return &MemoryManager{
		buffSize:    buffSize,
		blockPool:   pool,
		activeMap:   make(map[int64]*MemoryBlock),
		releaseChan: make(chan int64, 100), // 初始化 channel 并设置缓冲区
		done:        make(chan struct{}),
	}
}

// 申请新内存块并关联起始偏移
func (m *MemoryManager) AllocateBlock(startOffset int64) *MemoryBlock {
	block := m.blockPool.Get().(*MemoryBlock)
	block.Start = startOffset
	block.Usage = 0
	block.Status = BlockActive

	m.mu.Lock()
	m.activeMap[startOffset] = block
	m.mu.Unlock()
	return block
}

// 记录消费进度
func (m *MemoryManager) MarkConsumed(offset int64, length int) {
	// 计算对应的 block start offset
	blockStart := (offset / int64(m.buffSize)) * int64(m.buffSize)
	
	m.mu.RLock()
	block, exists := m.activeMap[blockStart]
	m.mu.RUnlock()

	if !exists {
		return // 块不存在，忽略
	}

	block.Usage += length
	// 状态转换逻辑 - 修复边界条件
	if block.Usage >= m.buffSize {
		block.Status = BlockReleasable
	} else if block.Usage > 0 {
		block.Status = BlockPartial
	}
}

// 根据消费记录释放内存
func (m *MemoryManager) ReleaseByOffset(offset int64, length int) {
	if length <= 0 {
		return
	}
	
	startBlockIdx := offset / int64(m.buffSize)   // 起始块索引
	endBlockIdx := (offset + int64(length) - 1) / int64(m.buffSize) // 结束块索引，修复边界

	for idx := startBlockIdx; idx <= endBlockIdx; idx++ {
		blockStart := idx * int64(m.buffSize)
		m.mu.RLock()
		block := m.activeMap[blockStart]
		m.mu.RUnlock()

		if block == nil || block.Status != BlockReleasable {
			continue
		}
		// 安全释放（异步避免阻塞主线程）
		select {
		case m.releaseChan <- blockStart:
		default:
			// 如果 channel 满了，直接释放
			m.mu.Lock()
			if b, exists := m.activeMap[blockStart]; exists {
				delete(m.activeMap, blockStart)
				b.Data = b.Data[:0]
				m.blockPool.Put(b)
			}
			m.mu.Unlock()
		}
	}
}

// 释放整块内存（无论是否完全消费）
func (m *MemoryManager) ReleaseBlock(block *MemoryBlock) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// 1. 从活跃映射中移除
	delete(m.activeMap, block.Start)

	// 2. 重置数据并归还内存池
	block.Data = block.Data[:0]  // 清空数据引用
	m.blockPool.Put(block)       // 归还到sync.Pool复用
}

// 异步释放协程
func (m *MemoryManager) releaseDaemon() {
	for {
		select {
		case startOffset := <-m.releaseChan:
			m.mu.Lock()
			block, exists := m.activeMap[startOffset]
			delete(m.activeMap, startOffset) // 移除活跃记录
			m.mu.Unlock()

			if exists {
				block.Data = block.Data[:0] // 清空数据（防泄漏）
				m.blockPool.Put(block)      // 归还内存池
			}
		case <-m.done:
			return
		}
	}
}

// 关闭管理器
func (m *MemoryManager) Close() {
	close(m.done)
	close(m.releaseChan)
}

// 实时统计内存状态
func (m *MemoryManager) Stats() (active, partial, releasable int) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for _, block := range m.activeMap {
		switch block.Status {
		case BlockActive: active++
		case BlockPartial: partial++
		case BlockReleasable: releasable++
		}
	}
	return
}
