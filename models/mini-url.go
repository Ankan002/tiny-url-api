package models

import "github.com/kamva/mgm/v3"

type MiniUrl struct {
	mgm.DefaultModel `bson:",inline"`
	ShortId          string `json:"shortId" bson:"short_id" validate:"unique"`
	Url              string `json:"url"`
}
