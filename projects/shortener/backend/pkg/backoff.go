package pkg

import (
	"time"
)

var shortBackoffTimeouts = []time.Duration{500 * time.Millisecond, 1 * time.Second, 1 * time.Second}                 // 500ms, 1s, 1s
var longBackoffTimeouts = []time.Duration{500 * time.Millisecond, 1 * time.Second, 2 * time.Second, 5 * time.Second} // 500ms, 1s, 2s, 5s

// ProgressiveRetry executes a BackoffOperation and waits an increasing time before retrying the operation.
func ProgressiveRetry(operation func() error) error {
	return CustomProgressiveRetry(operation, shortBackoffTimeouts)
}

func LongProgressiveRetry(operation func() error) error {
	return CustomProgressiveRetry(operation, longBackoffTimeouts)
}

func CustomProgressiveRetry(operation func() error, backoffTimeouts []time.Duration) error {
	var err error

	for attempts := 0; attempts < len(backoffTimeouts); attempts++ {
		err = operation()
		if err == nil {
			return nil
		}

		time.Sleep(backoffTimeouts[attempts])
	}

	return err
}
