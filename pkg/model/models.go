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

func (r *Rule) Apply(price *big.Float) *big.Float {
	res := new(big.Float).Set(price)
	switch r.Type {
	case ruletype.Discount:
		res.Set(
			new(big.Float).Sub(
				price,
				util.Min(
					new(big.Float).Add(
						price,
						new(big.Float).Quo(
							new(big.Float).Mul(
								price,
								r.Action.PercentageDisplacementAmount,
							),
							big.NewFloat(100),
						),
					),
					r.Action.MaximumDisplacementAmount,
				),
			),
		)
	case ruletype.Markup:
		res.Set(
			new(big.Float).Add(
				price,
				util.Min(
					new(big.Float).Add(
						price,
						new(big.Float).Quo(
							new(big.Float).Mul(
								price,
								r.Action.PercentageDisplacementAmount,
							),
							big.NewFloat(100),
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
	ID                           int64      `json:"id" bson:"id"`
	FixedDisplacementAmount      *big.Float `json:"fixed_displacement_amount" bson:"fixed_displacement_amount"`
	PercentageDisplacementAmount *big.Float `json:"percentage_displacement_amount" bson:"percentage_displacement_amount"`
	MaximumDisplacementAmount    *big.Float `json:"maximum_displacement_amount" bson:"maximum_displacement_amount"`
}

type Condition struct {
	UserType usertype.UserType `json:"user_type" bson:"user_type"`
}

type ApplyRule struct {
	ID           int64      `json:"id" bson:"id"`
	Name         string     `json:"name" bson:"name"`
	Sequence     int        `json:"sequence" bson:"sequence"`
	OldPrice     *big.Float `json:"old_price" bson:"old_price"`
	NewPrice     *big.Float `json:"new_price" bson:"new_price"`
	Displacement *big.Float `json:"displacement" bson:"displacement"`
}
