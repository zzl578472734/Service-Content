package vars

type UserLoginParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserSearchParam struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Status   int8   `json:"status"`
}

type UserInsertParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Status   int8   `json:"status"`
}
