package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/", top)
	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, gin.Version)
	})
	member := r.Group("/members")
	{
		member.GET("/new", newMember)
		member.GET("/:id", showMember)
		member.DELETE("/:id", deleteMember)
		member.GET("", indexMember)
		member.POST("", createMember)
		member.GET("/random", randomMember)
	}
	team := r.Group("/teams")
	{
		team.GET("/new", newTeam)
		team.GET("", indexTeam)
		team.GET("/:id/member", showTeamMember)
		team.GET("/:id/member/random", randomTeamMember)
		team.DELETE("/:id", deleteTeam)
		team.POST("", createTeam)
	}
}

func top(c *gin.Context) {
	c.HTML(http.StatusOK, "top/top.tmpl", gin.H{})
}
