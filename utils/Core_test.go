package utils

import "testing"

func TestNullOrLetter(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Null", args{""}, true},
		{"NotNull", args{"A"}, true},
		{"NotNull", args{"1"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NullOrLetter(tt.args.s); got != tt.want {
				t.Errorf("NullOrLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotNullAndLetter(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Null", args{""}, false},
		{"AlphaNum", args{"abc123"}, false},
		{"AlphaNum", args{"abcXyz"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NotNullAndLetter(tt.args.s); got != tt.want {
				t.Errorf("NotNullAndLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
