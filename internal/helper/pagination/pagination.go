package pagination

import (
	_const "github.com/vekaputra/tiger-kittens/internal/const"
	"github.com/vekaputra/tiger-kittens/internal/model"
)

func SQLOffset(page, perPage uint64) uint64 {
	if page == 0 {
		page = 1
	}
	return (page - 1) * perPage
}

func TotalPage(totalItem, perPage uint64) uint64 {
	return (totalItem + perPage - 1) / perPage
}

func DefaultPagination(page model.PaginationRequest) model.PaginationRequest {
	if page.Page == 0 {
		page.Page = 1
	}
	if page.PerPage == 0 {
		page.PerPage = _const.DefaultPerPage
	}
	return page
}
