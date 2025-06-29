package counter

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("initial value is zero", func(t *testing.T) {
		c := &Counter{}
		if got := c.Value(); got != 0 {
			t.Errorf("want 0, got %d", got)
		}
	})

	t.Run("increment once", func(t *testing.T) {
		c := &Counter{}
		c.Inc()
		if got := c.Value(); got != 1 {
			t.Errorf("want 1, got %d", got)
		}
	})

	t.Run("concurrent increments", func(t *testing.T) {
		t.Parallel() // run this test in parallel with others

		c := &Counter{}
		var wg sync.WaitGroup
		n := 1000

		for i := 0; i < n; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.Inc()
			}()
		}

		wg.Wait()

		if got := c.Value(); got != n {
			t.Errorf("want %d, got %d", n, got)
		}
	})
}

func BenchmarkCounter(b *testing.B)  {
  c := &Counter{}

  b.ResetTimer()

  for i := 0; i < b.N; i++ {
    c.Inc()
  }
}
