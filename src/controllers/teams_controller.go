package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sample/src/models"
	"sample/src/utility"
	"strconv"
)

var teamRepository *models.TeamRepository

func init() {
	var err error
	teamRepository, err = models.NewTeam()
	if err != nil {
		fmt.Printf("failed to create repository: %v", err)
	}
}
func createTeam(c *gin.Context) {
	var memberIds []int
	var err error
	name := c.PostForm("teamName")
	members := c.PostFormArray("members")
	memberIds, err = utility.ConvertStringSliceToIntSlice(members)

	if err != nil {
		println(err)
	}

	membersList := memberRepository.FindMemberByIds(memberIds)
	ret := teamRepository.CreateTeam(name, &membersList)
	if ret != nil {
		var errorMessages []string
		errorMessages = append(errorMessages, ret.Error())
		c.HTML(http.StatusBadRequest, "teams/new.tmpl", gin.H{"errorMessages": errorMessages})
		return
	}
	c.Redirect(http.StatusFound, "/teams")
}

func newTeam(c *gin.Context) {
	members := memberRepository.AllMember()

	c.HTML(http.StatusOK, "teams/new.tmpl", gin.H{"members": members})
}

func indexTeam(c *gin.Context) {
	teams := teamRepository.AllTeam()
	c.HTML(http.StatusOK, "teams/index.tmpl", gin.H{"teams": teams})
}

func showTeamMember(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	team := teamRepository.FindTeam(id)
	members, _ := memberRepository.TeamMembers(id)

	c.HTML(http.StatusOK, "teams/show.tmpl", gin.H{"team": team, "members": members})
}

func deleteTeam(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := teamRepository.DeleteTeam(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{})
}

func randomTeamMember(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	member := memberRepository.RandomTeamMember(id)

	c.JSON(200, gin.H{
		"id":   member.Id,
		"name": member.Name,
	})
}
