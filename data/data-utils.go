package data

import (
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

func (dal QuestionsDal) Paginate(r *http.Request, maxPageSize int, defaultPageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		q := r.URL.Query()
		page, _ := strconv.Atoi(q.Get("page"))
		if page == 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(q.Get("page_size"))
		switch {
		case pageSize > maxPageSize:
			pageSize = maxPageSize
		case pageSize <= 0:
			pageSize = defaultPageSize
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
