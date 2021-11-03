package sample

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrGradeIsUndefined = errors.New("grade is undefined")
var ErrCannotConvertToGrade = errors.New("cannot convert to certain number")
var ErrGradeIsOutOfRange = errors.New("grade is out of range")

// Grade is type to notify rubi-adding level
type Grade int

const (
	// UnknownGrade means grade is undefined
	UnknownGrade Grade = -1
	// DefaultGrade means use default setting
	DefaultGrade Grade = 0
	// FirstGrade means adding rubi to all text
	FirstGrade  Grade = 1
	SecondGrade Grade = 2
	ThirdGrade  Grade = 3
	ForthGrade  Grade = 4
	FifthGrade  Grade = 5
	SixthGrade  Grade = 6
	JuniorGrade Grade = 7
	// NormalGrade means not adding rubi to common Kanji
	NormalGrade Grade = 8
)

func gradeInRange(g Grade) bool {
	if g < DefaultGrade || NormalGrade < g {
		return false
	}
	return true
}

func (g Grade) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "%d", g)
	return buf.Bytes(), nil
}

func (g *Grade) UnmarshalJSON(b []byte) error {
	var sb strings.Builder
	sb.Write(b)

	num, err := strconv.Atoi(sb.String())
	*g = Grade(num)

	if err != nil {
		*g = UnknownGrade
		err = ErrCannotConvertToGrade
	} else if !gradeInRange(*g) {
		*g = UnknownGrade
		err = ErrGradeIsOutOfRange
	}
	return err
}
