package control

type (
	Semaphore interface {
		Acquire()
		Release()
	}
	_semaphore struct {
		c chan struct{}
	}
)

func NewSemaphore(count uint64) Semaphore {
	return &_semaphore{
		c: make(chan struct{}, count),
	}
}

func (s *_semaphore) Acquire() {
	s.c <- struct{}{}
}

func (s *_semaphore) Release() {
	<-s.c
}
