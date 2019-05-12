package vars

type AdminLoginParam struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type AdminParam struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Name     string `json:"name"`
	RoleId   int64  `json:"role_id"`
}

type AdminModifyPasswordParam struct {
	Id       int64  `json:"id"`
	Password string `json:"password"`
}

type AdminListParam struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	DefaultListParam
}
