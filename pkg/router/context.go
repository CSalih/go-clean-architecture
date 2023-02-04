package router

import (
	"encoding/json"
	"net/http"
)

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
	Params  map[string]string
}

func (c *Context) Json(code int, data interface{}) error {
	jsonString, err := json.Marshal(data)
	if err != nil {
		return err
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(code)
	_, err = c.Writer.Write(jsonString)
	if err != nil {
		return err
	}
	return nil
}
