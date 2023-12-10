package request

type UsersCreate struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}

type UserId struct {
	Id int `json:"id"`
}
