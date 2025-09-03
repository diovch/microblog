package service

import "runtime"

type WorkerPool struct {
	tasks chan func()
}

func NewWorkerPool() *WorkerPool {
	wp := &WorkerPool{
		tasks: make(chan func()),
	}

	for range runtime.GOMAXPROCS(-1) - 2{
		go func() {
			for task := range wp.tasks {
				task()
			}
		}()
	}

	return wp
}

func (w *WorkerPool) RunAsync(task func()) {
	w.tasks <- task // FAQ: blocks if no workers are available
}

func (w *WorkerPool) Close() {
	close(w.tasks)
}