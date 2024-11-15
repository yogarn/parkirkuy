package model

type UBAuthReq struct {
	Identifier string `json:"identifier" form:"identifier" validate:"required"`
	Password   string `json:"password" form:"password" validate:"required"`
}

type UBAuthRes struct {
	NIM          string `json:"nim"`
	FullName     string `json:"fullName"`
	Email        string `json:"email"`
	Faculty      string `json:"faculty"`
	StudyProgram string `json:"studyProgram"`
}
