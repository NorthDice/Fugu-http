package request

import (
	"fmt"
	"io"
	"strings"
)

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

type Request struct {
	RequestLine RequestLine
}

const (
	SEPARATOR = "\n\r"
)

var BAD_START_LINE = fmt.Errorf("bad start line")

func parseRequestLine(b string) (*RequestLine, string, error) {
	idx := strings.Index(b, SEPARATOR)
	if idx == -1 {
		return nil, b, nil
	}
}

func RequestFromReader(reader io.Reader) (*Request, error) {

}
