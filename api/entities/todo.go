package entities

type TodoWithId struct {
	Id int `json:"id"`
	Todo
}

type Todo struct {
	Description string `json:"description"`
}
