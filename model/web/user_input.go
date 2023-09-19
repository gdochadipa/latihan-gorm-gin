package web

type RegisterUserInput struct {
	Name       string `json:"name" binding:"required"  form:"name"`
	Occupation string `json:"occupation" binding:"required"  form:"occupation"`
	Email      string `json:"email" binding:"required,email"  form:"email"`
	Password   string `json:"password" binding:"required"  form:"password"`
}

type LoginInput struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email" form:"email"`
}

type FormCreateUserInput struct {
	Name       string `form:"name" binding:"required"`
	Email      string `form:"email" binding:"required,email"`
	Occupation string `form:"occupation" binding:"required"`
	Password   string `form:"password" binding:"required"`
	Error      error
}

type FormUpdateUserInput struct {
	ID         int
	Name       string `form:"name" binding:"required"`
	Email      string `form:"email" binding:"required,email"`
	Occupation string `form:"occupation" binding:"required"`
	Error      error
}
