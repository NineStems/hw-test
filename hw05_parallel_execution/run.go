package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var (
	ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
	ErrInvalidRoutineNum   = errors.New("count of select routines less or equal zero")
)

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if m < 1 {
		return ErrErrorsLimitExceeded
	}
	if n < 1 {
		return ErrInvalidRoutineNum
	}

	w := &sync.WaitGroup{}
	mt := &sync.Mutex{}
	var ecnt int32

	for i := 0; i < n; i++ {
		w.Add(1)
		go func(wg *sync.WaitGroup, mu *sync.Mutex) {
			defer wg.Done()
			var task func() error

			for len(tasks) > 0 && ecnt <= int32(m) {
				mu.Lock()
				if len(tasks) == 0 || ecnt >= int32(m) {
					mu.Unlock()
					break
				}
				task, tasks = tasks[0], tasks[1:]
				mu.Unlock()
				if err := task(); err != nil {
					atomic.AddInt32(&ecnt, 1)
				}
			}
		}(w, mt)
	}
	w.Wait()
	if ecnt >= int32(m) {
		return ErrErrorsLimitExceeded
	}
	return nil
}
