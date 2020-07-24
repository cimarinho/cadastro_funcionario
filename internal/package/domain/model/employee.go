package model

import (
	"encoding/json"
	"time"
)

const (
	layoutISO = "02/01/2006"
)

type JsonBirthDay time.Time

type Employee struct {
	Name               string       `json:"name" binding:"required" `
	Birthday           JsonBirthDay `json:"birthday" time_format:"02/01/2006" time_utc:"1"`
	Salary             float64      `json:"salary"`
	ExperienceTime     int8         `json:"experienceTime"`
	PhysicalDisability string       `json:"physicalDisability" message:"VALIDAR" binding:"required,oneof=sim não"`
}

func (mt *JsonBirthDay) UnmarshalJSON(bs []byte) error {
	var s string
	err := json.Unmarshal(bs, &s)
	if err != nil {
		return err
	}
	t, err := time.Parse(layoutISO, s)
	if err != nil {
		return err
	}
	*mt = JsonBirthDay(t)
	return nil
}

func (j JsonBirthDay) Format() string {
	t := time.Time(j)
	return t.String()
}

func (t JsonBirthDay) IsZero() bool {
	var myDate time.Time
	myDate = time.Time(t)
	return myDate.IsZero()
}
