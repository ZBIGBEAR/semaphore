/*
自己实现一个信号量。其实信号量就是一个最大值控制量。初始化以恶最大值，每次从中取值，如果取到了则说明得到了许可，
可进行下一步操作，操作完成释放许可，供其他协程使用。
信号量的作用就是控制最大并发数
*/

package semaphore

import "time"

var (
	DefaultVal = struct {}{}
)

type Semaphore interface {
	Acquire()
	Release()
	TryAcquire() bool
	TryAcquireWithTime(timeout time.Duration) bool
}

func NewSemaphore(n int) Semaphore {
	return &semaphore{
		permits:n,
		chans:make(chan struct{}, n),
	}
}

type semaphore struct {
	permits int // 信号量的总数
	chans chan struct{}
}

func (s *semaphore) Acquire() {
	s.chans <- DefaultVal
}

func (s *semaphore) Release() {
	<- s.chans
}

func (s *semaphore) TryAcquire() bool {
	select{
	case s.chans <- DefaultVal:
		return true
	default:
		return false
	}
}

func (s *semaphore) TryAcquireWithTime(timeout time.Duration) bool {
	for {
		select {
		case s.chans <- DefaultVal:
			return true
		case <-time.After(timeout):
			return false
		}
	}
}

