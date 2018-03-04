package main

import (
	"flag"
	"os"
	"time"

	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/gin"
	"github.com/kinnou02/gonavitia"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

func setupRouter() *gin.Engine {
	r := gin.New()
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	r.Use(ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, false))

	r.GET("/status", Index)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

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
	kraken_addr := flag.String("kraken", "tcp://localhost:30000", "zmq addr for kraken")
	logjson := flag.Bool("logjson", false, "enable json logging")
	flag.Parse()
	init_log(*logjson)

	kraken := gonavitia.NewKraken("default", *kraken_addr, *timeout)

	r := setupRouter()
	r.GET("/journeys", JourneysHandler(kraken))
	// Listen and Server in 0.0.0.0:8080
	err := r.Run(*listen)
	if err != nil {
		logrus.Errorf("failure to start: %+v", err)
	}
}
