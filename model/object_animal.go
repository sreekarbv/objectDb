package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"reflect"
)

type Animal struct {
	Name    string `json:"name"`
	ID      string `json:"id"`
	Type    string `json:"type"`
	OwnerID string `json:"owner_id"`
}

func (a *Animal) GetKind() string {
	return reflect.TypeOf(a).String()
}
func (a *Animal) GetID() string {
	return a.ID
}
func (a *Animal) GetName() string {
	return a.Name
}

func (a *Animal) GetOwnerID() string {
	return a.OwnerID
}

func (a *Animal) GetType() string {
	return a.Type
}
func (a *Animal) SetID(s string) {
	a.ID = s
}
func (a *Animal) SetName(s string) {
	a.Name = s
}

func (a *Animal) SetOwnerID(s string) {
	a.OwnerID = s
}

func (a *Animal) SetType(s string) {
	a.Type = s
}

func unmarshalAnimal(input []byte) (ret Animal, err error) {
	err = json.Unmarshal(input, &ret)
	if err == nil {
		if ret.Type != "" {
			fmt.Println("GetObjectByName Animal Object ", ret, err)
			return
		}
		err = errors.New("Not of type Person")
		return
	}
	return
}

func bindAnimal(ctx *gin.Context) (a *Animal, err error) {
	err = ctx.ShouldBindBodyWith(a, binding.JSON)
	if err == nil {
		if a.Type != "" {
			fmt.Println("Received Animal Object")
			return
		}
		err = errors.New("Not of type Animal")
		return
	}
	return
}
