package godlp

import "errors"

var (
	ErrStarting   = errors.New("[ERROR] godlp: FAILED WHILE STARTING")
	ErrFetching   = errors.New("[ERROR] godlp: FAILED WHILE FETCHING")
	ErrResponding = errors.New("[ERROR] godlp: FAILED WHILE PROVIDING RESPONSE")
	ErrPreparing  = errors.New("[ERROR] godlp: FAILED WHILE PREPARING A PAYLOAD")
)
