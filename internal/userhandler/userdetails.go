package userhandler

type UserDetails struct {
	Uuid     string `json:"uuid"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}
