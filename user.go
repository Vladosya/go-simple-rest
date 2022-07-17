package todo

type User struct {
	Id      int    `json:"id"` // json:".." чтобы корректо принимать и выводить в http запросах
	Name    string `json:"name" binding:"required"`
	Surname string `json:"surname" binding:"required"`
}
