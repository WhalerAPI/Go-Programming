package utils_test

import (
	"golang/utils"
	"testing"
)

func TestReverseArrayCopy(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    []string
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utils.ReverseArrayCopy(tt.s)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("ReverseArrayCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}
