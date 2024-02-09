package models

import (
	"time"

	db "github.com/emp1re/students/pkg/db/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

// goverter:converter
// goverter:extend PgTypeToInt4
// goverter:extend PgTypeToText
// goverter:extend PgTypeTimestamptz
// goverter:extend PgTypeToBool
type Converter interface {
	ConvertFromDbStudents(source []db.Student) []OutFromDbStudent
	ConvertFromDBStudent(source db.Student) OutFromDbStudent

	ConvertAddress(source db.Address) OutAddress
	ConvertAddressesItems(source []db.Address) []OutAddress
}

func PgTypeToInt4(i pgtype.Int4) int32 {

	return i.Int32
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
type OutFromDbStudent struct {
	ID             int32  `json:"id"`
	StudentID      int64  `json:"student_id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Age            int32  `json:"age"`
	Email          string `json:"email"`
	Gender         string `json:"gender"`
	FavouriteColor string `json:"favourite_color"`
	StudentAddress int64  `json:"student_address"`
	CreatedAt      int64  `json:"created_at"`
	UpdatedAt      int64  `json:"updated_at"`
	Deleted        bool   `json:"deleted"`
}
