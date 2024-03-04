package controller

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (a *Auth) Login() {}

func (a *Auth) Signup() {}
