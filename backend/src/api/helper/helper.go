package helper

import (
	"backend/src/entity"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Unmarshal(c *gin.Context, t interface{}) error {
	x, _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal(x, &t)
	if err != nil {
		return entity.ErrInvalidInput
	}
	return nil
}

func ErrHandler(err error, c *gin.Context) {

	switch err {
	case entity.ErrPhoneAlreadyExists:
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": err.Error()})
	case entity.ErrPhoneDoesNotExist:
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	case entity.ErrInvalidInput:
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	default:
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
