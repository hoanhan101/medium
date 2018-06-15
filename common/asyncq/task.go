package asyncq

// Task is an interface for an asynchronous task.
type Task interface {
	Perform()
}
