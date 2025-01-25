package todo

import "time"

type(

	TodoRequest struct{
		Text string `json:"text"`
		Complited bool `json:"complited"`
	}

	TodoResponse struct{
		Text string `json:"text"`
		Complited bool `json:"complited"`
		CreatedAt time.Time `json:"date"`
	}
)