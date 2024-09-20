package models

import (
	"github.com/EClaesson/go-luhn"
	"regexp"
)

type Card struct {
	ID             string `json:"id,omitempty"`
	Number         string `json:"number"`
	ExpirationDate string `json:"expiration_date"`
	HolderName     string `json:"holder_name"`
	CVV            string `json:"cvv"`
}

func (c *Card) IsValidNumber() bool {
	isValid, _ := luhn.IsValid(c.Number)
	return isValid
}

func (c *Card) IsValidDate() bool {
	isValid, _ := regexp.Match(`(0[1-9]|1[0-2])\/([0-9]{4}|[0-9]{2})`, []byte(c.ExpirationDate))
	return isValid
}

func (c *Card) IsValidCvv() bool {
	isValid, _ := regexp.Match(`([0-9]{4}|[0-9]{3})`, []byte(c.CVV))
	return isValid
}
