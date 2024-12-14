package part2

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

//go:embed input
var input string

func Solve() {
	safeReports := CountSafeReports(input)
	fmt.Println(safeReports)

}

func CountSafeReports(input string) int {
	rows := strings.Split(input, "\n")
	safeReports := 0
	invalidReports := []Report{}
	for _, row := range rows {
		r := strings.Split(row, " ")
		report := Report{
			asc:  []int{},
			desc: []int{},
		}

		for i := 0; i < len(r); i++ {
			n, _ := strconv.Atoi(r[i])
			report.asc = append(report.asc, n)
			n, _ = strconv.Atoi(r[len(r)-1-i])
			report.desc = append(report.desc, n)
		}

		if isValid(report) {
			safeReports++
			continue
		}

		invalidReports = append(invalidReports, report)
	}

	for _, report := range invalidReports {
		for i := 0; i < len(report.asc); i++ {
			dampenedReport := Report{
				asc:  removeItem(deepCopy(report.asc), i),
				desc: removeItem(deepCopy(report.desc), i),
			}

			if isValid(dampenedReport) {
				safeReports++
				break
			}
		}
	}

	return safeReports
}

type Report struct {
	asc  []int
	desc []int
}

func isValid(report Report) bool {
	isSorted := slices.IsSorted(report.asc) || slices.IsSorted(report.desc)

	if !isSorted {
		return false
	}

	r := report.asc
	if slices.IsSorted(report.desc) {
		r = report.desc
	}

	for i := 0; i < len(r)-1; i++ {
		difference := math.Abs(float64(r[i] - r[i+1]))
		if difference < 1 || difference > 3 {
			return false
		}
	}

	return true
}

func removeItem(arr []int, index int) []int {
	return append(arr[:index], arr[index+1:]...)
}

func deepCopy(arr []int) []int {
	newArr := make([]int, len(arr))
	copy(newArr, arr)
	return newArr
}
