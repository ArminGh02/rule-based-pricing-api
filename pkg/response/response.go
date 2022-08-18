package response

import "pricingapi/pkg/model"

type Apply struct {
	Applied      bool               `json:"applied" bson:"applied"`
	AppliedRules []*model.ApplyRule `json:"applied_rules" bson:"applied_rules"`
}
