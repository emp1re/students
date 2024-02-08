package service

import (
	"context"
	"math/rand"
	"strconv"
	"time"

	db "github.com/emp1re/students/pkg/db/sqlc"
	"github.com/emp1re/students/pkg/models"
	"github.com/emp1re/students/pkg/models/generated"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Service interface {
	CreateStudent(ctx context.Context, input models.InputStudent) (out models.OutStudent, err error)
	GetStudents(ctx context.Context) (out []models.OutStudent, err error)
	UpdateStudent(ctx context.Context, input models.InputUpdateStudent, id string) error
	DeleteStudent(ctx context.Context, id string) error
}
type Repository struct {
	context.Context
	*zap.Logger
	*pgx.Conn
	db *db.Queries
}

func NewService(ctx context.Context, log *zap.Logger, postg *pgx.Conn, q *db.Queries) *Repository {
	return &Repository{
		Context: ctx,
		Logger:  log,
		Conn:    postg,
		db:      db.New(postg)}
}
func (r *Repository) CreateStudent(ctx context.Context, input models.InputStudent) (out models.OutStudent, err error) {
	//fmt.Println(input)
	studentId := input.LastName + string(rune(rand.Intn(100)))
	//r.db.CreateStudent(context.Background())

	adrs := db.CreateAddressParams{
		AddressID: studentId,
		Street:    pgtype.Text{String: input.Addresses.Street, Valid: true},
		City:      pgtype.Text{String: input.Addresses.City, Valid: true},
		Planet:    pgtype.Text{String: input.Addresses.Planet, Valid: true},
		Phone:     input.Addresses.Phone,
	}

	adr, err := r.db.CreateAddress(ctx, adrs)
	if err != nil {
		r.Logger.Error("CreateAddress", zap.Error(err))
		return out, errors.WithStack(err)
	}

	arg := db.CreateStudentParams{
		StudentID:      studentId,
		FirstName:      input.FirstName,
		LastName:       input.LastName,
		Age:            pgtype.Int8{Int64: input.Age, Valid: true},
		Email:          input.Email,
		Gender:         pgtype.Text{String: input.Gender, Valid: true},
		FavouriteColor: pgtype.Text{String: input.FavouriteColor, Valid: true},
		StudentAddress: adr.AddressID,
		CreatedAt:      pgtype.Timestamptz{Time: time.Now().UTC(), Valid: true},
		UpdatedAt:      pgtype.Timestamptz{Time: time.Now().UTC(), Valid: true},
		Deleted:        pgtype.Bool{Bool: false, Valid: true},
	}

	stdnt, err := r.db.CreateStudent(r.Context, arg)
	if err != nil {
		r.Logger.Error("CreateStudent", zap.Error(err))
		return out, errors.WithStack(err)
	}

	impl := generated.ConverterImpl{}
	var outStudent = impl.Convert(stdnt)

	return outStudent, nil
}

func (r *Repository) GetStudents(ctx context.Context) (out []models.OutStudent, err error) {
	stdnts, err := r.db.ListStudents(ctx, db.ListStudentsParams{
		Limit:  10,
		Offset: 0,
	})
	if err != nil {
		r.Logger.Error("ListStudents", zap.Error(err))
		return nil, errors.WithStack(err)
	}
	impl := generated.ConverterImpl{}
	var outStudents = impl.ConvertItems(stdnts)
	return outStudents, nil
}
func (r *Repository) UpdateStudent(ctx context.Context, input models.InputUpdateStudent, id string) error {
	//fmt.Println(input)
	conv, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		r.Logger.Error("UpdateStudent", zap.Error(err))
		return errors.WithStack(err)
	}
	arg := db.UpdateStudentParams{
		ID:        int32(conv),
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Age:       pgtype.Int8{Int64: input.Age, Valid: true},
	}
	err = r.db.UpdateStudent(ctx, arg)
	if err != nil {
		r.Logger.Error("UpdateStudent", zap.Error(err))
		return errors.WithStack(err)
	}
	return nil
}
func (r *Repository) DeleteStudent(ctx context.Context, id string) error {
	conv, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		r.Logger.Error("UpdateStudent", zap.Error(err))
		return errors.WithStack(err)
	}
	err = r.db.DeleteStudent(ctx, int32(conv))
	if err != nil {
		r.Logger.Error("DeleteStudent", zap.Error(err))
		return errors.WithStack(err)
	}
	return nil
}
