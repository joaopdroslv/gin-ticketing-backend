package schemas

type ResponsePagination struct {
	Page      int64 `json:"page"`
	PageTotal int64 `json:"page_total"`
	Limit     int64 `json:"limit"`
	Total     int64 `json:"total"`
}
