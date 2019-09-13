package api

// Error message and code
const (
	OK                                 = "S00200"
	ErrUnknown                         = "E00001"
	ErrInvalidRequest                  = "E00002"
	ErrForbidden                       = "E00003"
	ErrHoge                            = "E00004"
	ErrThisIsAVeryVeryLongErrorMessage = "E00005"
)

// Request parameter errors
const (
	ParamEmptyValue    = "P00001"
	ParamInvalidFormat = "P00002"
	ParamHoge          = "P00003"
	ParamDuplicateUser = "P00004"
	ParamInvalidValue  = "P00005"
)
