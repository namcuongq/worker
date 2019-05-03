package woker

import (
	"sync"
)

type WokerPool struct {
	WG   sync.WaitGroup
	Jobs chan interface{}
}

func New(workers int, do func(interface{})) *WokerPool {
	var pool = new(WokerPool)
	pool.Jobs = make(chan interface{})
	for workerID := 0; workerID < workers; workerID++ {
		pool.WG.Add(1)
		go pool.work(workerID, do)
	}
	return pool
}

func (pool *WokerPool) work(id int, do func(interface{})) {
	defer pool.WG.Done()
	for {
		select {
		case job, ok := <-pool.Jobs:
			if !ok {
				return
			}

			do(job)
		}
	}
}

func (pool *WokerPool) Add(job interface{}) {
	pool.Jobs <- job
}

func (pool *WokerPool) WaitAndClose() {
	close(pool.Jobs)
	pool.WG.Wait()
}
