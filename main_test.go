package main

import (
	"fmt"
	"os"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name   string
		size   int
		sample int
	}{
		{
			name:   "case 4",
			size:   4,
			sample: 1,
		},
		{
			name:   "case 8",
			size:   8,
			sample: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldArgs := os.Args
			flags := []string{"cmd"}
			flags = append(flags, fmt.Sprintf("-size=%d", tt.size))
			flags = append(flags, fmt.Sprintf("-sample=%d", tt.sample))
			os.Args = flags
			main()
			os.Args = oldArgs
		})
	}
}
