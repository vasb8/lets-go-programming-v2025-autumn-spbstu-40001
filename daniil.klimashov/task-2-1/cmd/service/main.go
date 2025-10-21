package main

import (
	"errors"
	"fmt"
)

var ErrInvalidOperator = errors.New("invalid operator")

type department struct {
	minTemp int
	maxTemp int
}

func newDepartment() department {
	return department{minTemp: 15, maxTemp: 30}
}

func (d *department) updateDesiredTemperature(op string, temp int) error {
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
		return fmt.Errorf("operation '%s': %w", op, ErrInvalidOperator)
	}

	return nil
}

func (d *department) getDesiredTemperature() int {
	if d.minTemp > d.maxTemp {
		return -1
	} else {
		return d.minTemp
	}
}

func main() {
	var departmentAmount int

	if _, err := fmt.Scan(&departmentAmount); err != nil {
		fmt.Println("Invalid number or departments", err)

		return
	}

	for range departmentAmount {
		var employeeAmount int
		temperature := newDepartment()

		if _, err := fmt.Scan(&employeeAmount); err != nil {
			fmt.Println("Invalid number of employees", err)

			return
		}

		for range employeeAmount {
			var (
				operator string
				temp     int
			)

			if _, err := fmt.Scan(&operator); err != nil {
				fmt.Println("Invalid operator", err)

				return
			}

			if _, err := fmt.Scan(&temp); err != nil {
				fmt.Println("Invalid temperature limit value", err)

				return
			}

			err := temperature.updateDesiredTemperature(operator, temp)
			if err != nil {
				fmt.Println("Invalid operator")

				return
			}

			fmt.Println(temperature.getDesiredTemperature())
		}
	}
}
