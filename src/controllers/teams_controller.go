package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sample/src/models"
	"sample/src/utility"
	"strconv"
)

func createTeam(c *gin.Context) {
	name := c.PostForm("teamName")
	members := c.PostFormArray("members")
	repo, err := models.NewRepository()
	if err != nil {
		log.Fatalf("failed to create repository: %v", err)
	}

	var memberIds []int
	memberIds, err = utility.ConvertStringSliceToIntSlice(members)
	if err != nil {
		println(err)
	}

	membersList := repo.FindMemberByIds(memberIds)
	repoErr := repo.CreateTeam(name, &membersList)
	if repoErr != nil {
		var errorMessages []string
		errorMessages = append(errorMessages, repoErr.Error())
		c.HTML(http.StatusBadRequest, "teams/new.tmpl", gin.H{})
		return
	}
	c.Redirect(http.StatusFound, "/teams")
}

func newTeam(c *gin.Context) {
	repo, err := models.NewRepository()
	if err != nil {
		log.Fatalf("failed to create repository: %v", err)
	}
	members := repo.AllMember()

	c.HTML(http.StatusOK, "teams/new.tmpl", gin.H{"members": members})
}

func indexTeam(c *gin.Context) {
	repo, err := models.NewRepository()
	if err != nil {
		log.Fatalf("failed to create repository: %v", err)
	}
	teams := repo.AllTeam()
	c.HTML(http.StatusOK, "teams/index.tmpl", gin.H{"teams": teams})
}

func showTeamMember(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	repo, err := models.NewRepository()
	if err != nil {
		log.Fatalf("failed to create repository: %v", err)
	}

	team := repo.FindTeam(id)
	members := repo.TeamMembers(id)

	c.HTML(http.StatusOK, "teams/show.tmpl", gin.H{"team": team, "members": members})
}

func randomTeamMember(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	repo, err := models.NewRepository()
	if err != nil {
		log.Fatalf("failed to create repository: %v", err)
	}
	member := repo.RandomTeamMember(id)

	c.JSON(200, gin.H{
		"id":   member.Id,
		"name": member.Name,
	})
}
