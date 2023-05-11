package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Object interface {
	// GetKind returns the type of the object.
	GetKind() string
	// GetID returns a unique UUID for the object.
	GetID() string
	// GetName returns the name of the object. Names are not unique.
	GetName() string
	// SetID sets the ID of the object.
	SetID(string)
	// SetName sets the name of the object.
	SetName(string)
}

func UnmarshallObject(input []byte) (retObj Object, err error) {
	retPerson, err := unmarshalPerson(input)
	if err != nil {
		fmt.Println("Error while unmarshalling to Person object", err)
	} else {
		retObj = &retPerson
		return
	}
	retAnimal, err := unmarshalAnimal(input)
	if err != nil {
		fmt.Println("Error while unmarshalling to Animal object", err)
	} else {
		retObj = &retAnimal
		return
	}
	return
}

func BindObject(ctx *gin.Context) (obj Object, err error) {
	obj, err = bindPerson(ctx)
	if err == nil {
		return
	}
	return bindAnimal(ctx)
}
