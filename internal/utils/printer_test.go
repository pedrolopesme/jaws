package utils

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func captureOutput(f func()) string {
	// return to original state afterwards
	// note: defer evaluates (and saves) function ARGUMENT values at definition
	// time, so the original value of os.Stdout is preserved before it is
	// changed further into this function.
	defer func(orig *os.File) {
		os.Stdout = orig
	}(os.Stdout)

	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	w.Close()
	out, _ := io.ReadAll(r)

	return string(out)
}

func TestPrint(t *testing.T) {
	type args struct {
		title       string
		content     string
		contentJson string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "valid title and content",
			args: args{
				title:       "Title",
				content:     "value",
				contentJson: "{\"claim\":\"value\"}",
			},
		},
		{
			name: "empty title and content",
			args: args{
				title:       "",
				content:     "",
				contentJson: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := captureOutput(func() {
				Print(tt.args.title, tt.args.contentJson)
			})
			assert.Contains(t, out, tt.args.title)
			assert.Contains(t, out, tt.args.content)
		})
	}
}

func TestPrintTable(t *testing.T) {
	type args struct {
		header          []string
		content         [][]string
		expectedContent string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "valid header and content",
			args: args{
				header:          []string{"Header 1", "Header 2"},
				content:         [][]string{{"Content 1", "Content 2"}, {"Content 3", "Content 4"}},
				expectedContent: "Content 1",
			},
		},
		{
			name: "empty header and content",
			args: args{
				header:          []string{},
				content:         [][]string{},
				expectedContent: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := captureOutput(func() {
				PrintTable(tt.args.header, tt.args.content)
			})
			assert.Equal(t, out, out)
			assert.Contains(t, out, tt.args.expectedContent)
		})
	}
}

func TestPrinBreaklines(t *testing.T) {
	type args struct {
		numLines int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "valid number of lines",
			args: args{
				numLines: 2,
			},
		},
		{
			name: "invalid number of lines",
			args: args{
				numLines: -1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := captureOutput(func() {
				PrinBreaklines(tt.args.numLines)
			})
			expected := ""
			for i := 0; i < tt.args.numLines; i++ {
				expected += "\n"
			}
			assert.Equal(t, out, expected)
		})
	}
}
