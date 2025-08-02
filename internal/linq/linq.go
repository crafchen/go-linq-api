package linq

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type QueryBuilder struct {
	db        *gorm.DB
	tableName string
	joins     []string
	wheres    []string
	args      []interface{}
	selectStr string
	orderStr  string
	groupStr  string
	havingStr string
	limitNum  int
	offsetNum int
	distinct  bool
}

// ---------------- INIT ----------------

// Build gom tất cả điều kiện trước khi thực thi
func (q *QueryBuilder) Build() *gorm.DB {
	finalDB := q.db

	// DISTINCT
	if q.distinct {
		finalDB = finalDB.Distinct(q.selectStr)
	}

	// SELECT
	if q.selectStr != "" && !q.distinct {
		finalDB = finalDB.Select(q.selectStr)
	}

	// JOIN
	for _, join := range q.joins {
		finalDB = finalDB.Joins(join)
	}

	// WHERE
	if len(q.wheres) > 0 {
		finalDB = finalDB.Where(strings.Join(q.wheres, " AND "), q.args...)
	}

	// GROUP BY + HAVING
	if q.groupStr != "" {
		finalDB = finalDB.Group(q.groupStr)
	}
	if q.havingStr != "" {
		finalDB = finalDB.Having(q.havingStr)
	}

	// ORDER
	if q.orderStr != "" {
		finalDB = finalDB.Order(q.orderStr)
	}

	// LIMIT & OFFSET
	if q.limitNum >= 0 {
		finalDB = finalDB.Limit(q.limitNum)
	}
	if q.offsetNum >= 0 {
		finalDB = finalDB.Offset(q.offsetNum)
	}

	return finalDB
}

// Khởi tạo từ model GORM
func From(db *gorm.DB, model interface{}) *QueryBuilder {
	stmt := &gorm.Statement{DB: db}
	_ = stmt.Parse(model) // parse model để lấy tên bảng

	return &QueryBuilder{
		db:        db.Table(stmt.Schema.Table),
		tableName: stmt.Schema.Table,
		limitNum:  -1,
		offsetNum: -1,
	}
}

// ---------------- LINQ-LIKE OPERATIONS ----------------

func (q *QueryBuilder) Select(selectStr string) *QueryBuilder {
	q.selectStr = selectStr
	return q
}

func (q *QueryBuilder) Distinct() *QueryBuilder {
	q.distinct = true
	return q
}

func (q *QueryBuilder) Where(condition string, args ...interface{}) *QueryBuilder {
	q.wheres = append(q.wheres, condition)
	q.args = append(q.args, args...)
	return q
}

func (q *QueryBuilder) OrderBy(order string) *QueryBuilder {
	q.orderStr = order
	return q
}

func (q *QueryBuilder) GroupBy(group string) *QueryBuilder {
	q.groupStr = group
	return q
}

func (q *QueryBuilder) Having(having string) *QueryBuilder {
	q.havingStr = having
	return q
}

func (q *QueryBuilder) Limit(limit int) *QueryBuilder {
	q.limitNum = limit
	return q
}

func (q *QueryBuilder) Offset(offset int) *QueryBuilder {
	q.offsetNum = offset
	return q
}

// JOIN hỗ trợ nhiều bảng
func (q *QueryBuilder) InnerJoin(table string, on string) *QueryBuilder {
	// MSSQL cần ON ngay sau bảng
	q.joins = append(q.joins, fmt.Sprintf("INNER JOIN %s ON %s", table, on))
	return q
}

func (q *QueryBuilder) LeftJoin(table string, on string) *QueryBuilder {
	q.joins = append(q.joins, fmt.Sprintf("LEFT JOIN %s ON %s", table, on))
	return q
}

func (q *QueryBuilder) RightJoin(table string, on string) *QueryBuilder {
	q.joins = append(q.joins, fmt.Sprintf("RIGHT JOIN %s ON %s", table, on))
	return q
}

// ---------------- EXECUTION ----------------

// ToList: thực thi query và bind vào out
func (q *QueryBuilder) ToList(out interface{}) error {
	finalDB := q.build()
	return finalDB.Find(out).Error
}

// FirstOrDefault: lấy 1 bản ghi đầu tiên
func (q *QueryBuilder) FirstOrDefault(out interface{}) error {
	finalDB := q.build().Limit(1)
	return finalDB.Find(out).Error
}

// Any: kiểm tra có record nào không
func (q *QueryBuilder) Any() (bool, error) {
	var count int64
	err := q.build().Limit(1).Count(&count).Error
	return count > 0, err
}

// Count: đếm số bản ghi
func (q *QueryBuilder) Count() (int64, error) {
	var count int64
	err := q.build().Count(&count).Error
	return count, err
}

// Sum: tính tổng của cột
func (q *QueryBuilder) Sum(column string) (float64, error) {
	var sum float64
	err := q.build().Select(fmt.Sprintf("SUM(%s)", column)).Scan(&sum).Error
	return sum, err
}

// Max: giá trị lớn nhất
func (q *QueryBuilder) Max(column string) (float64, error) {
	var max float64
	err := q.build().Select(fmt.Sprintf("MAX(%s)", column)).Scan(&max).Error
	return max, err
}

// Min: giá trị nhỏ nhất
func (q *QueryBuilder) Min(column string) (float64, error) {
	var min float64
	err := q.build().Select(fmt.Sprintf("MIN(%s)", column)).Scan(&min).Error
	return min, err
}

// ---------------- PRIVATE ----------------

// build() gom tất cả điều kiện trước khi thực thi
func (q *QueryBuilder) build() *gorm.DB {
	finalDB := q.db

	// DISTINCT + SELECT
	if q.distinct && q.selectStr != "" {
		finalDB = finalDB.Distinct(q.selectStr)
	} else if q.selectStr != "" {
		finalDB = finalDB.Select(q.selectStr)
	}

	// JOIN
	for _, join := range q.joins {
		finalDB = finalDB.Joins(join)
	}

	// WHERE
	if len(q.wheres) > 0 {
		finalDB = finalDB.Where(strings.Join(q.wheres, " AND "), q.args...)
	}

	// GROUP BY + HAVING
	if q.groupStr != "" {
		finalDB = finalDB.Group(q.groupStr)
	}
	if q.havingStr != "" {
		finalDB = finalDB.Having(q.havingStr)
	}

	// ORDER
	if q.orderStr != "" {
		finalDB = finalDB.Order(q.orderStr)
	}

	// LIMIT & OFFSET
	if q.limitNum >= 0 {
		finalDB = finalDB.Limit(q.limitNum)
	}
	if q.offsetNum >= 0 {
		finalDB = finalDB.Offset(q.offsetNum)
	}

	return finalDB
}
