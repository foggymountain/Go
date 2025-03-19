package fruid

import (
	"fmt"
	"testing"
)

func TestGenerate(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		length  int
		wantLen int
		wantErr bool
	}{
		{"length one", 1, 1, false},
		{"length ten", 10, 10, false},
		{"length one hundred", 100, 100, false},
		{"length error", -1, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Generate(tt.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateAuid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.wantLen {
				t.Errorf("GenerateAuid() = %v, want %v", got, tt.wantLen)
			}
			fmt.Println(got)
		})
	}
}

func TestGenerateUUID(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		wantLen int
		wantErr bool
	}{
		{"standard length", 21, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateUUID()
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateAuid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.wantLen {
				t.Errorf("GenerateAuid() = %v, want %v", got, tt.wantLen)
			}
			fmt.Println(got)
		})
	}
}

func TestGenerateWithAlpha(t *testing.T) {
	type args struct {
		n  int
		al string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"With Alpha", args{n: 10, al: "abcdefghijklmnop"}, false},
		{"With Alpha", args{n: 0, al: ""}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateWithAlpha(tt.args.n, tt.args.al)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateWithAlpha() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.args.n {
				t.Errorf("GenerateWithAlpha() length %v, want %v", len(got), tt.args.n)
			}
		})
	}
}
