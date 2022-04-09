package controller

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	DefaultPageSize = 6
	MaxPageSize     = 100
	PageVar         = "page"
	PageSizeVar     = "pageSize"
)

type Page struct {
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
	Total    int         `json:"total"`
	Pages    int         `json:"pages"`
	Data     interface{} `json:"data"`
}

func New(page, pageSize int, total int, data interface{}) *Page {
	if pageSize == 0 {
		pageSize = DefaultPageSize
	}
	if pageSize > MaxPageSize {
		pageSize = MaxPageSize
	}
	pages := (total + pageSize - 1) / pageSize
	if page > pages {
		page = pages
	}
	if page <= 0 {
		page = 1
	}
	prevPage := page - 1
	if prevPage <= 0 {
		prevPage = 1
	}
	nextPage := page + 1
	if nextPage > pages {
		nextPage = pages
	}
	return &Page{
		Page:     page,
		PageSize: pageSize,
		Total:    total,
		Pages:    pages,
		Data:     data,
	}
}

func NewFromRequest(c *gin.Context, data interface{}, count int) *Page {
	page, _ := strconv.Atoi(c.Query(PageVar))
	pageSize, _ := strconv.Atoi(c.Query(PageSizeVar))
	return New(page, pageSize, count, data)
}

func GetPaginationParameterFromRequest(c *gin.Context) (int, int) {
	page, _ := strconv.Atoi(c.Query(PageVar))
	pageSize, _ := strconv.Atoi(c.Query(PageSizeVar))
	if pageSize == 0 {
		pageSize = DefaultPageSize
	}
	if pageSize > MaxPageSize {
		pageSize = MaxPageSize
	}
	return page, pageSize
}

func (p *Page) BuildLinkHeader(baseUrl string, defaultPageSize int) string {

	if p.PageSize == 0 {
		p.PageSize = defaultPageSize
	}
	if p.PageSize > MaxPageSize {
		p.PageSize = MaxPageSize
	}
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Pages <= 0 {
		p.Pages = 1
	}
	if p.Pages > 1 {
		var links []string
		if p.Page > 1 {
			links = append(links, getLink(baseUrl, p.Page-1, p.PageSize))
		}
		if p.Page < p.Pages {
			links = append(links, getLink(baseUrl, p.Page+1, p.PageSize))
		}
		return strings.Join(links, ", ")
	}
	return ""

}

func getLink(baseUrl string, page int, pageSize int) string {
	return "<" + baseUrl + "?" + PageVar + "=" + strconv.Itoa(page) + "&" + PageSizeVar + "=" + strconv.Itoa(pageSize) + ">; rel=\"next\""
}
