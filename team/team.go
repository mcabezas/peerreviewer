package team

import "peerreviewer/team/member"

type Team struct {
	ID int
	Name string
	Members []*member.Member
}

type Option func(*Team)

func New(name string, members ...*member.Member) *Team {
	return &Team{Members:members}
}
