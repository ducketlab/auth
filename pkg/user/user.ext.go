package user

import (
	"github.com/ducketlab/mingo/exception"
	"github.com/ducketlab/mingo/types/ftime"
	"golang.org/x/crypto/bcrypt"
)

const (
	// DefaultExiresDays todo
	DefaultExiresDays = 90
)

func New(req *CreateAccountRequest) (*User, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	pass, err := NewHashedPassword(req.Password)
	if err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	u := &User{
		CreateAt:       ftime.Now().Timestamp(),
		UpdateAt:       ftime.Now().Timestamp(),
		Profile:        req.Profile,
		DepartmentId:   req.DepartmentId,
		Account:        req.Account,
		CreateType:     req.CreateType,
		Type:           req.UserType,
		ExpiresDays:    req.ExpiresDays,
		Description:    req.Description,
		HashedPassword: pass,
		Status: &Status{
			Locked: false,
		},
	}

	if req.DepartmentId != "" && req.Profile != nil {
		u.IsInitialized = true
	}

	return u, nil
}

func NewHashedPassword(password string) (*Password, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, err
	}

	return &Password{
		Password: string(bytes),
		CreateAt: ftime.Now().Timestamp(),
		UpdateAt: ftime.Now().Timestamp(),
	}, nil
}

func (s *Set) Add(u *User) {
	s.Items = append(s.Items, u)
}

func NewProfile() *Profile {
	return &Profile{}
}

func (p *Password) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(p.Password), []byte(password))
	if err != nil {
		return exception.NewUnauthorized("user or password not correct")
	}
	return nil
}

func (u *User) Desensitize() {
	if u.HashedPassword != nil {
		u.HashedPassword.Password = ""
		u.HashedPassword.History = []string{}
	}
	return
}

func NewUserSet() *Set {
	return &Set{
		Items: []*User{},
	}
}

func NewDefaultUser() *User {
	return &User{
		Profile: NewProfile(),
		Status: &Status{
			Locked: false,
		},
	}
}