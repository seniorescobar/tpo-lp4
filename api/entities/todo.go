package entities

type TodoWithId struct {
	Id int `json:"id"`
	Todo
}

type Todo struct {
	Id          string `json:"_id"`
	Description string `json:"description"`
}
