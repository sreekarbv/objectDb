package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"reflect"
)

type Person struct {
	Name     string `json:"name"`
	ID       string `json:"id"`
	LastName string `json:"last_name"`
	Birthday string `json:"birthday"`
}

func (p *Person) GetKind() string {
	return reflect.TypeOf(p).String()
}
func (p *Person) GetID() string {
	return p.ID
}
func (p *Person) GetName() string {
	return p.Name
}

func (p *Person) GetLastName() string {
	return p.LastName
}

func (p *Person) GetBirthday() string {
	return p.Birthday
}
func (p *Person) SetID(s string) {
	p.ID = s
}
func (p *Person) SetName(s string) {
	p.Name = s
}

func (p *Person) SetLastName(s string) {
	p.LastName = s
}

func (p *Person) SetBirthday(s string) {
	p.Birthday = s
}

func unmarshalPerson(input []byte) (ret Person, err error) {
	err = json.Unmarshal(input, &ret)
	if err == nil {
		if ret.Birthday != "" {
			fmt.Println("GetObjectByName Person Object ", ret, err)
			return
		}
		err = errors.New("Not of type Person")
		return
	}
	return
}

func bindPerson(ctx *gin.Context) (p *Person, err error) {
	err = ctx.ShouldBindBodyWith(&p, binding.JSON)
	if err == nil {
		if p.Birthday != "" {
			fmt.Println("Received Person Object")
			return
		}
		err = errors.New("Not of type Person")
		return
	}
	return
}
