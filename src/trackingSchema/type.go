package trackingSchema

type (
	//Schema tracking data Schema
	Schema struct {
		Alias   string        `json:"alias"`
		EventID string        `json:"eventID"`
		Desc    string        `json:"description"`
		Params  []*Parameters `json:"parameters"`
	}

	//Parameters parameters of event
	Parameters struct {
		Name        string `json:"name"`
		DataType    string `json:"type"`
		Format      string `json:"format"`
		Description string `json:"description"`
		Mandatory   bool   `json:"mandatory"`
	}
)
