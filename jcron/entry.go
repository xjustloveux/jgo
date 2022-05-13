package jcron

import "sync"

type entry struct {
	sch  *schedule
	j    Job
	data map[string]interface{}
}

func (e *entry) run(w *sync.WaitGroup) {
	defer func() {
		w.Done()
	}()
	e.j.Run(e.data)
}
