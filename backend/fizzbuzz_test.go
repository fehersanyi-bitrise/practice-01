package main

import (
	"testing"
)

func Test_counter(t *testing.T) {
	tests := []struct {
		name string
		args int
		want string
	}{
		// TODO: Add test cases.
		{"", 1, "1"},
		{"", 3, "fizz"},
		{"", 5, "buzz"},
		{"", 15, "fizzbuzz"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := counter(tt.args); got != tt.want {
				t.Errorf("counter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_HandleFizzBuzz(t *testing.T) {

}
