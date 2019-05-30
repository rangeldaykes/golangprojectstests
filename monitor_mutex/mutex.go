package monitor_mutex

import (
	"sync"
)

type SafeCounter struct {
	Mu    *sync.Mutex
	Count int
}

func (s *SafeCounter) Inc() {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	s.Count++
}
