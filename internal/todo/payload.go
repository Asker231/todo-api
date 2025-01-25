package todo

type(

	TodoRequest struct{
		Text string `json:"text"`
		Complited bool `json:"complited"`
	}
)