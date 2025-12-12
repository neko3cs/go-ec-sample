package domain

type User struct {
	id       uint
	name     string
	email    string
	password string
	isAdmin  bool
}

func NewUser(id uint, name, email, password string, isAdmin bool) *User {
	return &User{
		id:       id,
		name:     name,
		email:    email,
		password: password,
		isAdmin:  isAdmin,
	}
}

func (u *User) Id() uint { return u.id }

func (u *User) Name() string { return u.name }

func (u *User) Email() string { return u.email }

func (u *User) Password() string { return u.password }

func (u *User) IsAdmin() bool { return u.isAdmin }
