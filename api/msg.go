package api

const (
	SUCCESS = 0
	UNKNOWN = iota + 1000
	FAILED
)

var ResMessage = map[int]string{
	SUCCESS: "OK",
	FAILED:  "Failed",
	UNKNOWN: "Unknown Error",
}
