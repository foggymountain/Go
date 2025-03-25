package randomphrase

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_alreadyUsed(t *testing.T) {
	type args struct {
		m    usage
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test alreadyUsed",
			args: args{
				m:    usage{"hello": struct{}{}, "world": struct{}{}},
				word: "world",
			},
			want: true,
		},
		{
			name: "Test alreadyUsed",
			args: args{
				m:    usage{"hello": struct{}{}, "world": struct{}{}},
				word: "foo",
			},
			want: false,
		},
		{
			name: "Test empty",
			args: args{
				m:    usage{},
				word: "foo",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alreadyUsed(tt.args.m, tt.args.word); got != tt.want {
				t.Errorf("alreadyUsed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_NoRepeatInDictionary(t *testing.T) {
	seen := make(map[string]bool)
	fmt.Printf("Dictionary length: %v\n", len(dict))
	for _, str := range dict {
		if seen[str] {
			t.Errorf("Duplicate word in dictionary: '%v'", str)
			t.FailNow()
		}
		seen[str] = true
	}
}

func Test_CodeSpace(t *testing.T) {
	l := len(dict)
	f1, err := CodeSpace(0)
	assert.Nil(t, err)
	assert.Equal(t, l, f1)

	f2, _ := CodeSpace(1)
	assert.Equal(t, l, f2)
}

func Test_GenerateWithFormat(t *testing.T) {
	res, err := generateWithFormat([]string{"hello"}, 1, "-", true)
	assert.Nil(t, err)
	assert.Equal(t, "Hello", res)
}

func Test_build(t *testing.T) {
	type args struct {
		l int
	}
	tests := []struct {
		name   string
		args   args
		result int
	}{
		{name: "Test build 1", args: args{l: 1}, result: 1},
		{name: "Test build 2", args: args{l: 10}, result: 10},
		{name: "Test build 3", args: args{l: 20}, result: 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := len(build(dict, tt.args.l, false)); got != tt.result {
				tt.result = len(build(dict, tt.args.l, false))
				t.Errorf("build() = %v, want %v", got, tt.result)

			}
		})
	}
}
