package users

type User struct {
	ID            int
	Name          string
	LastName      string
	EffectiveRate float32
}

type Service interface {
	GetUser(userID int) (User, error)
	CalcAvaiableAmount(grossRevenue, effectiveRate float32) float32
}

func NewUserService() Service {
	return &userService{}
}

type userService struct{}

func (s *userService) GetUser(userID int) (User, error) {
	u := User{
		ID:            1,
		Name:          "nono",
		LastName:      "lol",
		EffectiveRate: 0.05,
	}
	return u, nil
}

func (s *userService) CalcAvaiableAmount(grossRevenue, effectiveRate float32) float32 {
	amount := grossRevenue * effectiveRate
	return amount
}
