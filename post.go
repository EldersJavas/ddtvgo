// Created by EldersJavas(EldersJavas&gmail.com)

package ddtvgo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func test() error {
	router := gin.Default()
	router.POST("/events", events)
	err := router.Run(":5000")
	if err != nil {
		return err
	}
	return nil
}

func events(c *gin.Context) {
	buf := make([]byte, 1024)

	n, _ := c.Request.Body.Read(buf)

	fmt.Println(string(buf[0:n]))

	resp := map[string]string{"hello": "world"}
	c.JSON(http.StatusOK, resp)

	/*post_gwid := c.PostForm("name")
	  fmt.Println(post_gwid)*/

}
