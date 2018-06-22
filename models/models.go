package models

// Result Result of an arithmetic call.
type Result struct {
	Action string `json:"action"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
	Answer int    `json:"answer"`
	Cached bool   `json:"cached"`
	Error  string `json:"error"` //I've added this to ensure that we can inform clients of mistakes.
}
