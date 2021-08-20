package authentication

import "errors"

var (
	ipDoesNotMatchSessionError = errors.New("ip_does_not_match_session")
	unauthorizedError = errors.New("unauthorized")
)
