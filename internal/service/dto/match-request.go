package dto

type CreateMatchRequestReq struct {
	MatchId string `json:"match_id"`
}
type CreateMatchRequestResp struct{}

type AcceptMatchRequestReq struct {
	RequestId string `json:"request_id"`
}

type AcceptMatchRequestReesp struct{}

type RejectMatchRequestReq struct {
	RequestId string `json:"request_id"`
	// TODO: Reject reason
	// Reason string `json:"reason"`
}

type RejectMatchRequestResp struct{}
