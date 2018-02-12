package journeys

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/kinnou02/gormungandr/response"
    zmq "github.com/pebbe/zmq4"
    "github.com/kinnou02/gormungandr/navitia"
    "github.com/golang/protobuf/proto"
    "time"

)

func JourneysHandler(c *gin.Context) {

    req := BuildRequest(c.Query("from"), c.Query("to"))
    requester, _ := zmq.NewSocket(zmq.REQ)
    defer requester.Close()
    requester.Connect("tcp://localhost:3000")
    data, _ := proto.Marshal(req)
    requester.Send(string(data), 0)
    raw_resp, _ := requester.Recv(0)
    resp := &pbnavitia.Response{}
    _ = proto.Unmarshal([]byte(raw_resp), resp)
    r := response.NewJourneysReponse(resp)
//    fmt.Println(resp)
    c.JSON(http.StatusOK, r)
}


/*
type JourneysRequest struct {
	Origin                 []*LocationContext   `protobuf:"bytes,1,rep,name=origin" json:"origin,omitempty"`
	Destination            []*LocationContext   `protobuf:"bytes,2,rep,name=destination" json:"destination,omitempty"`
	Datetimes              []uint64             `protobuf:"varint,3,rep,name=datetimes" json:"datetimes,omitempty"`
	Clockwise              *bool                `protobuf:"varint,4,req,name=clockwise" json:"clockwise,omitempty"`
	ForbiddenUris          []string             `protobuf:"bytes,5,rep,name=forbidden_uris" json:"forbidden_uris,omitempty"`
	MaxDuration            *int32               `protobuf:"varint,6,req,name=max_duration" json:"max_duration,omitempty"`
	MaxTransfers           *int32               `protobuf:"varint,7,req,name=max_transfers" json:"max_transfers,omitempty"`
	StreetnetworkParams    *StreetNetworkParams `protobuf:"bytes,8,opt,name=streetnetwork_params" json:"streetnetwork_params,omitempty"`
	Wheelchair             *bool                `protobuf:"varint,9,opt,name=wheelchair,def=0" json:"wheelchair,omitempty"`
	ShowCodes              *bool                `protobuf:"varint,11,opt,name=show_codes" json:"show_codes,omitempty"`
	Details                *bool                `protobuf:"varint,13,opt,name=details" json:"details,omitempty"`
	RealtimeLevel          *RTLevel             `protobuf:"varint,14,opt,name=realtime_level,enum=pbnavitia.RTLevel" json:"realtime_level,omitempty"`
	MaxExtraSecondPass     *int32               `protobuf:"varint,15,opt,name=max_extra_second_pass,def=0" json:"max_extra_second_pass,omitempty"`
	WalkingTransferPenalty *int32               `protobuf:"varint,16,opt,name=walking_transfer_penalty,def=120" json:"walking_transfer_penalty,omitempty"`
	DirectPathDuration     *int32               `protobuf:"varint,17,opt,name=direct_path_duration" json:"direct_path_duration,omitempty"`
	BikeInPt               *bool                `protobuf:"varint,18,opt,name=bike_in_pt" json:"bike_in_pt,omitempty"`
	AllowedId              []string             `protobuf:"bytes,19,rep,name=allowed_id" json:"allowed_id,omitempty"`
	XXX_unrecognized       []byte               `json:"-"`
*/

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
