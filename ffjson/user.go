package ffjson

type User struct {
	Id        uint64 `json:"id"`
	UserName  string `json:"user_name"`
	CellPhone string `json:"cell_phone"`
}
