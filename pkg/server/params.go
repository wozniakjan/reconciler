package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Params struct {
	params map[string]string
}

func NewParams(r *http.Request) *Params {
	return &Params{
		params: mux.Vars(r),
	}
}
func (p *Params) String(name string) (string, error) {
	result, ok := p.params[name]
	if !ok {
		return "", fmt.Errorf("Parameter '%s' undefined", name)
	}
	return result, nil
}

func (p *Params) Int64(name string) (int64, error) {
	strResult, err := p.String(name)
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(strResult, 10, 64)
}
