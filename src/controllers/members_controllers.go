// Package controllers controllers/members_controllers.go
package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sample/src/models"
	"strconv"
)

var memberRepository *models.MemberRepository

func init() {
	var err error

	memberRepository, err = models.NewMember()
	if err != nil {
		fmt.Printf("failed to create repository: %v", err)
	}
}

func indexMember(c *gin.Context) {
	members := memberRepository.AllMember()
	c.HTML(http.StatusOK, "members/index.tmpl", gin.H{"members": members})
}

func showMember(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	member := memberRepository.FindMember(id)
	c.HTML(http.StatusOK, "members/show.tmpl", gin.H{
		"title": "show",
		"name":  member.Name,
		"id":    member.Id,
	})
}

func createMember(c *gin.Context) {
	name := c.PostForm("name")
	ret := memberRepository.CreateMember(name)
	if ret != nil {
		var errorMessages []string
		errorMessages = append(errorMessages, ret.Error())
		c.HTML(http.StatusBadRequest, "members/new.tmpl", gin.H{"errorMessages": errorMessages})
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
	memberRepository.DeleteMember(id)
	c.JSON(200, gin.H{
		"id": id,
	})
}

func randomMember(c *gin.Context) {
	member := memberRepository.RandomMember()
	c.JSON(200, gin.H{
		"id":   member.Id,
		"name": member.Name,
	})
}
