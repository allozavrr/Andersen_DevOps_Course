package main

import (
	
	"testing"

)

func TestHelloWithTable(t *testing.T) {
    tests := []struct {
        name string
        want string
    }{
        // TODO: Add test cases.
        {
            name: "test for hello",
            want: "Hello World",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Hello(); got != tt.want {
                t.Errorf("Hello() = %v, want %v", got, tt.want)
            }
        })
    }
}
