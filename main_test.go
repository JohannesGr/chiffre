package main

import (
	"reflect"
	"testing"
)

var zeichenErlaubt string = "abcdefghijklmnopqrst"
var mitLeerzeichen string = "abcdefgh ijklmnopqrst"
var mitZeichen string = "!§$%abcdefghijklmnopqrst123"

func Test_encrypt(t *testing.T) {

	tests := []struct {
		name string
		args string
		key  string
		want string
	}{
		{"encrypt", "LOREMIPSUMDOLORSIT", "CTMAGAZIN", "NHDESIOAHOWALURRQG"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encrypt(&tt.args, &tt.key); !reflect.DeepEqual(got, &tt.want) {
				t.Errorf("encrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeIllegalChars(t *testing.T) {

	tests := []struct {
		name string
		text *string
		want *string
	}{
		{"erlaubt", &zeichenErlaubt, &zeichenErlaubt},
		{"leerzeichen", &mitLeerzeichen, &zeichenErlaubt},
		{"mitZeichen", &mitZeichen, &zeichenErlaubt},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeIllegalChars(tt.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeIllegalChars() = %v, want %v", got, tt.want)
			}
		})
	}
}
