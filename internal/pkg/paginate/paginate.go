package paginate

import (
	"encoding/json"

	"github.com/acoshift/paginate"
)

const (
	defaultPerPage = 30
)

// Paginate type
type Paginate struct {
	paginate.Paginate

	page    int64
	perPage int64
}

// MarshalJSON marshals paginate into json
func (p Paginate) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Page    int64 `json:"page"`
		PerPage int64 `json:"perPage"`
		MaxPage int64 `json:"maxPage"`
		Total   int64 `json:"total"`
	}{p.Page(), p.PerPage(), p.MaxPage(), p.Count()})
}

// UnmarshalJSON unmarshals paginate from json
func (p *Paginate) UnmarshalJSON(b []byte) error {
	p.Paginate = paginate.Paginate{}

	var x struct {
		Page    int64 `json:"page"`
		PerPage int64 `json:"perPage"`
	}
	err := json.Unmarshal(b, &x)
	if err != nil {
		return err
	}

	p.page, p.perPage = x.Page, x.PerPage

	return nil
}

// SetCount sets item count
func (p *Paginate) SetCount(cnt int64) {
	if p.page <= 0 {
		p.page = 1
	}
	if p.perPage <= 0 {
		p.perPage = defaultPerPage
	}

	p.Paginate = *paginate.New(p.page, p.perPage, cnt)
}

// CountFrom sets count using result from f
func (p *Paginate) CountFrom(f func() (cnt int64, err error)) error {
	cnt, err := f()
	if err != nil {
		return err
	}
	p.SetCount(cnt)
	return nil
}
