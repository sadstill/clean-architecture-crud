package util

import (
	"fmt"
	"time"
)

func DoWithRetries(fn func() error, attempts int, delay time.Duration) error {
	var err error
	for attempts > 0 {
		if err = fn(); err == nil {
			return nil
		}
		time.Sleep(delay)
		attempts--
	}
	return fmt.Errorf("all attempts failed. Last error: %w", err)
}
