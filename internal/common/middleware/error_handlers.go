package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrResponse struct {
	Status     bool        `json:"status"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Result     interface{} `json:"result"`
}

func (e ErrResponse) ToJson() (raw []byte, err error) {
	raw, err = json.Marshal(e)
	return raw, err
}

func GinErrorHandle(h func(c *gin.Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := h(c); err != nil {
			if len(c.Errors) == 0 {
				c.Error(err)
			}

			GinRecoveryFn(c)
		}
	}
}

func GinRecoveryFn(c *gin.Context) {
	err := c.Errors.Last()
	if err == nil {
		return
	}

	resp := &ErrResponse{
		Status:     false,
		StatusCode: 999999,
		Message:    err.Err.Error(),
		Result:     struct{}{},
	}

	// Если ошибка уже была обработана самим фреймворком, дописываем только тело ответа
	if c.IsAborted() {
		resp.StatusCode = 1001
		rawResp, _ := resp.ToJson()
		c.Writer.Write(rawResp)
		return
	} else {
		c.Abort()
	}

	// ErrorTypeBind - ошибка парсинга тела. Отдаем 400 код
	if err.Type == gin.ErrorTypeBind {
		resp.StatusCode = 10002
		c.JSON(http.StatusBadRequest, resp)
		return
	}
}
