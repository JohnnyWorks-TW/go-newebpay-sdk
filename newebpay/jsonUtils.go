package newebpay

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type JsonPayTime time.Time

func (j *JsonPayTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02 15:04:05", s)

	if err != nil {
		return err
	}
	*j = JsonPayTime(t)
	return nil
}

func (j JsonPayTime) MarshalJSON() ([]byte, error) {
	return []byte("\"" + j.Format("2006-01-02 15:04:05") + "\""), nil
}

func (j JsonPayTime) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}

type StringInt int64

func (j *StringInt) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	value, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}

	*j = StringInt(value)
	return nil
}

func (j StringInt) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprint(j)), nil
}
