package utils

type ctxKey string

const (
	RequesterCtxKey ctxKey = "requesterCtxKey"
)

type ContextData struct {
	UserID    string
	UserEmail string
	UserName  string
}
