package msgraph

type TeamResponse struct {
	OdataContext string  `json:"@odata.context"`
	OdataCount   int     `json:"@odata.count"`
	Teams        []Teams `json:"value"`

	graphClient *GraphClient
}
type Teams struct {
	ID                          string      `json:"id"`
	CreatedDateTime             interface{} `json:"createdDateTime"`
	DisplayName                 string      `json:"displayName"`
	Description                 string      `json:"description"`
	InternalID                  interface{} `json:"internalId"`
	Classification              interface{} `json:"classification"`
	Specialization              interface{} `json:"specialization"`
	Visibility                  string      `json:"visibility"`
	WebURL                      interface{} `json:"webUrl"`
	IsArchived                  bool        `json:"isArchived"`
	IsMembershipLimitedToOwners interface{} `json:"isMembershipLimitedToOwners"`
	MemberSettings              interface{} `json:"memberSettings"`
	GuestSettings               interface{} `json:"guestSettings"`
	MessagingSettings           interface{} `json:"messagingSettings"`
	FunSettings                 interface{} `json:"funSettings"`
	DiscoverySettings           interface{} `json:"discoverySettings"`
}

func (t *TeamResponse) setGraphClient(gC *GraphClient) {
	t.graphClient = gC
}
