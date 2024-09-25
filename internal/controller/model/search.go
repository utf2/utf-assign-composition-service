package transfer

import (
	"time"

	"github.com/google/uuid"
)

type StudentFormStatus string

const (
	Passed       StudentFormStatus = "Пройдена"
	NotPassedYet StudentFormStatus = "Еще не пройдена"
	Deadlined    StudentFormStatus = "Просрочена"
)

type TeacherDTO struct {
	ID         uuid.UUID
	FirstName  string
	LastName   string
	MiddleName string
}

type StudentFormResultDTO struct {
	ScoredPoints      uint
	MaxPossiblePoints uint
}

type AssignedFormDTO struct {
	ID          uuid.UUID
	Name        string
	Description string
	StartDate   time.Time
	EndDate     time.Time
	Teacher     TeacherDTO
	Status      StudentFormStatus
	Result      StudentFormResultDTO
}

type SearchStudentAssignedFormsRequest struct {
	StudentID uuid.UUID
}

type SearchStudentAssignedFormsResponse struct {
	AssignedForms []AssignedFormDTO
}

// TODO
type CreatedFormDTO struct {
}

type SearchTeacherCreatedFormsRequest struct {
	TeacherID uuid.UUID
}

type SearchTeacherCreatedFormsResponse struct {
	CreatedForms []CreatedFormDTO
}
