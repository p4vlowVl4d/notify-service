package handlers

type User struct {
	ID                 uint64
	Name, Token, Email string
	Active             bool
	Role               UserRole
}

type UserRole struct {
	Admin  bool
	Save   bool
	Update bool
	Delete bool
	Name   string
}

func (u User) getRoleString() string {
	return u.Role.Name
}
