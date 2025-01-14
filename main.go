package main

import (
  "fmt"
  "log"
  "net/http"
  "os"

  "github.com/gin-gonic/gin"
)

func main() {
  log.Println("Starting server...")
  router := gin.New()
  router.GET("/", rootHandler)
  port := os.Getenv("PORT")
  if len(port) == 0 {
    port = "8080"
  }
  router.Run(fmt.Sprintf(":%s", port))
}

func rootHandler(ctx *gin.Context) {
  ctx.String(http.StatusOK, "This is a silly demo")
}
