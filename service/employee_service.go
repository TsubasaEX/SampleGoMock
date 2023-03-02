package service

import (
	"SampleGoMock/repo"
	"fmt"
)

type EmployeeService interface {
	GetSrEmployeeNumbers() int
}

type EmployeeSerivceImpl struct {
	EmpRepo repo.EmployeeRepo
}

func init() {
	fmt.Println("init service.....")
}

func (es *EmployeeSerivceImpl) GetSrEmployeeNumbers(age int) int {
	srEmps := es.EmpRepo.FindEmployeesAgeGreaterThan(age)
	if false {
		fmt.Println("123")
	}
	return len(srEmps)
}

func (es *EmployeeSerivceImpl) PrintNum(age int) {
	fmt.Println(age)
}
