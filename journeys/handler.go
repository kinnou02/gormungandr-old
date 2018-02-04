package journeys

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/kinnou02/gormungandr/response"
)

func JourneysHandler(c *gin.Context) {
    r := response.JourneysResponse{[]response.Journey{
        response.Journey{
            From: response.Place{"place:1", "my place"},
            To: response.Place{"place:2", "some other place"},
            Sections: []response.Section{
                {
                    From: response.Place{"place:1", "my place"},
                    To: response.Place{"place:2", "some other place"},
                },
            },
        },
    }}
    c.JSON(http.StatusOK, r)
}
