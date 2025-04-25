package models

type SearchParams struct {
	Query      string `json:"query"`       // For name/description search
	CategoryID *int   `json:"category_id"` // Optional category filter
	MinPrice   *int   `json:"min_price"`   // Optional minimum price
	MaxPrice   *int   `json:"max_price"`   // Optional maximum price
	Available  *bool  `json:"available"`   // Optional availability filter
}