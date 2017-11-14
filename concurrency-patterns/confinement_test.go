package concurrency_patterns

import "testing"

func TestConfinementWithChannels(t *testing.T) {
	ConfinementWithChannels()

	t.Logf("Confinement Channel test complete!")
}

func TestConfinementWithChannel2(t *testing.T) {
	ConfinementWithChannel2()
	t.Logf("Confinement Channel 2 test complete!")
	}
