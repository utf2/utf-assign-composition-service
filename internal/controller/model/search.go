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

type SearchStudentAssignedFormsRequest struct {
	StudentID uuid.UUID
}

type SearchStudentAssignedFormsResponse struct {
	AssignedForms []AssignedFormDTO
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

type SearchTeacherCreatedFormsRequest struct {
	TeacherID uuid.UUID
}

type SearchTeacherCreatedFormsResponse struct {
	CreatedForms []CreatedFormDTO
}

type CreatedFormDTO struct {
	ID             uuid.UUID
	Name           string
	Description    string
	CreateDate     time.Time
	StartDate      time.Time
	EndDate        time.Time
	AssignedGroups []GroupDTO
}

type GroupDTO struct {
	ID                 uuid.UUID
	SpecializationCode string
	GroupNumber        string
}
