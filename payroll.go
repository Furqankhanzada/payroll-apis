package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

		c.JSON(200, gin.H{"firstname": firstname, "lastname": lastname})
	})
	router.Run(":8080")
}
