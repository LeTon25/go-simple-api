package response

import "github.com/gin-gonic/gin"

type ResponseData struct {
	Code    int         `json:"code"`    // code
	Message string      `json:"message"` // thong bao
	Data    interface{} `json:"data"`    // Du lieu
}

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(code, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, code int) {
	c.JSON(code, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}
