package request

type BookResponse struct {
	Id        int    `json:"id"`
	Title     string `json:"title" validate:"required"`
	Author    string `json:"author" validate:"required"`
	Cover     string `json:"cover" validate:"required"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
