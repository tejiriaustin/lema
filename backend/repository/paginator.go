package repository

import "math"

type Paginator struct {
	CurrentPage int64 `json:"current_page"`
	PerPage     int64 `json:"per_page"`
	TotalRows   int64 `json:"total_rows"`
	TotalPages  int64 `json:"total_pages"`
	Offset      int64 `json:"offset"`
	NextPage    int64 `json:"next_page"`
	PrevPage    int64 `json:"prev_page"`
}

func newPaginator(page, perPage int64) *Paginator {
	return &Paginator{
		CurrentPage: page,
		PerPage:     perPage,
	}
}

func (p *Paginator) setOffset() {
	p.Offset = (p.CurrentPage - 1) * p.PerPage
}

func (p *Paginator) setTotalPages() {
	p.TotalPages = int64(math.Ceil(float64(p.TotalRows) / float64(p.PerPage)))
}

func (p *Paginator) setPrevPage() {
	if p.CurrentPage > 1 {
		p.PrevPage = p.CurrentPage - 1
	}
}

func (p *Paginator) setNextPage() {
	if p.CurrentPage < p.TotalPages {
		p.NextPage = p.CurrentPage + 1
	}
}
