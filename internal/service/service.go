package service

import (
	"context"
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
	CreateStudent(ctx context.Context, input models.InputStudent) (out models.ResponseApiStudent, err error)
	GetStudents(ctx context.Context) (resp []models.ResponseApiStudent, err error)
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
func (r *Repository) CreateStudent(ctx context.Context, input models.InputStudent) (out models.ResponseApiStudent, err error) {

	studentId, err := r.db.GetIndexes(context.Background())
	if err != nil {

		studentId, err = r.db.CreateIndex(ctx, 1)
		if err != nil {
			r.Logger.Error("CreateIndex", zap.Error(err))
			return out, errors.WithStack(err)
		}
	}

	adrs := db.CreateStudentAddressParams{
		AddressID: studentId,
		Street:    pgtype.Text{String: input.Addresses.Street, Valid: true},
		City:      pgtype.Text{String: input.Addresses.City, Valid: true},
		Planet:    pgtype.Text{String: input.Addresses.Planet, Valid: true},
		Phone:     input.Addresses.Phone,
	}

	adr, err := r.db.CreateStudentAddress(context.Background(), adrs)
	if err != nil {
		r.Logger.Error("CreateAddress", zap.Error(err))
		return out, errors.WithStack(err)
	}

	arg := db.CreateStudentParams{
		StudentID:      studentId,
		FirstName:      input.FirstName,
		LastName:       input.LastName,
		Age:            pgtype.Int4{Int32: input.Age, Valid: true},
		Email:          input.Email,
		Gender:         pgtype.Text{String: input.Gender, Valid: true},
		FavouriteColor: pgtype.Text{String: input.FavouriteColor, Valid: true},
		StudentAddress: adr.AddressID,
		CreatedAt:      time.Now().Unix(),
		UpdatedAt:      time.Now().Unix(),
		Deleted:        pgtype.Bool{Bool: false, Valid: true},
	}

	stdnt, err := r.db.CreateStudent(context.Background(), arg)
	if err != nil {
		r.Logger.Error("CreateStudent", zap.Error(err))
		return out, errors.WithStack(err)
	}

	impl := generated.ConverterImpl{}
	var outStudent = impl.ConvertFromDBStudent(stdnt)
	var address = impl.ConvertAddress(adr)
	ApiImpl := models.ImplementationConvertFromDbToApi{}
	var response = ApiImpl.ConvertResponseApi(outStudent, address)

	err = r.db.UpdateIndex(context.Background(), db.UpdateIndexParams{
		IndexID:   studentId,
		IndexID_2: studentId + 1,
	})
	if err != nil {
		r.Logger.Error("UpdateIndex", zap.Error(err))
		return out, errors.WithStack(err)
	}
	return response, nil
}

func (r *Repository) GetStudents(ctx context.Context) (resp []models.ResponseApiStudent, err error) {

	stdnts, err := r.db.ListStudents(ctx, db.ListStudentsParams{
		Limit:  10,
		Offset: 0,
	})
	if err != nil {
		r.Logger.Error("ListStudents", zap.Error(err))
		return nil, errors.WithStack(err)
	}
	addresses, err := r.db.GetStudentAddresses(context.Background(), db.GetStudentAddressesParams{Limit: 10, Offset: 0})
	if err != nil {
		r.Logger.Error("GetStudentAddresses", zap.Error(err))
		return nil, errors.WithStack(err)
	}

	impl := generated.ConverterImpl{}
	var outStudents = impl.ConvertFromDbStudents(stdnts)
	var convertedAddresses = impl.ConvertAddressesItems(addresses)
	ApiImpl := models.ImplementationConvertFromDbToApi{}
	resp = ApiImpl.ConvertResponseApiWithItems(outStudents, convertedAddresses)

	return resp, nil
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
		Age:       pgtype.Int4{Int32: input.Age, Valid: true},
		UpdatedAt: time.Now().Unix(),
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
		r.Logger.Error("DeleteParseIntStudentID", zap.Error(err))
		return errors.WithStack(err)
	}
	err = r.db.DeleteStudent(ctx, int32(conv))
	if err != nil {
		r.Logger.Error("DeleteStudent", zap.Error(err))
		return errors.WithStack(err)
	}
	return nil
}
