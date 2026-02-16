package pedalea

type UserData struct {
	Email     string `json:"email" validation:"required,email"`
	FirstName string `json:"first_name" validation:"required"`
	LastName  string `json:"last_name" validation:"required"`
}

type InsertUser struct {
	Password string `json:"password" validation:"required"`
	UserData
}

type LoginUser struct {
	Email    string `json:"email" validation:"required,email"`
	Password string `json:"password" validation:"required"`
}

type User struct {
	ID             int
	CreatedAt      string
	UpdatedAt      string
	HashedPassword string
	UserData
}
