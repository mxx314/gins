package models

type User struct {
	Name string `json:"name" form:"name" `
	Age  int    `json:"age" form:"age" `
	City string `json:"city" form:"city" `
}
