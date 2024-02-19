// Package controllers controllers/members_controllers.go
package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sample/src/models"
	"strconv"
)

func indexMember(c *gin.Context) {
	repo, err := models.NewRepository()
	if err != nil {
		log.Fatalf("failed to create repository: %v", err)
	}
	notificationMessage := c.Query("notificationMessage")
	people := repo.AllPerson()
	c.HTML(http.StatusOK, "members/index.tmpl", gin.H{
		"people":              people,
		"notificationMessage": notificationMessage,
	})
}

func showMember(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	repo, err := models.NewRepository()
	if err != nil {
		log.Fatalf("failed to create repository: %v", err)
	}
	person := repo.FindPerson(id)
	c.HTML(http.StatusOK, "members/show.tmpl", gin.H{
		"title": "show",
		"name":  person.Name,
		"id":    person.Id,
	})
}

func createMember(c *gin.Context) {
	name := c.PostForm("name")
	repo, err := models.NewRepository()
	if err != nil {
		log.Fatalf("failed to create repository: %v", err)
	}
	repoErr := repo.CreatePerson(name)
	if repoErr != nil {
		var errorMessages []string
		errorMessages = append(errorMessages, repoErr.Error())
		c.HTML(http.StatusBadRequest, "members/new.tmpl", gin.H{})
		return
	}
	c.Redirect(http.StatusFound, "/members")
	return
}

func newMember(c *gin.Context) {
	c.HTML(http.StatusOK, "members/new.tmpl", gin.H{})
}

func deleteMember(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	repo, err := models.NewRepository()
	if err != nil {
		log.Fatalf("failed to create repository: %v", err)
	}
	repo.DeletePerson(id)
	c.JSON(200, gin.H{
		"id": id,
	})
}

func selectMember(c *gin.Context) {
	repo, err := models.NewRepository()
	if err != nil {
		log.Fatalf("failed to create repository: %v", err)
	}
	person := repo.SamplePerson()
	c.JSON(200, gin.H{
		"id":   person.Id,
		"name": person.Name,
	})
}
