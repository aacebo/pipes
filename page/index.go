package page

import (
	"net/http"
	"strconv"
)

type Page struct {
	Number int
	Size   int
}

func New(r *http.Request) *Page {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))

	if page == 0 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

	if pageSize == 0 {
		pageSize = 10
	}

	v := Page{
		Number: page,
		Size:   pageSize,
	}

	return &v
}

func (self *Page) GetOffset() int {
	return (self.Number - 1) * self.Size
}
