package model

type PaginationRequest struct {
	Page    uint64 `query:"page"`
	PerPage uint64 `query:"per_page"`
}

type PaginationResponse struct {
	Page      uint64 `json:"page"`
	PerPage   uint64 `json:"per_page"`
	TotalPage uint64 `json:"total_page"`
	TotalItem uint64 `json:"total_item"`
}
