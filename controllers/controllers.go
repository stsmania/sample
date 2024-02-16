// controllers/controllers.go
package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sample/models/person"
	"strconv"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/", indexMember)

	member := r.Group("/members")
	{
		member.GET("/new", newMember)
		member.GET("/:id", showMember)
		member.DELETE("/:id", deleteMember)
		member.GET("", indexMember)
		member.POST("", createMember)
		member.GET("/select", selectMember)
	}
}

func indexMember(c *gin.Context) {
	people := person.Get_all()
	c.HTML(http.StatusOK, "members/index.tmpl", gin.H{
		"title":  "Hage",
		"people": people,
	})
}

func showMember(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	person := person.Db_find(id)
	c.HTML(http.StatusOK, "members/show.tmpl", gin.H{
		"title": "show",
		"name":  person.Name,
		"id":    person.Id,
	})
}

func createMember(c *gin.Context) {
	name := c.PostForm("name")
	person.Create(name)
	c.Redirect(http.StatusFound, "/")
}

func newMember(c *gin.Context) {
	c.HTML(http.StatusOK, "members/new.tmpl", gin.H{})
}

func deleteMember(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	person.Db_delete(id)
	c.JSON(200, gin.H{
		"id": id,
	})
}

func selectMember(c *gin.Context) {
    people := person.Sample()
    c.JSON(200, gin.H{
        "id": people.Id,
        "name": people.Name,
    })
}

