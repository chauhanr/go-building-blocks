package routines

import "fmt"

// Task is the unit of work that the worker needs to take up.
type Task struct {
	Id int
}

func (t *Task) Run() {
	fmt.Printf("Running task with id : %d\n", t.Id)
}

func StartTasks(tasksChannel chan *Task, n int) {
	for i := 0; i < n; i++ {
		t := Task{i + 1}
		tasksChannel <- &t
	}
	close(tasksChannel)
}

func Worker(in chan *Task) {
	t := <-in
	t.Run()
}
