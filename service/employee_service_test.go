package service

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"localTestGOLand/mocks"
	"localTestGOLand/model"
	"testing"
)

//// define mock type
//type EmployeeRepoImplMock struct {
//	mock.Mock
//}
//
//// use mock to implement EmployeeRepo's method
//func (repoMock *EmployeeRepoImplMock) FindEmployeesAgeGreaterThan(age int) []model.Employee {
//	args := repoMock.Called(age)
//	return args.Get(0).([]model.Employee)
//}

type DyInterface interface {
	//Len() int
	Less(i, j int) bool
	//Swap(i, j int)
}

type MyDyStruct struct {
	DyInterface
}

func (m MyDyStruct) Less(i, j int) bool {
	return m.DyInterface.Less(j, i)
}

type DyDef struct {
}

func (d DyDef) Less(i, j int) bool { return false }

func TestGetSrEmployeeNumbers_Age40(t *testing.T) {
	my := new(MyDyStruct)
	fmt.Println(my.Less(1, 3))
	//fmt.Println(my.Len())
	fmt.Println(abc)
	repoMockery := mocks.NewEmployeeRepo(t)
	fmt.Println(10000)
	repoMockery.On("FindEmployeesAgeGreaterThan", 40).
		Return([]model.Employee{
			{ID: 99, Name: "Jack", Age: 70}, // mock return value
		})
	//repoMock := new(EmployeeRepoImplMock)
	//repoMock.On("FindEmployeesAgeGreaterThan", 40).
	//	Return([]model.Employee{
	//		{ID: 99, Name: "Jack", Age: 70}, // mock return value
	//	})

	expected := 1

	es := EmployeeSerivceImpl{
		EmpRepo: repoMockery,
		//EmpRepo: repoMock,
	}

	actial := es.GetSrEmployeeNumbers(40)
	assert.Equal(t, expected, actial)
}
