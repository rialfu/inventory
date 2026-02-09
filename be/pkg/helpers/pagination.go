package helpers

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Pagination struct {
	Page  int
	Limit int
	Sort  string
	Order string
}

type ColumnInfo struct {
	IsNumeric bool
}

var schemaCache = sync.Map{}

func (p *Pagination) Normalize() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Limit <= 0 || p.Limit > 100 {
		p.Limit = 10
	}
	if p.Order != "asc" && p.Order != "desc" {
		p.Order = "asc"
	}
}

func GetModelColumns(db *gorm.DB, model any) map[string]bool {
	stmt := &gorm.Statement{DB: db}
	_ = stmt.Parse(model)

	cols := make(map[string]bool)
	for _, field := range stmt.Schema.Fields {
		cols[field.DBName] = true
	}
	return cols
}

func ApplyPagination(
	db *gorm.DB,
	model any,
	queryParams map[string][]string,
) (*gorm.DB, Pagination) {
	modelType := reflect.TypeOf(model).String()

	// Cek apakah cache sudah ada
	val, exists := schemaCache.Load(modelType)
	var columnCache map[string]bool
	if !exists {
		// Jika belum ada, baru hitung (Parse Schema)
		stmt := &gorm.Statement{DB: db}
		stmt.Parse(model)

		columnCache = make(map[string]bool)
		for name, field := range stmt.Schema.FieldsByDBName {
			dt := string(field.DataType)
			columnCache[name] = (dt == "int" || dt == "uint" || dt == "float")
		}
		// Simpan ke cache
		schemaCache.Store(modelType, columnCache)
	} else {
		columnCache = val.(map[string]bool)
	}

	p := Pagination{
		Page:  parseInt(queryParams["page"], 1),
		Limit: parseInt(queryParams["limit"], 10),
		Sort:  parseString(queryParams["sort"]),
		Order: parseString(queryParams["order"]),
	}
	p.Normalize()
	if p.Limit > 100 {
		p.Limit = 100
	}
	validColumns := GetModelColumns(db, model)

	// ===== Sorting =====
	if p.Sort != "" && validColumns[p.Sort] {
		db = db.Order(clause.OrderByColumn{
			Column: clause.Column{Name: p.Sort},
			Desc:   p.Order == "desc",
		})
	}

	// ===== Search / Filter =====
	for key, values := range queryParams {
		if !validColumns[key] {
			continue
		}
		if len(values) == 0 || values[0] == "" {
			continue
		}
		isNumeric, exists := columnCache[key]
		if !exists {
			continue
		}
		val := values[0]
		if val == "" {
			continue
		}
		if isNumeric {
			// Logika untuk angka: Mendukung operator <, >, <=, >=, =, !=
			op, num, ok := parseNumericFilter(val)
			if ok {
				// Jika ada operator (misal: ">10")
				db = db.Where(fmt.Sprintf("%s %s ?", key, op), num)
			} else {
				// Jika hanya angka tanpa operator (misal: "10"), defaultnya adalah "="
				if _, err := strconv.ParseFloat(val, 64); err == nil {
					db = db.Where(fmt.Sprintf("%s = ?", key), val)
				}
			}
		} else {
			db = db.Where(
				fmt.Sprintf("%s ILIKE ?", key),
				"%"+values[0]+"%",
			)
		}

	}

	offset := (p.Page - 1) * p.Limit
	db = db.Offset(offset).Limit(p.Limit)

	return db, p
}
func isNumericColumn(db *gorm.DB, model any, columnName string) bool {
	stmt := &gorm.Statement{DB: db}
	_ = stmt.Parse(model) // Mengisi schema berdasarkan model

	field, ok := stmt.Schema.FieldsByDBName[columnName]
	if !ok {
		return false
	}

	switch string(field.DataType) {
	case "int", "uint", "float":
		return true
	default:
		return false
	}
}
func parseNumericFilter(input string) (string, string, bool) {
	// Daftar operator yang didukung (Urutan penting: yang 2 karakter harus di atas)
	operators := []string{">=", "<=", "!=", ">", "<", "="}

	for _, op := range operators {
		if strings.HasPrefix(input, op) {
			value := strings.TrimPrefix(input, op)
			// Validasi apakah setelah operator adalah angka/decimal yang valid
			if _, err := strconv.ParseFloat(value, 64); err == nil {
				return op, value, true
			}
		}
	}
	return "", "", false
}

type PaginateData[T any] struct {
	Data  []T   `json:"data"`
	Limit int   `json:"limit"`
	Page  int   `json:"page"`
	Total int64 `json:"total"`
}
