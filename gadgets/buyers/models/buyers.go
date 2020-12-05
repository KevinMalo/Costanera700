package models

import "errors"

const maxLengthInId = 8
const maxLengthInName = 100
const maxLengthInAge = 255

//Buyer model structure for buyers
type Buyer struct {
	Id   string	// max 8 chars
	Name string // max 100 chars
	Age  uint8 // 0 to 255
}

//CreateBuyerCMD command to create a new review
type CreateBuyerCMD struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

func (cmd *CreateBuyerCMD) validate() error {
	if len(cmd.Id) > maxLengthInId {
		return errors.New("id must be between 1-8 chars")
	}

	if len(cmd.Name) > maxLengthInName {
		return errors.New("name must be less than 100 chars")
	}

	if cmd.Age > maxLengthInAge {
		return errors.New("name must be less than 255 chars")
	}

	return nil
}