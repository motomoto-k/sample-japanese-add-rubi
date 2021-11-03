package sample

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

type Response struct {
	Result []Word
	Error  error
}

type Word struct {
	Furigana string `json:"furigana"`
	Roman    string `json:"roman"`
	SubWord  []Word `json:"subword"`
	Surface  string `json:"surface"`
}

type response struct {
	ID      string `json:"id"`
	JsonRPC string `json:"jsonrpc"`
	Result  struct {
		Word []Word `json:"word"`
	} `json:"result"`
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

type invalidError struct {
	Error struct {
		Message string `json:"Message"`
	} `json:"Error"`
}

func readResponseMessage(src io.Reader, id string, res *Response) error {
	var r response
	err := json.NewDecoder(src).Decode(&r)
	if err != nil {
		return err
	}

	if id != r.ID {
		return errors.New("no match id string")
	}

	if r.Error.Code == 0 {
		res.Result = r.Result.Word
	} else {
		res.Error = fmt.Errorf("responsed error : Code=%d Message=%s", r.Error.Code, r.Error.Message)
	}
	return nil
}

func readResponseMessageAsInvalid(src io.Reader, res *Response) error {
	var e invalidError
	err := json.NewDecoder(src).Decode(&e)
	if err != nil {
		return err
	}

	res.Error = fmt.Errorf("responsed error : %s", e.Error.Message)
	return nil
}
