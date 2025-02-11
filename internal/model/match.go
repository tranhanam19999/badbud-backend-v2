package model

type Match struct {
	Base
	Name string
	// Id of court
	CourtID string
	// Number of court on court (court 1, 2, 3,...)
	CourtNum  int
	FeeMale   string
	FeeFemale string
	// Limit on how many person per court
	Limit             int
	MatchParticipants []*MatchParticipant
}

type MatchParticipant struct {
	Base
	MatchId string
	UserId  string
}
