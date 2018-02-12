package response

import "github.com/kinnou02/gormungandr/navitia"

type Place struct{
    Id *string
    Name *string
}


func NewPlace(pb *pbnavitia.PtObject) *Place{
    if pb == nil {
        return nil
    }
    place := Place{pb.Uri, pb.Name}
    return &place
}
