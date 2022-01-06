package scheduler

import (
	"GO-CRAWLER/engine"
)

type Scheduler struct {
	requestChan chan engine.Request
}

func (s *Scheduler) ConfigureScheduler(c chan engine.Request) {
	s.requestChan = c
}

func (s *Scheduler) Submit(r engine.Request) {
	go func() {
		for {
			s.requestChan <- r
		}
	}()
}
