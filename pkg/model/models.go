package model

import (
	"math/big"
	"pricingapi/pkg/model/ruletype"
	"pricingapi/pkg/model/usertype"
	"pricingapi/pkg/util"
)

type Rule struct {
	ID         int64             `json:"id" bson:"id"`
	Name       string            `json:"name" bson:"name"`
	Type       ruletype.RuleType `json:"type" bson:"type"`
	Action     Action            `json:"action" bson:"action"`
	Conditions []Condition       `json:"conditions" bson:"conditions"`
}

func (r *Rule) Apply(price *big.Int) *big.Int {
	res := new(big.Int).Set(price)
	switch r.Type {
	case ruletype.Discount:
		res.Set(
			new(big.Int).Sub(
				price,
				util.Min(
					new(big.Int).Add(
						price,
						new(big.Int).Div(
							new(big.Int).Mul(
								price,
								r.Action.PercentageDisplacementAmount,
							),
							big.NewInt(100),
						),
					),
					r.Action.MaximumDisplacementAmount,
				),
			),
		)
	case ruletype.Markup:
		res.Set(
			new(big.Int).Add(
				price,
				util.Min(
					new(big.Int).Add(
						price,
						new(big.Int).Div(
							new(big.Int).Mul(
								price,
								r.Action.PercentageDisplacementAmount,
							),
							big.NewInt(100),
						),
					),
					r.Action.MaximumDisplacementAmount,
				),
			),
		)
	}
	return res
}

type Action struct {
	ID                           int64    `json:"id" bson:"id"`
	FixedDisplacementAmount      *big.Int `json:"fixed_displacement_amount" bson:"fixed_displacement_amount"`
	PercentageDisplacementAmount *big.Int `json:"percentage_displacement_amount" bson:"percentage_displacement_amount"`
	MaximumDisplacementAmount    *big.Int `json:"maximum_displacement_amount" bson:"maximum_displacement_amount"`
}

type Condition struct {
	UserType usertype.UserType `json:"user_type" bson:"user_type"`
}

type ApplyRule struct {
	ID           int64    `json:"id" bson:"id"`
	Name         string   `json:"name" bson:"name"`
	Sequence     int      `json:"sequence" bson:"sequence"`
	OldPrice     *big.Int `json:"old_price" bson:"old_price"`
	NewPrice     *big.Int `json:"new_price" bson:"new_price"`
	Displacement *big.Int `json:"displacement" bson:"displacement"`
}
