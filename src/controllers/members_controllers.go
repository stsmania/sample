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
	members := repo.AllMember()
	c.HTML(http.StatusOK, "members/index.tmpl", gin.H{"members": members})
}

func showMember(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	repo, err := models.NewRepository()
	if err != nil {
		log.Fatalf("failed to create repository: %v", err)
	}
	member := repo.FindMember(id)
	c.HTML(http.StatusOK, "members/show.tmpl", gin.H{
		"title": "show",
		"name":  member.Name,
		"id":    member.Id,
	})
}

func createMember(c *gin.Context) {
	name := c.PostForm("name")
	repo, err := models.NewRepository()
	if err != nil {
		log.Fatalf("failed to create repository: %v", err)
	}
	repoErr := repo.CreateMember(name)
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
	repo.DeleteMember(id)
	c.JSON(200, gin.H{
		"id": id,
	})
}

func randomMember(c *gin.Context) {
	repo, err := models.NewRepository()
	if err != nil {
		log.Fatalf("failed to create repository: %v", err)
	}
	member := repo.RandomMember()
	c.JSON(200, gin.H{
		"id":   member.Id,
		"name": member.Name,
	})
}
