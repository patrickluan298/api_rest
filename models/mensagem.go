package models

type Mensagem struct {
	Mensagem string `json:"mensagem"`
}

type MensagemDeDados struct {
	Mensagem string      `json:"mensagem"`
	Dados    interface{} `json:"dados"`
}
