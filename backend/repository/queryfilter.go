package repository

type Query struct {
	query string
	args  []interface{}
}

func NewQueryFilter() *Query {
	return &Query{}
}

func (f *Query) Where(query string, args ...interface{}) *Query {
	f.query += query
	f.args = append(f.args, args)
	return f
}

func (f *Query) Raw(query string, args ...interface{}) *Query {
	f.query += query
	f.args = append(f.args, args...)
	return f
}
