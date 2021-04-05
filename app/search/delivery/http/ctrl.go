package http

import (
	"github.com/gin-gonic/gin"
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

func (s SearchController) Search(ctx *gin.Context) {

}