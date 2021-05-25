package controllers

import (
	"github.com/aiuzu42/battlelog/pkg/roll"
	"github.com/aiuzu42/battlelog/pkg/stratagems"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func RollController(c *gin.Context) {
	s, err := strconv.Atoi(c.Query("sides"))
	if err != nil || s < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid \"sides\" parameter"})
		return
	}
	n, err := strconv.Atoi(c.Query("number"))
	if err != nil || n < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid \"number\" parameter"})
		return
	}
	c.JSON(http.StatusOK, roll.Roll(n, s))
}

func GetStratagemsController(c *gin.Context) {
	c.JSON(http.StatusOK, stratagems.GetAllStratagemFactions())
}

func GetStratagemByName(c *gin.Context) {
	subs := c.QueryArray("sub")
	name := c.Param("name")
	phases := c.QueryArray("phase")
	for _, p := range phases {
		if !stratagems.ValidatePhase(p) {
			phaseError := "Invalid phase: " + p
			c.JSON(http.StatusBadRequest, gin.H{"error": phaseError})
			return
		}
	}
	f := stratagems.GetFactionStratagems(name, subs, phases)
	if f.Name == "" {
		c.JSON(http.StatusNoContent, f)
		return
	}
	c.JSON(http.StatusOK, f)
}

func GetPhases(c *gin.Context) {
	c.JSON(http.StatusOK, stratagems.GetPhases())
}
