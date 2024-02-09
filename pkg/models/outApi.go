package models

import "time"

type ConvertFromDbToApi interface {
	ConvertResponseApi(source OutFromDbStudent, sourceAdress OutAddress) ResponseApiStudent
	ConvertResponseApiWithItems(source []OutFromDbStudent, sourceAdress []OutAddress) ResponseApiStudentWithItems
}
type ResponseApiStudent struct {
	ID             int32      `json:"id"`
	StudentID      int64      `json:"student_id"`
	FirstName      string     `json:"first_name"`
	LastName       string     `json:"last_name"`
	Age            int32      `json:"age"`
	Email          string     `json:"email"`
	Gender         string     `json:"gender"`
	FavouriteColor string     `json:"favourite_color"`
	StudentAddress OutAddress `json:"student_address"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	Deleted        bool       `json:"deleted"`
}

type ResponseApiStudentWithItems []ResponseApiStudent

type ImplementationConvertFromDbToApi struct{}

func (s *ImplementationConvertFromDbToApi) ConvertResponseApi(source OutFromDbStudent, sourceAdress OutAddress) ResponseApiStudent {
	var responseApiStudent ResponseApiStudent
	responseApiStudent.ID = source.ID
	responseApiStudent.StudentID = source.StudentID
	responseApiStudent.FirstName = source.FirstName
	responseApiStudent.LastName = source.LastName
	responseApiStudent.Age = source.Age
	responseApiStudent.Email = source.Email
	responseApiStudent.Gender = source.Gender
	responseApiStudent.FavouriteColor = source.FavouriteColor
	responseApiStudent.StudentAddress = sourceAdress
	responseApiStudent.CreatedAt = FromUnixToTime(source.CreatedAt)
	responseApiStudent.UpdatedAt = FromUnixToTime(source.UpdatedAt)
	responseApiStudent.Deleted = source.Deleted
	return responseApiStudent
}
func (s *ImplementationConvertFromDbToApi) ConvertResponseApiWithItems(source []OutFromDbStudent, sourceAddress []OutAddress) ResponseApiStudentWithItems {
	var responseApiStudentWithItems ResponseApiStudentWithItems
	if source != nil {
		responseApiStudentWithItems = make([]ResponseApiStudent, len(source))
		for i := 0; i < len(source); i++ {
			responseApiStudentWithItems[i] = s.ConvertResponseApi(source[i], sourceAddress[i])
		}
	}
	return responseApiStudentWithItems
}
func FromUnixToTime(i int64) time.Time {
	return time.Unix(i, 0)
}
