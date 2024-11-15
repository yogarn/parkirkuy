package model

type UBAuthReq struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

type UBAuthRes struct {
	NIM          string `json:"nim"`
	FullName     string `json:"fullName"`
	Email        string `json:"email"`
	Faculty      string `json:"faculty"`
	StudyProgram string `json:"studyProgram"`
}
