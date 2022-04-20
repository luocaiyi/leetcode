package _491_average_salary_excluding_the_minimum_and_maximum_salary

import "testing"

func TestAverage(t *testing.T) {
	salary := []int{4000, 3000, 1000, 2000}
	println(average(salary)) // 2500
}
