package domain

import "fmt"

var (
	MALFORMED_REQUEST_LINE  = fmt.Errorf("malformed request-line")
	UNSUPORTED_HTTP_VERSION = fmt.Errorf("unsuported http version")
)
