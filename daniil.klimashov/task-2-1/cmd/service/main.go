package main

import "fmt"

type department struct {
	minTemp int
	maxTemp int
}

func newDepartment() *department {
	return &department{
		minTemp: 15,
		maxTemp: 30,
	}
}

func (d *department) updateDesiredTemperature(op string, temp int) {
	switch op {
	case ">=":
		if temp > d.minTemp {
			d.minTemp = temp
		}
	case "<=":
		if temp < d.maxTemp {
			d.maxTemp = temp
		}
	}
}

func (d *department) getDesiredTemperature() int {
	if d.minTemp <= d.maxTemp {
		return d.minTemp
	}
	return -1
}

func main() {
	var (
		n, k int
	)

	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Invalid number or departments")
	}

	if _, err := fmt.Scan(&k); err != nil {
		fmt.Println("Invalid number of employees")
	}

	for i := 0; i < n; i++ {

		dept := newDepartment()

		for j := 0; j < k; j++ {

			var op string
			if _, err := fmt.Scan(&op); err != nil {
				fmt.Println("Invalid operator")
			}

			var temp int
			if _, err := fmt.Scan(&temp); err != nil {
				fmt.Println("Invalid temperature limit value")
			}

			dept.updateDesiredTemperature(op, temp)

			desiredTemp := dept.getDesiredTemperature()
			fmt.Println(desiredTemp)
		}
	}

}
