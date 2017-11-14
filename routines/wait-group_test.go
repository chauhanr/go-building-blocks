package routines

import "testing"

func TestForkJoinFunction(t *testing.T) {
	ForkJoinFunction()
	t.Logf("After the Routine run.")
}


func TestSleepingRoutines(t *testing.T) {
	SleepingRoutines()
	t.Logf("After Sleeping Routines test.")
}

func TestBulkWGUseage(t *testing.T) {
	BulkWGUseage()
	t.Logf("After Bulk WG call")
}