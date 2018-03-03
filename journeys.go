package gormungandr

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"github.com/kinnou02/gonavitia"
	"github.com/kinnou02/gonavitia/serializer"
	"github.com/kinnou02/pbnavitia"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func JourneysHandler(kraken *gonavitia.Kraken) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		req := BuildRequest(c.Query("from"), c.Query("to"))
		resp, err := kraken.Call(req)
		if err != nil {
			log.Error("failure to call kraken %s", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		r := serializer.NewJourneysReponse(resp)
		//    fmt.Println(resp)
		c.JSON(http.StatusOK, r)
	}
	return gin.HandlerFunc(fn)
}

func BuildRequest(from, to string) *pbnavitia.Request {
	j := &pbnavitia.JourneysRequest{
		Origin: []*pbnavitia.LocationContext{{
			Place:          proto.String(from),
			AccessDuration: proto.Int32(0)},
		},
		Destination: []*pbnavitia.LocationContext{{
			Place:          proto.String(to),
			AccessDuration: proto.Int32(0)},
		},
		Datetimes:              []uint64{uint64(time.Now().Unix())},
		Clockwise:              proto.Bool(true),
		MaxDuration:            proto.Int(86400),
		MaxTransfers:           proto.Int(10),
		Wheelchair:             proto.Bool(false),
		ShowCodes:              proto.Bool(false),
		RealtimeLevel:          pbnavitia.RTLevel_BASE_SCHEDULE.Enum(),
		MaxExtraSecondPass:     proto.Int(0),
		WalkingTransferPenalty: proto.Int(120),
		DirectPathDuration:     proto.Int(30 * 60),
		BikeInPt:               proto.Bool(false),
	}
	req := &pbnavitia.Request{
		RequestedApi: pbnavitia.API_PLANNER.Enum(),
		Journeys:     j,
	}

	return req
}