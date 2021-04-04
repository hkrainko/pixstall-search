package model

import "time"

type OpenCommission struct {
	ID               string             `json:"id" bson:"id"`
	ArtistID         string             `json:"artistId" bson:"artistId"`
	Title            string             `json:"title" bson:"title"`
	Desc             string             `json:"desc" bson:"desc"`
	Price            Price              `json:"price" bson:"price"`
	DayNeed          DayNeed            `json:"dayNeed" bson:"dayNeed"`
	SampleImagePaths []string           `json:"sampleImagePaths" bson:"sampleImagePaths"`
	IsR18            bool               `json:"isR18" bson:"isR18"`
	AllowBePrivate   bool               `json:"allowBePrivate" bson:"allowBePrivate"`
	AllowAnonymous   bool               `json:"allowAnonymous" bson:"allowAnonymous"`
	State            OpenCommissionState `json:"state" bson:"state"`
	CreateTime       time.Time          `json:"createTime" bson:"createTime"`
	LastUpdatedTime  time.Time          `json:"lastUpdatedTime" bson:"lastUpdatedTime"`
}

type OpenCommissionState string

const (
	OpenCommissionStateActive  OpenCommissionState = "A"
	OpenCommissionStateHidden  OpenCommissionState = "H"
	OpenCommissionStateRemoved OpenCommissionState = "R"
)

type DayNeed struct {
	From int `json:"from" bson:"from"`
	To   int `json:"to" bson:"to"`
}
