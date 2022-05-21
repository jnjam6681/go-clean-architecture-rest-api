package entity

type Todo struct {
	ID        int32  `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

func (t Todo) TableName() string {
	return "todo"
}
