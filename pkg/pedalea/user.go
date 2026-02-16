package pedalea

type UserData struct {
	Email     string
	FirstName string
	LastName  string
}

type InsertUser struct {
	Password string
	UserData
}

type LoginUser struct {
	Email    string
	Password string
}

type User struct {
	ID             int
	CreateAt       string
	UpdateAt       string
	HashedPassword string
	UserData
}
