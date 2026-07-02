package request

import (
	"fugu-http/internal/domain"
	"io"
	"strings"
)

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func (r *RequestLine) ValidHTTP() bool {
	return r.HttpVersion == "HTTP/1.1"
}

type Request struct {
	RequestLine RequestLine
}

const (
	SEPARATOR = "\n\r"
)

func parseRequestLine(b string) (*RequestLine, string, error) {
	idx := strings.Index(b, SEPARATOR)
	if idx == -1 {
		return nil, b, nil
	}

	startLine := b[:idx]
	restOfMessage := b[idx+len(SEPARATOR):]

	parts := strings.Split(startLine, " ")
	if len(parts) != 3 {
		return nil, restOfMessage, domain.MALFORMED_REQUEST_LINE
	}
	rl := RequestLine{
		Method:        parts[0],
		RequestTarget: parts[1],
		HttpVersion:   parts[2],
	}, restOfMessage, nil

}

func RequestFromReader(reader io.Reader) (*Request, error) {

}
