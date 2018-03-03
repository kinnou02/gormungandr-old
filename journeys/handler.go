package journeys

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/kinnou02/gonavitia/serializer"
    zmq "github.com/pebbe/zmq2"
    "github.com/kinnou02/pbnavitia"
    "github.com/golang/protobuf/proto"
    "time"
    "errors"
    log "github.com/sirupsen/logrus"
)

func JourneysHandler(kraken string, timeout time.Duration) gin.HandlerFunc {
    fn := func(c *gin.Context) {
        req := BuildRequest(c.Query("from"), c.Query("to"))
        resp, err := CallKraken(kraken, req, timeout)
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

func CallKraken(kraken string, request *pbnavitia.Request,
                timeout time.Duration) (*pbnavitia.Response, error){
    requester, _ := zmq.NewSocket(zmq.REQ)
    requester.Connect(kraken)
    defer requester.Close()
    data, _ := proto.Marshal(request)
    requester.Send(string(data), 0)
    poller := zmq.NewPoller()
    poller.Add(requester, zmq.POLLIN)
    p, err := poller.Poll(timeout)
    if err != nil {
        return nil, err
    }
    if len(p) < 1 {
        return nil, errors.New("fucked")
    }
    raw_resp, _ := p[0].Socket.Recv(0)
    resp := &pbnavitia.Response{}
    _ = proto.Unmarshal([]byte(raw_resp), resp)
    return resp, nil
}



func BuildRequest(from, to string) *pbnavitia.Request{
    j := &pbnavitia.JourneysRequest{
        Origin: []*pbnavitia.LocationContext{{
            Place: proto.String(from),
            AccessDuration: proto.Int32(0)},
        },
        Destination: []*pbnavitia.LocationContext{{
            Place: proto.String(to),
            AccessDuration: proto.Int32(0)},
        },
        Datetimes: []uint64{uint64(time.Now().Unix())},
        Clockwise: proto.Bool(true),
        MaxDuration: proto.Int(86400),
        MaxTransfers: proto.Int(10),
        Wheelchair: proto.Bool(false),
        ShowCodes: proto.Bool(false),
        RealtimeLevel: pbnavitia.RTLevel_BASE_SCHEDULE.Enum(),
        MaxExtraSecondPass: proto.Int(0),
        WalkingTransferPenalty: proto.Int(120),
        DirectPathDuration: proto.Int(30*60),
        BikeInPt: proto.Bool(false),
    }
    req := &pbnavitia.Request{
        RequestedApi: pbnavitia.API_PLANNER.Enum(),
        Journeys: j,
    }


    return req
}
