package base

type stats struct {
	Running int32 `json:"running"`
}

//monitor container stats
var Statics = &stats{}
