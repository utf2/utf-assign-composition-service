package controller

import transfer "github.com/utf2/utf-assign-composition-service/internal/controller/model"

type AssignCompositionControllerAPI interface {
	SearchStudentForms(transfer.SearchStudentAssignedFormsRequest) transfer.SearchStudentAssignedFormsResponse
	SearchTeacherForms(transfer.SearchTeacherCreatedFormsRequest) transfer.SearchTeacherCreatedFormsResponse
}
