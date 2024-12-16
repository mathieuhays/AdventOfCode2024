package main

import (
	"errors"
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_lineToInts(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  error
		want1 int
		want2 int
	}{
		{"one space", "1 3", nil, 1, 3},
		{"two spaces", "2  4", nil, 2, 4},
		{"three spaces", "3  6", nil, 3, 6},
		{"padded", "  4  7  ", nil, 4, 7},
		{"too many numbers", "  4 1 7  ", errors.New("invalid"), 0, 0},
		{"too few numbers", "  4  ", errors.New("invalid"), 0, 0},
		{"no numbers", "   ", errors.New("invalid"), 0, 0},
		{"letters", " a b  ", errors.New("invalid"), 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := lineToInts(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lineToInts() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("lineToInts() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("lineToInts() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func Test_inputToLists(t *testing.T) {
	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name   string
		reader io.Reader
		want   error
		want1  []int
		want2  []int
	}{
		{"test 1", strings.NewReader("1 3\n2 4"), nil, []int{1, 2}, []int{3, 4}},
		{"test 2", strings.NewReader(" 1  4\n  2    4"), nil, []int{1, 2}, []int{4, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := inputToLists(tt.reader)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inputToLists() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("inputToLists() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("inputToLists() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
