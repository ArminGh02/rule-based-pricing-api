package api

import (
	"math/big"
	"pricingapi/pkg/db"
	"pricingapi/pkg/model"
	"pricingapi/pkg/model/usertype"
	"sync"
)

type API struct {
	db *db.DB
}

var (
	instance *API
	once     sync.Once
)

func Instance() *API {
	once.Do(func() {
		instance = &API{}
	})
	return instance
}

func (a *API) Init(db *db.DB) {
	a.db = db
}

func (a *API) Apply(price *big.Int, userType usertype.UserType) []*model.ApplyRule {
	var appliedRules []*model.ApplyRule
	rules := a.db.FilterRulesByUserType(userType)
	for i, r := range rules {
		newPrice := r.Apply(price)

		appliedRules = append(appliedRules, &model.ApplyRule{
			ID:           r.ID,
			Name:         r.Name,
			Sequence:     i,
			OldPrice:     price,
			NewPrice:     newPrice,
			Displacement: new(big.Int).Abs(new(big.Int).Sub(price, newPrice)),
		})
	}
	return appliedRules
}
