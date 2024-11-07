package response

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/aa-ar/httpx/errors"
)

type Response struct {
	status  int
	Body    interface{}
	Cookies []*http.Cookie
}

func NewResponse(status int, body interface{}, cookies []*http.Cookie) *Response {
	if err, ok := body.(errors.Error); ok {
		if err != nil {
			body = Error{
				Error:   err.Error(),
				Details: err,
			}
		}
	}

	return &Response{
		status:  status,
		Body:    body,
		Cookies: cookies,
	}
}

func (res *Response) marshalBody() []byte {
	if res.Body == nil {
		return *new([]byte)
	}
	if res, err := json.Marshal(res.Body); err == nil {
		return res
	} else {
		slog.Error(err.Error())
		return *new([]byte)
	}
}

func (res *Response) WriteTo(w http.ResponseWriter) {
	r := res.marshalBody()
	if len(r) == 0 {
		res.status = http.StatusNoContent
	}
	if w, ok := w.(http.ResponseWriter); ok {
		for _, cookie := range res.Cookies {
			http.SetCookie(w, cookie)
		}
		w.WriteHeader(res.status)
	}
	fmt.Fprint(w, string(r))
}
