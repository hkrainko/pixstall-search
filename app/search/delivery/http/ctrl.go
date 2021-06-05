package http

import (
	"github.com/gin-gonic/gin"
	"pixstall-search/app/error/http"
	error2 "pixstall-search/domain/error"
	"pixstall-search/domain/search"
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
	searchType, exist := c.GetQuery("t")
	if !exist || (searchType != "all" &&  searchType != "open-commissions" && searchType != "artists" && searchType != "artworks"){
		c.JSON(http.NewErrorResponse(error2.BadRequestError))
		return
	}
}