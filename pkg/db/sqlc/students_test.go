package db

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	var userId string = "1"
	adrs := CreateAddressParams{
		AddressID: userId,
		Street:    pgtype.Text{String: "main", Valid: true},
		City:      pgtype.Text{String: "Kiev", Valid: true},
		Planet:    pgtype.Text{String: "Earth", Valid: true},
		Phone:     "123456789",
	}
	//adr := []interface{}{adrs}
	adr, err := testQueries.CreateAddress(context.Background(), adrs)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(adr)

	arg := CreateStudentParams{
		StudentID:      userId,
		FirstName:      "John",
		LastName:       "Doe",
		Age:            pgtype.Int8{Int64: 20, Valid: true},
		Email:          "email@email18.com",
		Gender:         pgtype.Text{String: "Male", Valid: true},
		FavouriteColor: pgtype.Text{String: "255.255.255", Valid: true},
		StudentAddress: adrs.AddressID,
		CreatedAt:      pgtype.Timestamptz{Time: time.Now().UTC(), Valid: true},
		UpdatedAt:      pgtype.Timestamptz{Time: time.Now().UTC(), Valid: true},
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
	required.Equal(stdnt.StudentAddress, arg.StudentAddress)
	required.Equal(stdnt.Deleted.Bool, arg.Deleted.Bool)
	required.NotZero(stdnt.ID)

}
func TestGetStudent(t *testing.T) {
	arg, _ := testQueries.GetStudent(context.Background(), 1)
	required := require.New(t)
	required.NotEmpty(arg)
}
func TestUpdateStudent(t *testing.T) {
	arg := testQueries.UpdateStudent(context.Background(), UpdateStudentParams{
		ID:        1,
		FirstName: "Johny",
		LastName:  "SDG",
		Age:       pgtype.Int8{Int64: 20, Valid: true}},
	)
	if arg != nil {
		t.Fatal(arg)
	}

}
func TestDeleteStudent(t *testing.T) {
	arg := testQueries.DeleteStudent(context.Background(), 3)
	if arg != nil {
		t.Fatal(arg)
	}
}
