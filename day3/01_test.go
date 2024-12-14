package day0301

import (
	_ "embed"
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "should return total"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Solve()
		})
	}
}

func Test_getMulOperations(t *testing.T) {
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
			args: args{input: "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"},
			want: []string{
				"mul(2,4)", "mul(5,5)", "mul(11,8)", "mul(8,5)",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMulOperations(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getMulOperations() = %v, want %v", got, tt.want)
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
		{
			name: "should get 2, 4",
			args: args{mulOperation: "mul(2,4)"},
			want: [2]int{2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMultiples(tt.args.mulOperation); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getMultiples() = %v, want %v", got, tt.want)
			}
		})
	}
}
