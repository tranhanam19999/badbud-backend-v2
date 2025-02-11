package model

type MatchRequestStatus string

type MatchRequest struct {
	Base
	UserID  string
	MatchID string
	Status  MatchRequestStatus
	User    User  `gorm:"foreignKey:UserID"`
	Match   Match `gorm:"foreignKey:MatchID"`
}

var (
	MatchRequestStatusRequested MatchRequestStatus = "REQUESTED"
	MatchRequestStatusRejected  MatchRequestStatus = "REJECTED"
	MatchRequestStatusAccepted  MatchRequestStatus = "ACCEPTED"
)
