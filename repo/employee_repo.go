package repo

import (
	"SampleGoMock/model"
	"fmt"
	"os"
	"sort"
	"strings"
)

type EmployeeRepo interface {
	FindEmployeesAgeGreaterThan(age int) []model.Employee
}

type EmployeeRepoImpl struct {
}

func init() {
	for _, v := range os.Args {
		if strings.Contains(v, "test") {
			fmt.Println("skipping testing")
			return
		}
	}
	fmt.Println("init repo.....")
}

func (er *EmployeeRepoImpl) FindEmployeesAgeGreaterThan(age int) []model.Employee {
	m := map[int]model.Employee{
		1: {ID: 1, Name: "John", Age: 33},
		2: {ID: 2, Name: "Mary", Age: 21},
		3: {ID: 3, Name: "Mike", Age: 45},
	}
	emps := []model.Employee{}
	for _, emp := range m {
		if emp.Age > age {
			emps = append(emps, emp)
		}
	}
	sort.Slice(emps, func(i, j int) bool {
		return emps[i].Age < emps[j].Age // sort by Age asc
	})
	return emps
}

func (es *EmployeeRepoImpl) PrintNum(age int) {
	fmt.Println(age)
}
