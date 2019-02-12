package main

import (
	"encoding/json"
	"log"
	"reflect"
	"testing"
)

func Test_handleBadReqest(t *testing.T) {
	want, err := json.Marshal("Bad request")
	if err != nil {
		log.Fatal(err)
	}
	tests := []struct {
		name string
		want []byte
	}{
		// TODO: Add test cases.
		{"", want},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handleBadReqest(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handleBadReqest() = %v, want %v", got, tt.want)
			}
		})
	}
}
