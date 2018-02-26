package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUint(c *gin.Context, key string) (uint, error) {
	i, err := strconv.Atoi(c.Param(key))
	return uint(i), err
}
