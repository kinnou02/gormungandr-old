package response

type JourneysResponse struct{
    Journeys []Journey `json:"journeys"`
}

type Journey struct{
    From Place
    To Place
    Sections []Section
}

type Section struct{
    From Place
    To Place
}
