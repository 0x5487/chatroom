package model

type Payload struct {
	Kind   string      `json:"kind"`
	Member Member      `json:"member"`
	Data   interface{} `json:"data"`
}

type Member struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type Message struct {
	Text string `json:"text"`
}
