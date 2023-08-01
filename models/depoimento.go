package models

type Depoimento struct {
	Foto       string `json:"foto"`
	Depoimento string `json:"depoimento"`
	NomePessoa string `json:"nome_pessoa"`
}
