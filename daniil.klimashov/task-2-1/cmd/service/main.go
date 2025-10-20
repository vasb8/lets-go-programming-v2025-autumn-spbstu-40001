package main

import "fmt"

const (
	defaultMinTemperature = 15
	defaultMaxTemperature = 30
)

type department struct {
	minTemp int
	maxTemp int
}

func (d *department) updateAndGetDesiredTemperature(op string, temp int) {
	switch op {
	case ">=":
		if temp > d.minTemp {
			d.minTemp = temp
		}
	case "<=":
		if temp < d.maxTemp {
			d.maxTemp = temp
		}
	default:
		fmt.Println("Wrong operation")

		return
	}

	if d.minTemp > d.maxTemp {
		fmt.Println(-1)
	} else {
		fmt.Println(d.minTemp)
	}
}

func main() {
	var departmentAmount int

	if _, err := fmt.Scan(&departmentAmount); err != nil {
		fmt.Println("Invalid number or departments")

		return
	}

	for range departmentAmount {
		var (
			employeeAmount int
			temperature    = department{15, 30}
		)

		if _, err := fmt.Scan(&employeeAmount); err != nil {
			fmt.Println("Invalid number of employees")

			return
		}

		for range employeeAmount {
			var (
				operator string
				temp     int
			)

			if _, err := fmt.Scan(&operator); err != nil {
				fmt.Println("Invalid operator")

				return
			}

			if _, err := fmt.Scan(&temp); err != nil {
				fmt.Println("Invalid temperature limit value")

				return
			}

			temperature.updateAndGetDesiredTemperature(operator, temp)
		}
	}
}
