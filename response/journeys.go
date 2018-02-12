package response

import "github.com/kinnou02/gormungandr/navitia"
import "time"

type JourneysResponse struct{
    Journeys []*Journey `json:"journeys"`
}

func NewJourneysReponse(pb *pbnavitia.Response) JourneysResponse{
    r := JourneysResponse{}
    for _, pb_journey := range pb.Journeys {
        r.Journeys = append(r.Journeys, NewJourney(pb_journey))
    }
    return r
}

type Journey struct{
    From *Place `json:"from,omitempty"`
    To *Place `json:"to,omitempty"`
    Duration int32 `json:"duration"`
    NbTransfers int32 `json:"nb_transfers"`
    DepartureDateTime time.Time `json:"departure_date_time"`
    ArrivalDateTime time.Time `json:"arrival_date_time"`
    RequestedDateTime time.Time `json:"requested_date_time"`
    Type *string `json:"type"`
    Tags []string `json:"tags"`
    Sections []*Section `json:"sections"`

}

func NewJourney(pb *pbnavitia.Journey) *Journey{
    journey := Journey{
        From: NewPlace(pb.Origin),
        To: NewPlace(pb.Destination),
        Duration: pb.GetDuration(),
        NbTransfers: pb.GetNbTransfers(),
        DepartureDateTime: time.Unix(int64(pb.GetDepartureDateTime()), 0),
        ArrivalDateTime: time.Unix(int64(pb.GetArrivalDateTime()), 0),
        RequestedDateTime: time.Unix(int64(pb.GetRequestedDateTime()), 0),
    }
    for _, pb_section := range pb.Sections {
        journey.Sections = append(journey.Sections, NewSection(pb_section))
    }
    return &journey
}

type Section struct{
    From *Place `json:"from,omitempty"`
    To *Place `json:"to,omitempty"`
}

func NewSection(pb *pbnavitia.Section) *Section{
    section := Section{
        From: NewPlace(pb.Origin),
        To: NewPlace(pb.Destination),
    }
    return &section
}
