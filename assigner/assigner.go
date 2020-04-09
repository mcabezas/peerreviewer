package assigner

import (
	"peerreviewer/team"
	"peerreviewer/team/member"
)

type Config struct {
	MaxAssignersPerMember   int
	MinExperiencedReviewers int
}

type Backend int
type constructor func(config *Config) Assigner

const(
	BestEffortBackend = iota
)

var backends = map[Backend]constructor {BestEffortBackend: NewBestEffort}

type MemberReviewers struct {
	Member    *member.Member
	Reviewers []*member.Member
}

type Assigner interface {
	DistributeAssignations(team *team.Team) []*MemberReviewers
}

func New(config *Config, backend Backend) Assigner {
	return backends[backend](config)
}
