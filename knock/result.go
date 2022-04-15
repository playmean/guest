package knock

type StatusFlag string

const (
	StatusFlagGood    StatusFlag = "good"
	StatusFlagWarning StatusFlag = "warning"
	StatusFlagBad     StatusFlag = "bad"
)

type ResultEntity struct {
	Title  string      `json:"title"`
	Entity interface{} `json:"result"`
}

type ResultStatus struct {
	Title string       `json:"title"`
	Value string       `json:"value"`
	Flags []StatusFlag `json:"flags,omitempty"`
	// TODO implement extended status information
}

type Result struct {
	Entities []ResultEntity `json:"entities"`
	Statuses []ResultStatus `json:"statuses"`
}
