package scheduler

import "learn-go/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(
	requests chan engine.Request) {
	s.workerChan = requests
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	// send request down to worker chan
	go func() {
		s.workerChan <- request
	}()
	//s.workerChan <- request //循环等待 死锁
}
