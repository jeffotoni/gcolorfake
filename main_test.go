/*
* @project     gcolor
* @package     gcolor
* @author      @jeffotoni
* @size        21/08/2017
*
 */

package gcolorfake

import "testing"

func TestBlackCor(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"black_", args{"texto para colorir black"}, "\033[0;30mtexto para colorir black\033[0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BlackCor(tt.args.msg); got != tt.want {
				t.Errorf("BlackCor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedCor(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"red_", args{"texto para colorir red"}, "\033[0;31mtexto para colorir red\033[0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RedCor(tt.args.msg); got != tt.want {
				t.Errorf("RedCor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGreenCor(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"green_", args{"texto para colorir green"}, "\033[0;32mtexto para colorir green\033[0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GreenCor(tt.args.msg); got != tt.want {
				t.Errorf("GreenCor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYellowCor(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"yellow_", args{"texto para colorir yellow"}, "\033[0;33mtexto para colorir yellow\033[0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := YellowCor(tt.args.msg); got != tt.want {
				t.Errorf("YellowCor() = %v, want %v", got, tt.want)
			}
		})
	}
}
