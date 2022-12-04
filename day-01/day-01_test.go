package day01

import (
	_ "embed"
	"testing"
)

func Test_day01(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Day01()
		})
	}
}
