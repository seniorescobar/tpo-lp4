package entities

type TodoWithId struct {
	Id int `json:"id"`
	Todo
}

type Todo struct {
	Id          string `bson:"_id" json:"_id""`
	Description string `json:"description"`
}
