package paginator

type Paginator struct {
	Page       int `json:"page"`
	PageSize   int `json:"page_size"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

func New(page, pageSize int) *Paginator {
	return &Paginator{
		Page:     page,
		PageSize: pageSize,
	}
}

func NewWithTotal(page, pageSize, total int) *Paginator {
	var totalPages int
	if pageSize > 0 {
		totalPages = (total + pageSize - 1) / pageSize
	}
	return &Paginator{
		Page:       page,
		PageSize:   pageSize,
		Total:      total,
		TotalPages: totalPages,
	}
}

func (p *Paginator) SetTotal(total int) *Paginator {
	p.Total = total
	if p.PageSize > 0 {
		p.TotalPages = (total + p.PageSize - 1) / p.PageSize
	}
	return p
}

func (p *Paginator) SetPage(page int) *Paginator {
	p.Page = page
	return p
}

func (p *Paginator) SetPageSize(pageSize int) *Paginator {
	p.PageSize = pageSize
	return p
}

func (p *Paginator) GetOffset() int {
	return (p.Page - 1) * p.PageSize
}

// Paginate paginates the slice.
func Paginate[T any](items []T, p *Paginator) []T {
	if len(items) == 0 {
		return items
	}

	offset := p.GetOffset()
	if offset < 0 || offset >= len(items) {
		return nil
	}
	offsetEnd := offset + p.PageSize
	if offsetEnd > len(items) {
		offsetEnd = len(items)
	}
	return items[offset:offsetEnd]
}
