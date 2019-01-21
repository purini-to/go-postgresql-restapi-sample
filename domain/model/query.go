package model

var (
	// DefaultLimit is default limit.
	DefaultLimit = uint32(1000)
	// DefaultOffset is default offset.
	DefaultOffset = uint32(0)
	// DefaultSort is default sort.
	DefaultSort = "updated_at desc, created_at desc, id"
)

// Query is query parameters.
type Query struct {
	Limit  string `json:"limit" valid:"int,range(1|1000)"`
	Offset string `json:"offset" valid:"int,range(0|1000000000)"`
	Sort   string `json:"-"`
}
