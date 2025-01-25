package todo

import "gorm.io/gorm"


type(
	Todo struct{
		*gorm.Model
		Text string `json:"text"`
		Complited bool `json:"complited"`
	}
)

func NewTodo(text string , complited bool)*Todo{
	return &Todo{
		Text: text,
		Complited: complited,
	}
}