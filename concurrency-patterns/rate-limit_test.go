package concurrency_patterns

import "testing"

func TestRateLimitFunc(t *testing.T){
	RateLimitFunc()

	t.Logf("Api connections tested.")
}
