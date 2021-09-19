package handler

import (
	"backend/src/api/helper"
	"backend/src/api/model"
	"backend/src/usecase/phoneNumber"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinHandler struct {
	PhoneNumberUseCase phoneNumber.UseCase
}

func NewGinHandler(useCase phoneNumber.UseCase) (r *gin.Engine) {
	h := &GinHandler{
		useCase,
	}
	r = gin.Default()
	r.Use(CORSMiddleware())
	r.GET("/phone-numbers", h.GetPhoneNumbers)
	r.POST("/phone-numbers", h.CreateNumber)
	return r
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (h *GinHandler) GetPhoneNumbers(c *gin.Context) {
	limit := c.DefaultQuery("limit", "10")
	offset := c.DefaultQuery("offset", "0")
	country := c.Query("country")
	state := c.Query("state")

	phonenumbers, err := h.PhoneNumberUseCase.FindAllPhoneNumbers(offset, limit, country, state)
	if err != nil {
		helper.ErrHandler(err, c)
		return
	}
	c.JSON(http.StatusOK, phonenumbers)
}

func (h *GinHandler) CreateNumber(c *gin.Context) {

	var numberPayload model.CreatePhoneNumberInput
	err := helper.Unmarshal(c, &numberPayload)
	if err != nil {
		helper.ErrHandler(err, c)
		return
	}

	if err := h.PhoneNumberUseCase.CreatePhoneNumber(numberPayload.PhoneNumber); err != nil {
		helper.ErrHandler(err, c)
		return
	}

	response := model.CreatePhoneNumberOutput{Message: "PhoneNumber created successfully"}
	c.JSON(http.StatusCreated, response)
}
