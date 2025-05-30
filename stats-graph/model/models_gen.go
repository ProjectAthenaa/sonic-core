// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type AxisNames struct {
	X string `json:"X"`
	Y string `json:"Y"`
}

type StatCategory struct {
	ID          string     `json:"ID"`
	Name        string     `json:"Name"`
	Labels      *AxisNames `json:"Labels"`
	Type        StatType   `json:"Type"`
	AntibotName *string    `json:"AntibotName"`
}

type Statistic struct {
	XValue int `json:"X_Value"`
	YValue int `json:"Y_Value"`
}

type StatType string

const (
	StatTypeCheckouts      StatType = "CHECKOUTS"
	StatTypeDeclines       StatType = "DECLINES"
	StatTypeErrors         StatType = "ERRORS"
	StatTypeFailed         StatType = "FAILED"
	StatTypeCookieGens     StatType = "COOKIE_GENS"
	StatTypeRecaptchaUsage StatType = "RECAPTCHA_USAGE"
	StatTypeTasksRunning   StatType = "TASKS_RUNNING"
	StatTypeMoneySpent     StatType = "MONEY_SPENT"
)

var AllStatType = []StatType{
	StatTypeCheckouts,
	StatTypeDeclines,
	StatTypeErrors,
	StatTypeFailed,
	StatTypeCookieGens,
	StatTypeRecaptchaUsage,
	StatTypeTasksRunning,
	StatTypeMoneySpent,
}

func (e StatType) IsValid() bool {
	switch e {
	case StatTypeCheckouts, StatTypeDeclines, StatTypeErrors, StatTypeFailed, StatTypeCookieGens, StatTypeRecaptchaUsage, StatTypeTasksRunning, StatTypeMoneySpent:
		return true
	}
	return false
}

func (e StatType) String() string {
	return string(e)
}

func (e *StatType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = StatType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid StatType", str)
	}
	return nil
}

func (e StatType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
