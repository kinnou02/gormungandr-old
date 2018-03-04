package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"github.com/kinnou02/gonavitia"
	"github.com/kinnou02/gonavitia/pbnavitia"
	"github.com/kinnou02/gonavitia/serializer"
	log "github.com/sirupsen/logrus"
)

type JourneysRequest struct {
	From     string    `form:"from" binding:"required"`
	To       string    `form:"to" binding:"required"`
	Datetime time.Time `form:"datetime" time_format:"20060102150405"`
}

func JourneysHandler(kraken *gonavitia.Kraken) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var request JourneysRequest
		if err := c.ShouldBindQuery(&request); err != nil {
			log.Errorf("FATAL: %+v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		pb_req := BuildRequest(request)
		resp, err := kraken.Call(pb_req)
		if err != nil {
			log.Errorf("FATAL: %+v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		r := serializer.NewJourneysReponse(resp)
		//    fmt.Println(resp)
		c.JSON(http.StatusOK, r)
	}
	return gin.HandlerFunc(fn)
}

func BuildRequest(req JourneysRequest) *pbnavitia.Request {
	j := &pbnavitia.JourneysRequest{
		Origin: []*pbnavitia.LocationContext{{
			Place:          proto.String(req.From),
			AccessDuration: proto.Int32(0)},
		},
		Destination: []*pbnavitia.LocationContext{{
			Place:          proto.String(req.To),
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
		StreetnetworkParams: &pbnavitia.StreetNetworkParams{
			OriginMode:             proto.String("walking"),
			DestinationMode:        proto.String("walking"),
			WalkingSpeed:           proto.Float64(1.11),
			BikeSpeed:              proto.Float64(1.11),
			BssSpeed:               proto.Float64(1.11),
			CarSpeed:               proto.Float64(1.11),
			MaxWalkingDurationToPt: proto.Int32(30 * 60),
			MaxBikeDurationToPt:    proto.Int32(30 * 60),
			MaxBssDurationToPt:     proto.Int32(30 * 60),
			MaxCarDurationToPt:     proto.Int32(30 * 60),
		},
	}
	pb_req := &pbnavitia.Request{
		RequestedApi: pbnavitia.API_PLANNER.Enum(),
		Journeys:     j,
	}

	return pb_req
}
