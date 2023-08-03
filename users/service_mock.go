package users

type ServiceMock struct {
	GetUserFunc            func(userID int) (User, error)
	CalcAvaiableAmountFunc func(grossRevenue, effectiveRate float32) float32
}

func (s *ServiceMock) GetUser(userID int) (User, error) {
	if s.GetUserFunc != nil {
		return s.GetUserFunc(userID)
	}
	return User{}, nil
}

func (s *ServiceMock) CalcAvaiableAmount(grossRevenue, effectiveRate float32) float32 {
	if s.CalcAvaiableAmountFunc != nil {
		return s.CalcAvaiableAmountFunc(grossRevenue, effectiveRate)
	}
	return 0.00
}
