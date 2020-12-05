package models

import (
	"testing"
)

func NewBuyer(id string, name string, age uint8) *CreateBuyerCMD {
	return &CreateBuyerCMD{
		Id:   id,
		Name: name,
		Age:  age,
	}
}

func Test_withCorrectParams(t *testing.T) {
	b := NewBuyer("4f0f052", "Rostand", 72)
	err := b.validate()
	if err != nil {
		t.Error("the validation did not past")
		t.Fail()
	}
}

func Test_shouldFailWithWrongLenOfId(t *testing.T) {
	b := NewBuyer("dfd236a44", "Davita", 52)
	err := b.validate()
	if err == nil {
		t.Error("should fail with more id chars")
		t.Fail()
	}
}