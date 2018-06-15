package asyncq

import (
	"log"
)

var (
	// TaskQueue is a place to submit tasks.
	TaskQueue chan Task

	// TaskWorkerQueue is a buffered channel of Task channel.
	TaskWorkerQueue chan chan Task
)

// TaskWorker holds a task worker.
type TaskWorker struct {
	ID              int
	TaskChannel     chan Task
	TaskWorkerQueue chan chan Task
}

func init() {
	// Initialize TaskQueue as a buffered channel of size 108.
	TaskQueue = make(chan Task, 108)
}

// NewTaskWorker creates a new task worker.
func NewTaskWorker(id int, taskWorkerQueue chan chan Task) TaskWorker {
	taskWorker := TaskWorker{
		ID:              id,
		TaskChannel:     make(chan Task),
		TaskWorkerQueue: taskWorkerQueue,
	}

	return taskWorker
}

// Start starts a new task worker in its own Goroutine.
func (t *TaskWorker) Start() {
	go func() {
		// Run forever to make sure that the task runner will always be
		// running. It receives incoming TaskChannel and call the Perform
		// method.
		for {
			t.TaskWorkerQueue <- t.TaskChannel

			select {
			case task := <-t.TaskChannel:
				log.Printf("Asynchronous task worker #%d is performing a task.\n", t.ID)
				task.Perform()
			}
		}
	}()
}

// StartTaskDispatcher is responsible for accepting tasks from the task queue
// and dispatching them to the next available task worker. It takes an integer
// of taskWorkerSize, which specifies the maximum number of task workers to
// start.
func StartTaskDispatcher(taskWorkerSize int) {
	// Create a buffered channel of TaskWorkerQueue given the size.
	TaskWorkerQueue = make(chan chan Task, taskWorkerSize)

	// Create task workers and call Start to let them wait for the incoming
	// tasks.
	for i := 0; i < taskWorkerSize; i++ {
		log.Printf("Starting asynchronous task worker #%d.\n", i+1)
		taskWorker := NewTaskWorker(i+1, TaskWorkerQueue)
		taskWorker.Start()
	}

	// Spin up a new Goroutine to listen to task requests.
	// Use TaskQueue channel to receive new tasks to perform.
	go func() {
		for {
			select {
			case task := <-TaskQueue:
				go func() {
					taskChannel := <-TaskWorkerQueue
					taskChannel <- task
				}()
			}
		}
	}()
}
