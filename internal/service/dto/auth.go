package dto

type AuthUser struct {
	ID string
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	Token string `json:"token"`
}

type RegisterReq struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
type RegisterResp struct {
	Token string `json:"token"`
}
