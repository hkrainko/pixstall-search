package req

type GetSuggestionsRequest struct {
	Query string `json:"query"`
	// TODO: Add fields and size
}

type GetSuggestionsResponse struct {

}

func NewGetSuggestionsRequest(query string) GetSuggestionsRequest {
	return GetSuggestionsRequest{
		Query: query,
	}
}