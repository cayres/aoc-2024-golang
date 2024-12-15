package day0302

import (
	_ "embed"
	"reflect"
	"testing"
)

func TestSolvePart2(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "should return total",
			want: 48,
		},
	}
	input = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolvePart2(); got != tt.want {
				t.Errorf("SolvePart2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getInstructions(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "should get 4 mul operations",
			args: args{input: "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"},
			want: []string{
				"mul(2,4)", "don't()", "mul(5,5)", "mul(11,8)", "do()", "mul(8,5)",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getInstructions(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getInstructions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMultiples(t *testing.T) {
	type args struct {
		mulOperation string
	}
	tests := []struct {
		name string
		args args
		want [2]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMultiples(tt.args.mulOperation); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getMultiples() = %v, want %v", got, tt.want)
			}
		})
	}
}
