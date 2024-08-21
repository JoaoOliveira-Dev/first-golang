package entity

import (
	"encoding/json"
	"log"
)

// Interaface são regras para que determinada função ou código possa executar ou entrar em algum lugar
type UserInterface interface {
	String() string
}

type User struct{
	ID int `json:"id"`
	Nome string `json:"nome"`
	Usuario string `json:"usuario"`
	Email string `json:"email"`
	Senha string `json:"senha"`
}

func (u *User) String() string {
	data, err := json.Marshal(u)

	if err != nil{
		log.Println(err.Error())
	}

	return string(data)
}

type ListUser struct {
	Users []*User `json:"users"`
}

func NewUser(usuario *User) *User {
	return &User{
		Nome:    usuario.Nome,
		Usuario: usuario.Usuario,
		Email:   usuario.Email,
		Senha:   usuario.Senha,
	}
}