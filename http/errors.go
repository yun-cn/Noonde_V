package http

import (
	"encoding/json"
	"github.com/yun313350095/Noonde/api"
	"net/http"
)

var status = map[string]int{
	api.ErrUnknown:                         http.StatusInternalServerError,
	api.ErrInvalidRequest:                  http.StatusInternalServerError,
	api.ErrForbidden:                       http.StatusForbidden,
	api.ErrHoge:                            http.StatusInternalServerError,
	api.ErrThisIsAVeryVeryLongErrorMessage: http.StatusInternalServerError,
}

var message = map[string]string{
	api.ErrUnknown:                         "Unknown error",
	api.ErrInvalidRequest:                  "Invalid request",
	api.ErrForbidden:                       "Forbidden",
	api.ErrHoge:                            "Hoge",
	api.ErrThisIsAVeryVeryLongErrorMessage: "Thi is a very very long error message",
}

var paramMessage = map[string]string{
	api.ParamEmptyValue:    "Empty value",
	api.ParamInvalidFormat: "Invalid format",
	api.ParamHoge:          "Hoge",
	api.ParamDuplicateUser: "Duplicate user",
	api.ParamInvalidValue:  "Invalid value",
}

type paramError map[string]struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Error ..
func (perr paramError) Error() string {
	body, _ := json.Marshal(perr)
	return string(body)
}

// IsNotEmpty ..
func (perr paramError) IsNotEmpty() bool {
	return len(perr) > 0
}

// PushIfNotExists ..
func (perr paramError) PushIfNotExists(key string, code string) {
	if _, ok := perr[key]; ok {
		return
	}

	perr[key] = struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{
		code,
		paramMessage[code],
	}
}
