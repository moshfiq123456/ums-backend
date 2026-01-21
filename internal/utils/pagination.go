package utils

type Pagination struct {
	Page int `form:"page"`
	Size int `form:"size"`
}

func (p *Pagination) Normalize() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Size <= 0 {
		p.Size = 20
	}

	if p.Size > 100 {
		p.Size = 100
	}
}

