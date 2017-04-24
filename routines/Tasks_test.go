package routines

import "testing"

func TestTasksWithChannels(t *testing.T) {
	tasksChannel := make(chan *Task)
	N := 5

	go StartTasks(tasksChannel, N)

	for i := 0; i < N; i++ {
		// starting N channels to process each task.
		go Worker(tasksChannel)
	}
}
