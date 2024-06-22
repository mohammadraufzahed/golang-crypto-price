package worker

import (
	"fmt"
	"sync"
)

var JobPool WorkerPool

type Worker interface {
	Process(job Job) error
}

type DefaultWorker struct{}

func (w *DefaultWorker) Process(job Job) error {
	err := job.Execute()
	if err != nil {
		return err
	}
	return nil
}

type Job interface {
	Execute() error
}

type WorkerPool struct {
	workers   []Worker
	jobQueue  chan Job
	waitGroup sync.WaitGroup
}

func NewWorkerPool(workers, queueSize int) *WorkerPool {
	pool := &WorkerPool{
		workers:  make([]Worker, workers),
		jobQueue: make(chan Job, queueSize),
	}

	for i := 0; i < workers; i++ {
		pool.workers[i] = &DefaultWorker{}
	}

	return pool
}

func (pool *WorkerPool) Start() {
	for _, worker := range pool.workers {
		pool.waitGroup.Add(1)
		go func(worker Worker) {
			defer pool.waitGroup.Done()
			for job := range pool.jobQueue {
				err := worker.Process(job)
				if err != nil {
					fmt.Printf("Failed to process the job: %v", err)
				}
			}
		}(worker)
	}
}

func (pool *WorkerPool) Add(job Job) {
	pool.jobQueue <- job
}

func InitWorkerPool() WorkerPool {
	JobPool = *NewWorkerPool(10, 100)

	JobPool.Start()

	return JobPool
}
