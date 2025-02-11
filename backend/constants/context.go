package env

type contextKey string

const (
	// ContextKeyPageNumber is the key used to set pagination page number in context
	ContextKeyPageNumber contextKey = "_ctx.middlewares.key-page-number_"

	// ContextKeyPageSize is the key used to set pagination per_page value in context
	ContextKeyPageSize contextKey = "_ctx.middlewares.key-page-size_"
)
