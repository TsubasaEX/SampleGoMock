package service

import (
	"fmt"
	"localTestGOLand/repo"
)

var abc = 100

type EmployeeService interface {
	GetSrEmployeeNumbers() int
}

type EmployeeSerivceImpl struct {
	EmpRepo repo.EmployeeRepo
}

func (es *EmployeeSerivceImpl) GetSrEmployeeNumbers(age int) int {
	srEmps := es.EmpRepo.FindEmployeesAgeGreaterThan(age)
	if false {
		fmt.Println("123")
	}
	return len(srEmps)
}
