package responseHandler

import (
	"github.com/PhillBeck/golang-odm"
)

type PaginationResponse struct {
	Docs    interface{} `json:"docs"`
	Page    int         `json:"page"`
	Pages   int         `json:"pages"`
	PerPage int         `json:"perPage"`
	Total   int         `json:"total"`
}

func MakePaginationResponse(docs interface{}, info *odm.PaginationInfo) PaginationResponse {
	if docs == nil {
		docs = []int{}
	}

	return PaginationResponse{
		Docs:    docs,
		Page:    info.CurrentPage,
		Pages:   info.NumPages,
		PerPage: info.RecordsPerPage,
		Total:   info.NumRecords}
}
