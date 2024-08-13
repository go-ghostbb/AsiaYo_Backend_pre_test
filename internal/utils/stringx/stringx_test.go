package stringx_test

import (
	"AsiaYo-Backend-pre-test/internal/utils/stringx"
	"testing"
)

func TestContainNonEnglish(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{s: "Hello World"},
			want: false,
		},
		{
			name: "case 2",
			args: args{s: "Hello World!"}, // "!" is non english
			want: true,
		},
		{
			name: "case 3",
			args: args{s: "Chang Yu Wei"},
			want: false,
		},
		{
			name: "case 4",
			args: args{s: "張 Yu Wei"}, // "張" is non english
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringx.ContainNonEnglish(tt.args.s); got != tt.want {
				t.Errorf("stringx.ContainNonEnglish() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsFirstLetterUp(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{s: "Hello World"},
			want: true,
		},
		{
			name: "case 2",
			args: args{s: "hello World"},
			want: false,
		},
		{
			name: "case 3",
			args: args{s: "Hello world"},
			want: false,
		},
		{
			name: "case 4",
			args: args{s: "hello wOrld"},
			want: false,
		},
		{
			name: "case 5",
			args: args{s: "Hello WOrld"},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringx.IsFirstLetterUp(tt.args.s); got != tt.want {
				t.Errorf("stringx.IsFirstLetterUp() = %v, want %v", got, tt.want)
			}
		})
	}
}
