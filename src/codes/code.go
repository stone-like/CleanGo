package codes

type Code string

const (
	OK             Code = "OK"
	InvalidRequest Code = "invalid_request"
	NotFound       Code = "not_found"
	Database       Code = "database_error"
	Internal       Code = "internal_error"

	Unknown Code = "unknown"
)
