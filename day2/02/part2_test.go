package part2

import (
	"testing"
)

func TestSolve(t *testing.T) {
	Solve()
}

func TestValidReport(t *testing.T) {
	cases := []struct {
		report  Report
		expeted bool
	}{
		{
			report:  Report{asc: []int{1, 2, 3}, desc: []int{3, 2, 1}},
			expeted: true,
		},
		{
			report:  Report{asc: []int{2, 1, 3}, desc: []int{3, 1, 2}},
			expeted: false,
		},
		{
			report:  Report{asc: []int{1, 5, 10}, desc: []int{1, 5, 10}},
			expeted: false,
		},
	}

	for _, c := range cases {
		result := isValid(c.report)
		if result != c.expeted {
			t.Fatalf("Invalid report!")
		}

	}
}
