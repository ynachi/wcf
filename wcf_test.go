package wcf

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestCountLines(t *testing.T) {
	type args struct {
		rd io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{"one_line", args{strings.NewReader("hello\n")}, 1, false},
		{"two_line", args{strings.NewReader("hello world\nAlban")}, 2, false},
		{"empty_line", args{strings.NewReader("\n")}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CountLines(tt.args.rd)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountLines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CountLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountWords(t *testing.T) {
	type args struct {
		rd io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    int
		want1   map[string]int
		wantErr bool
	}{
		// TODO: Add test cases.
		{"one_line", args{strings.NewReader("hello\n")}, 1, map[string]int{"hello": 1}, false},
		{"two_line", args{strings.NewReader("hello world\nhello.")}, 3, map[string]int{"hello": 2, "world": 1}, false},
		{"non_alpha_only", args{strings.NewReader(". \n.")}, 0, map[string]int{}, false},
		{"empty_line", args{strings.NewReader("\n")}, 0, map[string]int{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := CountWords(tt.args.rd)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountWords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CountWords() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CountWords() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTopK(t *testing.T) {
	type args struct {
		input map[string]int
		k     int
	}
	tests := []struct {
		name    string
		args    args
		want    []Word
		wantErr bool
	}{
		{"tree_words", args{map[string]int{"hello": 2, "world": 1, "force": 3}, 2}, []Word{{"force", 3}, {"hello", 2}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TopK(tt.args.input, tt.args.k)
			if (err != nil) != tt.wantErr {
				t.Errorf("TopK() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TopK() = %v, want %v", got, tt.want)
			}
		})
	}
}
