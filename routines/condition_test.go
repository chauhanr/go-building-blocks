package routines

import "testing"

func TestQueueImplUsingCond(t *testing.T) {
	QueueImplUsingCond()

	t.Logf("End Condition usage.\n")
}

func TestBroadCastwithCondFunc(t *testing.T) {
	BroadCastwithCondFunc()
	t.Logf("Broadcast testing done !")
}
