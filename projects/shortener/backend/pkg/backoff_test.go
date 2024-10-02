package pkg

import (
	"errors"
	"testing"
)

func TestProgressiveRetry(t *testing.T) {
	var retries int

	type args struct {
		operation func() error
	}
	tests := []struct {
		name            string
		args            args
		wantErr         bool
		expectedRetries int
	}{
		{
			name: "Should fail and return error",
			args: args{
				operation: func() error {
					retries++
					return errors.New("Operation Failed")
				},
			},
			wantErr:         true,
			expectedRetries: 3,
		},
		{
			name: "Should succeed after two retries",
			args: args{
				operation: func() error {
					retries++
					if retries == 2 {
						return nil
					}

					return errors.New("Operation Failed")
				},
			},
			wantErr:         false,
			expectedRetries: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			retries = 0

			err := ProgressiveRetry(tt.args.operation)
			if !tt.wantErr && (err != nil) {
				t.Errorf("expected success but got error %s", err.Error())
			}

			if tt.expectedRetries != retries {
				t.Errorf("mismatch retries count")
			}
		})
	}
}
