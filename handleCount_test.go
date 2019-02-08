package main

import (
	"reflect"
	"testing"
)

func Test_charCount(t *testing.T) {
	tests := []struct {
		name string
		args string
		want map[string]int
	}{
		// TODO: Add test cases.
		{"", "aaa", map[string]int{"a": 3}},
		{"", "bbbb", map[string]int{"b": 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := charCount(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("charCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
