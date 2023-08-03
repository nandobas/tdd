package users_test

import (
	"testing"

	"github.com/nandobas/tdd/users"
	"github.com/stretchr/testify/suite"
)

type testUserServiceSuite struct {
	suite.Suite
}

func TestUserService(t *testing.T) {
	suite.Run(t, new(testUserServiceSuite))
}

func (t *testUserServiceSuite) TestUserServiceMock_WhenGetUser_ExpectedNoError() {
	//Arrange
	var userID = 1
	var userName = "nono"
	var userLastName = "lol"

	serviceMock := users.ServiceMock{
		GetUserFunc: func(id int) (users.User, error) {
			u := users.User{
				ID:       1,
				Name:     "nono",
				LastName: "lol",
			}

			return u, nil
		},
	}

	//Action
	u, err := serviceMock.GetUser(userID)
	t.NoError(err)

	//Assert
	t.Equal(userID, u.ID)
	t.Equal(userName, u.Name)
	t.Equal(userLastName, u.LastName)

}

func (t *testUserServiceSuite) TestUserService_WhenGetUser_ExpectedNoError() {
	//Arrange
	var userID = 1
	var userName = "nono"
	var userLastName = "lol"

	userService := users.NewUserService()

	//Action
	u, err := userService.GetUser(userID)
	t.NoError(err)

	//Assert
	t.Equal(userID, u.ID)
	t.Equal(userName, u.Name)
	t.Equal(userLastName, u.LastName)

}
func (t *testUserServiceSuite) TestUserService_WhenCalcAvaiableAmount_ExpectedAmountToPay() {
	//Arrange
	grossRevenue := float32(2372.00)
	effectiveRate := float32(0.05)
	expectedAmountToPay := float32(118.6)

	userService := users.NewUserService()

	//Action
	givenAmountToPay := userService.CalcAvaiableAmount(grossRevenue, effectiveRate)

	//Assert
	t.Equal(expectedAmountToPay, givenAmountToPay)

}
