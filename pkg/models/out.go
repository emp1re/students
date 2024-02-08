package models

import (
	"time"

	db "github.com/emp1re/students/pkg/db/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

// goverter:converter
// goverter:extend PgTypeToInt8
// goverter:extend PgTypeToText
// goverter:extend PgTypeTimestamptz
// goverter:extend PgTypeToBool
type Converter interface {
	ConvertItems(source []db.Student) []OutStudent
	Convert(source db.Student) OutStudent
}

func PgTypeToInt8(i pgtype.Int8) int64 {
	val, err := i.Value()

	if err != nil {

		return 0
	}
	return val.(int64)
}
func PgTypeToBool(i pgtype.Bool) bool {
	val, err := i.Value()
	if err != nil {
		return false
	}
	return val.(bool)
}
func PgTypeTimestamptz(i pgtype.Timestamptz) time.Time {
	val, err := i.Value()
	if err != nil {
		return time.Time{}
	}
	return val.(time.Time)
}
func PgTypeToText(i pgtype.Text) string {
	val, err := i.Value()
	if err != nil {
		return ""
	}
	return val.(string)
}

type OutAddress struct {
	AddressID int64  `json:"address_id"`
	Street    string `json:"street"`
	City      string `json:"city"`
	Planet    string `json:"planet"`
	Phone     string `json:"phone"`
}
type OutStudent struct {
	ID             int32     `json:"id"`
	StudentID      string    `json:"student_id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Age            int64     `json:"age"`
	Email          string    `json:"email"`
	Gender         string    `json:"gender"`
	FavouriteColor string    `json:"favourite_color"`
	StudentAddress string    `json:"student_address"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Deleted        bool      `json:"deleted"`
}
