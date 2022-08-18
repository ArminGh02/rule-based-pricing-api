package request

import (
	"math/big"
	"pricingapi/pkg/model/usertype"
)

type Apply struct {
	UserType usertype.UserType `json:"user_type" bson:"user_type"`
	Price    *big.Int          `json:"price" bson:"price"`
}
