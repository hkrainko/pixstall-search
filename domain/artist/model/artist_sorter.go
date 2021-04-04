package model

import "pixstall-search/domain/model"

type ArtistSorter struct {
	ArtistID               *model.SortOrder
	UserName               *model.SortOrder
	RegTime                *model.SortOrder
	YearOfDrawing          *model.SortOrder
	CommissionRequestCount *model.SortOrder
	CommissionAcceptCount  *model.SortOrder
	CommissionSuccessCount *model.SortOrder
	AvgRatings             *model.SortOrder
	LastRequestTime        *model.SortOrder
}
