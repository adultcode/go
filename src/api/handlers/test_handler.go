package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type personData struct {
	FirstName string `json:"first_name" binding:"required,min=4"`
	LastName  string
}

type header struct {
	UserId  string
	Browser string
}
type TestHandler struct{}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (h *TestHandler) Test(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"result": "Test",
	})
}

func (h *TestHandler) UserById(c *gin.Context) {

	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"user": id,
	})

}

func (h *TestHandler) HeaderBinder1(c *gin.Context) {

	userid := c.GetHeader("userid")
	c.JSON(http.StatusOK, gin.H{
		"userID": userid,
	})

}

func (h *TestHandler) HeaderBinder2(c *gin.Context) {

	head := header{}
	c.BindHeader(&head)
	c.JSON(http.StatusOK, gin.H{
		"userID": head.UserId,
	})
}

func (h *TestHandler) QueryBinder1(c *gin.Context) {

	//head := header{}
	id := c.Query("id")
	//c.BindHeader(&head)
	c.JSON(http.StatusOK, gin.H{
		"userID": id,
	})
}

func (h *TestHandler) BodyBinder(c *gin.Context) {

	//head := header{}
	//id := c.Query("id")
	p := personData{}
	err := c.BindJSON(&p)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})

	} else {
		c.JSON(http.StatusOK, p)

	}
}
