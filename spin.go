package main

import "sync"

type spinner struct {
	mu     sync.Mutex
	pos    int
	frames []rune
	width  int
}

func (s *spinner) bump() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	w := s.width
	if w < 1 {
		w = 1
	}
	r := make([]rune, w)
	for j := 0; j < w; j++ {
		i := (s.pos + j) % len(s.frames)
		r[j] = s.frames[i]
	}
	s.pos++
	return string(r)
}
