package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		// call HTML method of context to render a new template
		c.HTML(
			http.StatusOK, // set http status to OK
			"index.html", // specify template
			gin.H{ // pass the data that the page uses, here the title (see ln 8 of header.html)
				"title": "Home Page",
			},
		)
	})

	router.Run()
}
