package resp

import (
	"pixstall-search/domain/artwork/model"
	model2 "pixstall-search/domain/model"
	"strconv"
	"time"
)

type SearchArtworksResponse struct {
	Meta struct {
		Page Page `json:"page"`
	} `json:"meta"`
	Results []SearchArtworkResponseResult `json:"results"`
}

type SearchArtworkResponseResult struct {
	ID struct {
		Raw string `json:"raw"`
	} `json:"id"`
	OpenCommissionID struct {
		Raw string `json:"raw"`
	} `json:"open_commission_id"`

	ArtistID struct {
		Raw string `json:"raw"`
	} `json:"artist_id"`
	ArtistName struct {
		Raw string `json:"raw"`
	} `json:"artist_name"`
	ArtistProfilePath struct {
		Raw *string `json:"raw"`
	} `json:"artist_profile_path"`

	DayUsed struct {
		Raw float64 `json:"raw"`
	} `json:"day_used"`
	IsR18 struct {
		Raw string `json:"raw"`
	} `json:"is_r18"`
	Anonymous struct {
		Raw string `json:"raw"`
	} `json:"anonymous"`

	Path struct {
		Raw string `json:"raw"`
	} `json:"path"`
	Rating struct {
		Raw float64 `json:"raw"`
	} `json:"rating"`

	Title struct {
		Raw string `json:"raw"`
	} `json:"title"`
	TextContent struct {
		Raw string `json:"raw"`
	} `json:"text_content"`
	Views struct {
		Raw float64 `json:"raw"`
	} `json:"views"`
	FavorCount struct {
		Raw float64 `json:"raw"`
	} `json:"favor_count"`

	CreateTime struct {
		Raw time.Time `json:"raw"`
	} `json:"create_time"`
	StartTime struct {
		Raw time.Time `json:"raw"`
	} `json:"start_time"`
	CompletedTime struct {
		Raw time.Time `json:"raw"`
	} `json:"completed_time"`
	LastUpdatedTime struct {
		Raw time.Time `json:"raw"`
	} `json:"last_update_time"`
	State struct {
		Raw model.ArtworkState `json:"raw"`
	} `json:"state"`
}

func (r SearchArtworksResponse) ToDomainResult() model.GetArtworksResult {
	var artworks []model.Artwork
	for _, v := range r.Results {
		isR18, err := strconv.ParseBool(v.IsR18.Raw)
		if err != nil {
			continue
		}
		anonymous, err := strconv.ParseBool(v.Anonymous.Raw)
		if err != nil {
			continue
		}
		artworks = append(artworks, model.Artwork{
			ID:                v.ID.Raw,
			OpenCommissionID:  v.OpenCommissionID.Raw,
			ArtistID:          v.ArtistID.Raw,
			ArtistName:        v.ArtistName.Raw,
			ArtistProfilePath: v.ArtistProfilePath.Raw,
			DayUsed:           time.Duration(v.DayUsed.Raw),
			IsR18:             isR18,
			Anonymous:       anonymous,
			Path:            v.Path.Raw,
			Rating:          int(v.Rating.Raw),
			Title:           v.Title.Raw,
			TextContent:     v.TextContent.Raw,
			Views:           int(v.Views.Raw),
			FavorCount:      int(v.FavorCount.Raw),
			CreateTime:      v.CreateTime.Raw,
			StartTime:       v.StartTime.Raw,
			CompletedTime:   v.CompletedTime.Raw,
			LastUpdatedTime: v.LastUpdatedTime.Raw,
			State:           v.State.Raw,
		})
	}

	return model.GetArtworksResult{
		Page: model2.Page{
			Current:     r.Meta.Page.Current,
			TotalPage:   r.Meta.Page.TotalPage,
			TotalResult: r.Meta.Page.TotalResult,
			Size:        r.Meta.Page.Size,
		},
		Artworks: artworks,
	}
}
