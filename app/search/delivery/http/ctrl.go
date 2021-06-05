package http

import (
	"github.com/gin-gonic/gin"
	http2 "net/http"
	"pixstall-search/app/error/http"
	search_open_commissions "pixstall-search/app/search/delivery/http/resp/search-open-commissions"
	error2 "pixstall-search/domain/error"
	model3 "pixstall-search/domain/model"
	model2 "pixstall-search/domain/open-commission/model"
	"pixstall-search/domain/search"
	"pixstall-search/domain/search/model"
	"strconv"
)

type SearchController struct {
	useCase search.UseCase
}

func NewSearchController(useCase search.UseCase) SearchController {
	return SearchController{
		useCase: useCase,
	}
}

func (s SearchController) Search(c *gin.Context) {
	t, exist := c.GetQuery("t")
	if !exist {
		c.JSON(http.NewErrorResponse(error2.BadRequestError))
		return
	}
	searchType := model.NewSearchType(t)
	if searchType == nil {
		c.JSON(http.NewErrorResponse(error2.BadRequestError))
		return
	}

	switch *searchType {
	case model.AllSearchType:
		s.searchAll(c)
	case model.OpenCommissionsSearchType:
		s.searchOpenCommissions(c)
	case model.ArtistsSearchType:
		s.searchArtists(c)
	case model.ArtworksSearchType:
		s.searchArtworks(c)
	}
}

func (s SearchController) searchAll(c *gin.Context) {
	_ = c.Query("s")

}

func (s SearchController) searchOpenCommissions(c *gin.Context) {
	searchText := c.Query("s")

	priceAmount := model3.GetFloatRange(getFloatFromQuery("price.from", c), getFloatFromQuery("price.to", c))
	priceCurrency := model2.NewCurrency(c.Query("currency"))
	if priceAmount != nil && priceCurrency == nil {
		c.AbortWithStatusJSON(http2.StatusBadRequest, error2.BadRequestError)
		return
	}
	dayNeed := model3.GetIntRange(getIntFromQuery("day-need.from", c), getIntFromQuery("day-need.to", c))

	isR18 := getBool("is_r18", c)
	allowBePrivate := getBool("allow-be-private", c)
	allowAnonymous := getBool("allow-anonymous", c)

	pageCurrent := c.Query("page.current")
	intPageCurrent, err := strconv.Atoi(pageCurrent)
	if err != nil {
		c.AbortWithStatusJSON(http2.StatusBadRequest, error2.BadRequestError)
		return
	}
	pageSize := c.Query("page.size")
	intPageSize, err := strconv.Atoi(pageSize)
	if err != nil {
		c.AbortWithStatusJSON(http2.StatusBadRequest, error2.BadRequestError)
		return
	}

	filter := model2.OpenCommissionFilter{
		State:          nil,
		PriceAmount:    priceAmount,
		PriceCurrency:  priceCurrency,
		DayNeed:        dayNeed,
		IsR18:          isR18,
		AllowBePrivate: allowBePrivate,
		AllowAnonymous: allowAnonymous,
		PageFilter: model3.PageFilter{
			Current: intPageCurrent,
			Size:    intPageSize,
		},
	}
	sorter := GetOpenCommissionSorter(c.Query("sort"))
	if sorter == nil {
		c.AbortWithStatusJSON(http2.StatusBadRequest, error2.BadRequestError)
		return
	}

	result, err := s.useCase.SearchOpenCommissions(c, searchText, filter, *sorter)
	if err != nil {
		c.JSON(http.NewErrorResponse(err))
		return
	}
	c.JSON(http2.StatusOK, search_open_commissions.NewResponse(*result))
}

func (s SearchController) searchArtists(c *gin.Context) {
	//s, exist := c.GetQuery("s")
}

func (s SearchController) searchArtworks(c *gin.Context) {
	//s, exist := c.GetQuery("s")
}

// Private

func getBool(q string, c *gin.Context) *bool {
	str, exist := c.GetQuery(q)
	if !exist {
		return nil
	}
	result := str == "true"
	return &result
}

func getFloatFromQuery(q string, c *gin.Context) *float64 {
	str, exist := c.GetQuery(q)
	if !exist {
		return nil
	}
	result, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return nil
	}
	return &result
}

func getIntFromQuery(q string, c *gin.Context) *int {
	str, exist := c.GetQuery(q)
	if !exist {
		return nil
	}
	result, err := strconv.Atoi(str)
	if err != nil {
		return nil
	}
	return &result
}

func GetOpenCommissionSorter(str string) *model2.OpenCommissionSorter {
	sorter := model2.OpenCommissionSorter{}
	symbol := str[:1]
	if symbol == "-" {
		switch str[1:len(str)] {
		case "artist-id":
			v := model3.SortOrderDescending
			sorter.ArtistID = &v
			return &sorter
		case "price":
			v := model3.SortOrderDescending
			sorter.Price = &v
			return &sorter
		case "day-need-from":
			v := model3.SortOrderDescending
			sorter.DayNeedFrom = &v
			return &sorter
		case "day-need-to":
			v := model3.SortOrderDescending
			sorter.DayNeedTo = &v
			return &sorter
		case "create-time":
			v := model3.SortOrderDescending
			sorter.CreateTime = &v
			return &sorter
		case "last-updated-time":
			v := model3.SortOrderDescending
			sorter.LastUpdatedTime = &v
			return &sorter
		}
	} else {
		switch str {
		case "artist-id":
			v := model3.SortOrderAscending
			sorter.ArtistID = &v
			return &sorter
		case "price":
			v := model3.SortOrderAscending
			sorter.Price = &v
			return &sorter
		case "day-need-from":
			v := model3.SortOrderAscending
			sorter.DayNeedFrom = &v
			return &sorter
		case "day-need-to":
			v := model3.SortOrderAscending
			sorter.DayNeedTo = &v
			return &sorter
		case "create-time":
			v := model3.SortOrderAscending
			sorter.CreateTime = &v
			return &sorter
		case "last-updated-time":
			v := model3.SortOrderAscending
			sorter.LastUpdatedTime = &v
			return &sorter
		}
	}
	return nil
}
