package model

var (
	defaultLimit  = uint32(1000)
	defaultOffset = uint32(0)
	defaultSort   = "updated_at desc, created_at desc, id"
)

// Query is query parameters.
type Query struct {
	Limit  *uint32 `json:"limit"`
	Offset *uint32 `json:"offset"`
	Sort   *string `json:"sort"`
}

// DefaultQuery create default query.
func DefaultQuery() *Query {
	return &Query{
		Limit:  &defaultLimit,
		Offset: &defaultOffset,
		Sort:   &defaultSort,
	}
}
