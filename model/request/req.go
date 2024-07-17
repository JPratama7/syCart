package request

type Register struct {
	Email    string `json:"email,required"`
	Username string `json:"username,required"`
	Password string `json:"password,required"`
}

type Login struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password,required"`
}
