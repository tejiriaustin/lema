package env

type contextKey string

const (
	// ContextKeyPageNumber is the key used to set pagination page number in context
	ContextKeyPageNumber contextKey = "_ctx.middlewares.key-page-number_"

	// ContextKeyPerPageLimit is the key used to set pagination per_page value in context
	ContextKeyPerPageLimit contextKey = "_ctx.middlewares.per-page-limit_"
)
