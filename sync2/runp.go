package sync2

import "sync"

func RunParallel[T any](data []T, fn func(T) error) []error {
	var wg sync.WaitGroup
	var errMu sync.Mutex
	var errs []error

	wg.Add(len(data))
	for i := range data {
		go func(i int) {
			defer wg.Done()
			if err := fn(data[i]); err != nil {
				errMu.Lock()
				errs = append(errs, err)
				errMu.Unlock()
			}
		}(i)
	}
	wg.Wait()

	return errs
}
