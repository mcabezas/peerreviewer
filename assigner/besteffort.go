package assigner

import (
	"peerreviewer/team"
	"peerreviewer/team/member"
)

type BestEffort struct {
	config *Config
}

func NewBestEffort(config *Config) Assigner {
	return &BestEffort{config: config}
}

func (a *BestEffort) DistributeAssignations(t *team.Team) []*MemberReviewers {
	// Splitting members base on the experience they have
	var experiencedMembers []member.Member
	var nonExperiencedMembers []member.Member
	for _, m := range t.Members {
		if m.Category == member.Experienced {
			experiencedMembers = append(experiencedMembers, *m)
			continue
		}
		nonExperiencedMembers = append(nonExperiencedMembers, *m)
	}
	expRandomizer := NewRandomizer(experiencedMembers)
	nonExpRandomizer := NewRandomizer(nonExperiencedMembers)

	// Splitting assignations
	assignations := make([]*MemberReviewers, len(t.Members))
	for ii, m := range t.Members {
		assignation := &MemberReviewers{Member: m}
		assignation.Reviewers = make([]*member.Member, a.config.MaxAssignersPerMember)
		for jj := 0; jj < a.config.MaxAssignersPerMember; jj++ {
			if jj < a.config.MinExperiencedReviewers {
				assignation.Reviewers[jj] = expRandomizer.Next(m.ID)
				continue
			}
			assignation.Reviewers[jj] = nonExpRandomizer.Next(m.ID)
		}
		assignations[ii] = assignation
	}
	return assignations
}

