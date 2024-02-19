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

	people := repo.FindPersonByIds(memberIds)
	repoErr := repo.CreateTeam(name, &people)
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
	people := repo.AllPerson()

	c.HTML(http.StatusOK, "teams/new.tmpl", gin.H{"people": people})
}

func indexTeam(c *gin.Context) {
	repo, err := models.NewRepository()
	if err != nil {
		log.Fatalf("failed to create repository: %v", err)
	}
	teams := repo.AllTeam()
	c.HTML(http.StatusOK, "teams/index.tmpl", gin.H{"teams": teams})
}

func showTeamPeople(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	repo, err := models.NewRepository()
	if err != nil {
		log.Fatalf("failed to create repository: %v", err)
	}

	team := repo.FindTeam(id)
	people := repo.TeamPeople(id)

	c.HTML(http.StatusOK, "teams/show.tmpl", gin.H{"team": team, "people": people})
}

func selectTeamMember(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	repo, err := models.NewRepository()
	if err != nil {
		log.Fatalf("failed to create repository: %v", err)
	}
	person := repo.SampleTeamPeople(id)

	c.JSON(200, gin.H{
		"id":   person.Id,
		"name": person.Name,
	})
}
