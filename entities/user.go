package entities

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserRequest struct {
	Name            string `json:"name" binding:"required,min=3,max=255"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=6,max=255"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}
