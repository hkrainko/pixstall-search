package resp

import (
	"pixstall-search/domain/artist/model"
	model2 "pixstall-search/domain/model"
	"time"
)

type SearchArtistsResponse struct {
	Meta struct {
		Page Page `json:"page"`
	} `json:"meta"`
	Results []SearchArtistsResponseResult `json:"results"`
}

type SearchArtistsResponseResult struct {
	ID struct {
		Raw string `json:"raw"`
	} `json:"id"`

	// User
	UserID struct {
		Raw string `json:"raw"`
	} `json:"user_id"`
	UserName struct {
		Raw string `json:"raw"`
	} `json:"user_name"`
	ProfilePath struct {
		Raw *string `json:"raw,omitempty"`
	} `json:"profile_path"`
	State struct {
		Raw model.UserState `json:"raw"`
	} `json:"state"`
	RegTime struct {
		Raw time.Time `json:"raw"`
	} `json:"reg_time"`
	LastUpdatedTime struct {
		Raw time.Time `json:"raw"`
	} `json:"last_updated_time"`

	ArtistID struct {
		Raw string `json:"raw"`
	} `json:"artist_id"`
	PaymentMethods struct {
		Raw []string `json:"raw"`
	} `json:"payment_methods"`

	// ArtistIntro
	YearOfDrawing struct {
		Raw float64 `json:"raw"`
	} `json:"year_of_drawing"`
	ArtTypes struct {
		Raw []string `json:"raw"`
	} `json:"art_types"`

	// ArtistBoard
	Desc struct {
		Raw string `json:"raw"`
	} `json:"desc"`

	// CommissionDetails
	CommissionRequestCount struct {
		Raw float64 `json:"raw"`
	} `json:"commission_request_count"`
	CommissionAcceptCount struct {
		Raw float64 `json:"raw"`
	} `json:"commission_accept_count"`
	CommissionSuccessCount struct {
		Raw float64 `json:"raw"`
	} `json:"commission_success_count"`
	AvgRatings struct {
		Raw *float64 `json:"raw,omitempty"`
	} `json:"avg_ratings"`
	LastRequestTime struct {
		Raw *time.Time `json:"raw,omitempty"`
	} `json:"last_request_time"`
}

func (r SearchArtistsResponse) ToDomainResult() model.GetArtistsResult {

	var artists []model.Artist
	for _, v := range r.Results {
		artists = append(artists, model.Artist{
			User: model.User{
				UserID:          v.UserID.Raw,
				UserName:        v.UserName.Raw,
				ProfilePath:     v.ProfilePath.Raw,
				State:           v.State.Raw,
				RegTime:         v.RegTime.Raw,
				LastUpdatedTime: v.LastUpdatedTime.Raw,
			},
			ArtistID: v.ArtistID.Raw,
			ArtistIntro: model.ArtistIntro{
				YearOfDrawing: int(v.YearOfDrawing.Raw),
				ArtTypes:      v.ArtTypes.Raw,
			},
			ArtistBoard: model.ArtistBoard{
				Desc: v.Desc.Raw,
			},
			PaymentMethods: v.PaymentMethods.Raw,
			CommissionDetails: model.CommissionDetails{
				CommissionRequestCount: int(v.CommissionRequestCount.Raw),
				CommissionAcceptCount:  int(v.CommissionAcceptCount.Raw),
				CommissionSuccessCount: int(v.CommissionSuccessCount.Raw),
				AvgRatings:             v.AvgRatings.Raw,
				LastRequestTime:        v.LastRequestTime.Raw,
			},
		})
	}

	return model.GetArtistsResult{
		Page: model2.Page{
			Current:     r.Meta.Page.Current,
			TotalPage:   r.Meta.Page.TotalPage,
			TotalResult: r.Meta.Page.TotalResult,
			Size:        r.Meta.Page.Size,
		},
		Artists: artists,
	}
}
