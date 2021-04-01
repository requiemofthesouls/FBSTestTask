package models

type FibonacciRequest struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

type FibonacciResponse struct {
	Result string `json:"result"`
}
