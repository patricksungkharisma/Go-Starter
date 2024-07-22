package pkg

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateContextMock(method, uri string, bodyMap, uriParam *map[string]string, bodyJSON interface{}, postFromMap *map[string]string) *gin.Context {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	var body io.Reader
	if bodyJSON == nil {
		form := url.Values{}
		if bodyMap != nil {
			for k, v := range *bodyMap {
				form.Add(k, v)
			}
		}
		body = strings.NewReader(form.Encode())
	} else {
		bodyByte, _ := json.Marshal(bodyJSON)
		body = bytes.NewBuffer(bodyByte)
	}

	c.Request, _ = http.NewRequest(method, uri, body)

	if bodyJSON == nil {
		c.Request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	} else {
		c.Request.Header.Add("Content-Type", "application/json")
	}

	if uriParam != nil {
		for k, v := range *uriParam {
			c.Params = append(c.Params, gin.Param{
				Key:   k,
				Value: v,
			})
		}
	}

	if postFromMap != nil {
		c.Request.PostForm = url.Values{}
		for i, j := range *postFromMap {
			c.Request.PostForm.Set(i, j)
		}
	}

	return c
}
