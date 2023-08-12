package pagination

type RequestDTO struct {
	Size int `query:"size" form:"size,default=10" binding:"gt=0,lt=100"`
	Page int `query:"page" form:"page,default=1"  binding:"gt=0"`
}

func (p *RequestDTO) ToCmd() Query {
	return Query{p.Size, p.Page}
}
