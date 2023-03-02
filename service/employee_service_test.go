package service

import (
	"SampleGoMock/mocks"
	"SampleGoMock/model"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestGetSrEmployeeNumbers_Age40(t *testing.T) {
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
