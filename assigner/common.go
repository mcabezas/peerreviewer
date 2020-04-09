package assigner

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"peerreviewer/team"
	"peerreviewer/team/member"
)

func testEveryMemberHasAssignedReviewers(t *testing.T, assigner Assigner) {
	teamA := setupTeam()
	assignations := sliceToMap(assigner.DistributeAssignations(teamA))
	for _, m := range teamA.Members {
		assert.NotNil(t, assignations[m])
	}
}

func testEveryMemberHasAtLeastOneExperiencedReviewer(t *testing.T, assigner Assigner) {
	teamA := setupTeam()
	assignations := assigner.DistributeAssignations(teamA)
	for _, m := range assignations {
		hasExperiencedReviewer := false
		for _, r := range m.Reviewers {
			if r.Category == member.Experienced {
				hasExperiencedReviewer = true
				break
			}
		}
		if !hasExperiencedReviewer {
			assert.Fail(t, "no experienced reviewer")
		}
	}
}

func setupTeam() *team.Team {
	memberA := member.NewMember("random1", member.WithCategory(member.NeedsCoaching))
	memberB := member.NewMember("pro1", member.WithCategory(member.Experienced))
	memberC := member.NewMember("random2", member.WithCategory(member.NeedsCoaching))
	memberD := member.NewMember("pro2", member.WithCategory(member.Experienced))
	teamA := team.New("Tiburones", memberA, memberB, memberC, memberD)
	return teamA
}

func sliceToMap(slice []*MemberReviewers) map[*member.Member]*MemberReviewers {
	memberMap := map[*member.Member]*MemberReviewers{}
	for _, m := range slice {
		memberMap[m.Member] = m
	}
	return memberMap
}