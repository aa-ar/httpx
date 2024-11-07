package request

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"

	"github.com/aa-ar/httpx/errors"
)

type Decoder struct {
	*json.Decoder
}

func Decode(r io.Reader) *Decoder {
	return &Decoder{json.NewDecoder(r)}
}

func (d *Decoder) Into(v interface{}) error {
	if err := d.Decode(&v); err != nil {
		slog.Error(err.Error())
		jerr, ok := err.(*json.UnmarshalTypeError)
		if !ok {
			return &errors.BadRequestError{Message: "Bad request body"}
		}
		msg := fmt.Sprintf("Expected %s to be %s, %s found instead", jerr.Field, jerr.Type.String(), jerr.Value)
		return &BadRequestBodyError{Message: msg}
	}
	return nil
}
