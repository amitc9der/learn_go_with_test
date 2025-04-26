package sync_counter

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("increment the  counter 3 times leaves it at 3", func(t *testing.T) {
		counter := &Counter{}

		counter.Increment()
		counter.Increment()
		counter.Increment()

		assertCounter(t, counter, 3)
	})

	t.Run("increment the  counter 1000 times ", func(t *testing.T) {

		wantedCount := 1000
		counter := &Counter{}

		var wg sync.WaitGroup
		wg.Add(wantedCount)
		for range wantedCount {
			go func() {
				counter.Increment()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t testing.TB, c *Counter, want int) {
	t.Helper()

	if c.Value() != want {
		t.Errorf("got %d want %d ", c.Value(), want)
	}

}
