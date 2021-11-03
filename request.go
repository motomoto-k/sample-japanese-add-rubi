package sample

import (
	"encoding/json"
	"io"
)

type Request struct {
	ID    string
	Text  string
	Grade Grade
}

type request struct {
	ID      string `json:"id"`
	JsonRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Q     string `json:"q"`
		Grade *Grade `json:"grade,omitempty"`
	} `json:"params"`
}

func initializeRequest(r *request) {
	r.JsonRPC = "2.0"
	r.Method = "jlp.furiganaservice.furigana"
}

func writeRequestMessage(dst io.Writer, id string, req *Request) error {
	var r request
	initializeRequest(&r)

	r.ID = id
	r.Params.Q = req.Text
	if req.Grade == UnknownGrade {
		return ErrGradeIsUndefined
	} else if !gradeInRange(req.Grade) {
		return ErrGradeIsOutOfRange
	} else if req.Grade == DefaultGrade {
		r.Params.Grade = nil
	} else {
		r.Params.Grade = &req.Grade
	}

	return json.NewEncoder(dst).Encode(&r)
}
