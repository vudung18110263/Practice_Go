package entities

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Login struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
