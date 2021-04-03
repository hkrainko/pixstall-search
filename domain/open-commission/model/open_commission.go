package model

import "time"

type OpenCommission struct {
	ID               string             `json:"id" bson:"id"`
	ArtistID         string             `json:"artistId" bson:"artistId"`
	Title            string             `json:"title" bson:"title"`
	Desc             string             `json:"desc" bson:"desc"`
	PriceAmount      float64            `json:"priceAmount" bson:"priceAmount"`
	PriceCurrency    float64            `json:"priceCurrent" bson:"priceCurrency"`
	DayNeedFrom      int                `json:"dayNeedFrom" bson:"dayNeedFrom"`
	DayNeedTo        int                `json:"dayNeedTo" bson:"dayNeedTo"`
	SampleImagePaths []string           `json:"sampleImagePaths" bson:"sampleImagePaths"`
	IsR18            bool               `json:"isR18" bson:"isR18"`
	AllowBePrivate   bool               `json:"allowBePrivate" bson:"allowBePrivate"`
	AllowAnonymous   bool               `json:"allowAnonymous" bson:"allowAnonymous"`
	State            OpenCommissionSate `json:"state" bson:"state"`
	CreateTime       time.Time          `json:"createTime" bson:"createTime"`
	LastUpdatedTime  time.Time          `json:"lastUpdatedTime" bson:"lastUpdatedTime"`
}

type OpenCommissionSate string

const (
	OpenCommissionStateActive  OpenCommissionSate = "A"
	OpenCommissionStateHidden  OpenCommissionSate = "H"
	OpenCommissionStateRemoved OpenCommissionSate = "R"
)
