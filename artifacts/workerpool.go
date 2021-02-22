package main

import (
	"sync"
)

/*
This system is run like this:
	concurrency := 10
	// Build out the tasks to run
	var tasks []*Task
	for _, item := range items { // declare what to iterate over as tasks
		tasks = append(tasks, NewTask(func() error { return SOME_FUNCTION_TO_RUN }))
	}

	// Create and run the worker pool
	p := NewPool(tasks, concurrency)
	p.Run()

	// Check for failures in the workers.
	for _, task := range p.Tasks {
		log.Println(task.Err, "batch failed to process")
	}
*/

// Pool holds all of the tasks for the worker pool, it's the metadata of the pool essentially
type Pool struct {
	Tasks []*Task

	concurrency int
	tasksChan   chan *Task
	wg          sync.WaitGroup
}

// Task holds the actual function that will run concurrently
type Task struct {
	Err error
	f   func() error
}

// NewTask initializes a new task based on a given work
// function.
func NewTask(f func() error) *Task {
	return &Task{f: f}
}

// Run runs a Task and does appropriate accounting via a
// given sync.WorkGroup.
func (t *Task) Run(wg *sync.WaitGroup) {
	t.Err = t.f()
	wg.Done()
}

// NewPool creates a new worker pool seeded with all of the Tasks/functions that will be run
func NewPool(tasks []*Task, concurrency int) *Pool {
	return &Pool{
		Tasks:       tasks,
		concurrency: concurrency,
		tasksChan:   make(chan *Task),
	}
}

// Run actually starts the go routines for the pool
func (p *Pool) Run() {
	for i := 0; i < p.concurrency; i++ {
		go p.work()
	}

	p.wg.Add(len(p.Tasks))
	for _, task := range p.Tasks {
		p.tasksChan <- task
	}

	close(p.tasksChan)

	p.wg.Wait()
}

func (p *Pool) work() {
	for task := range p.tasksChan {
		task.Run(&p.wg)
	}
}
