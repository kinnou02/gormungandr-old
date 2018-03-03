package main

import (
    "github.com/gin-gonic/gin"
    "github.com/kinnou02/gormungandr/handlers"
    "github.com/kinnou02/gormungandr/journeys"
    "github.com/gin-gonic/contrib/ginrus"
    "github.com/sirupsen/logrus"
    "time"
    "flag"
    "os"
)


func setupRouter() *gin.Engine {
    r := gin.New()
    // Recovery middleware recovers from any panics and writes a 500 if there was one.
    r.Use(gin.Recovery())

    r.Use(ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, false))

    r.GET("/status", handlers.Index)


    return r
}

func init_log(logjson bool) {
  // Log as JSON instead of the default ASCII formatter.
  if logjson {
    logrus.SetFormatter(&logrus.JSONFormatter{})
  }

  // Output to stdout instead of the default stderr
  // Can be any io.Writer, see below for File example
  logrus.SetOutput(os.Stdout)

  // Only log the warning severity or above.
  logrus.SetLevel(logrus.DebugLevel)
}

func main() {
    listen := flag.String("listen", ":8080", "[IP]:PORT to listen")
    timeout := flag.Duration("timeout", time.Second, "timeour for call to kraken")
    logjson := flag.Bool("logjson", false, "enable json logging")
    flag.Parse()
    init_log(*logjson)
    logrus.Info("timeout: %s", *timeout)

    r := setupRouter()
    r.GET("/journeys", journeys.JourneysHandler(*timeout))
    // Listen and Server in 0.0.0.0:8080
    r.Run(*listen)
}
