package nap

import (
	"golang.org/x/sync/errgroup"
)

func scatter(n int, fn func(i int) error) error {
	var g errgroup.Group
	for i := 0; i < n; i++ {
		x := i
		g.Go(func() error {
			return fn(x)
		})
	}

	return g.Wait()
}
