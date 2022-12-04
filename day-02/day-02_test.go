package day02

import (
	_ "embed"
	"testing"
)

func Test_day02(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			Day02()
		})
	}
}
