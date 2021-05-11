package main

import "testing"

func TestPrint_helloworld(t *testing.T) {
	tests := []struct {
		name string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Print_helloworld()
		})
	}
}
