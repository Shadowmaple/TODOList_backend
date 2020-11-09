package main

import (
	"net/http"

	"todolist_backend/log"
	"todolist_backend/router"

	"github.com/gin-gonic/gin"
)

const (
	MaxPingCount = 10
)

func main() {
	g := gin.Default()

	router.Load(g)

	log.Info(http.ListenAndServe("localhost:8080", g).Error())
}
