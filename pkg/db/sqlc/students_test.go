package db

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	arg := CreateStudentParams{
		FirstName:      "John",
		LastName:       "Doe",
		Age:            pgtype.Int8{Int64: 20, Valid: true},
		Email:          "email@email.com",
		FavouriteColor: pgtype.Text{String: "255.255.255", Valid: true},
		Addresses:      pgtype.Text{String: "1234 Main St, Springfield, Earth", Valid: true},
		CreatedAt:      pgtype.Timestamptz{Time: time.Now(), Valid: true},
		UpdatedAt:      pgtype.Timestamptz{Time: time.Now(), Valid: true},
		Deleted:        pgtype.Bool{Bool: false, Valid: true},
	}
	stdnt, err := testQueries.CreateStudent(context.Background(), arg)
	if err != nil {
		t.Fatal(err)
	}
	required := require.New(t)
	required.NotEmpty(stdnt)
	required.Equal(stdnt.FirstName, arg.FirstName)
	required.Equal(stdnt.LastName, arg.LastName)
	required.Equal(stdnt.Age.Int64, arg.Age.Int64)
	required.Equal(stdnt.Email, arg.Email)
	required.Equal(stdnt.FavouriteColor.String, arg.FavouriteColor.String)
	required.Equal(stdnt.Addresses.String, arg.Addresses.String)
	required.Equal(stdnt.CreatedAt.Time, arg.CreatedAt.Time)
	required.Equal(stdnt.UpdatedAt.Time, arg.UpdatedAt.Time)
	required.Equal(stdnt.Deleted.Bool, arg.Deleted.Bool)
	required.NotZero(stdnt.ID)

}
