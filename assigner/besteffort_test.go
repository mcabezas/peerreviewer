package assigner

import (
	"testing"
)

func TestEveryMemberHasAssignedReviewersBestEffort(t *testing.T) {
	assigner := New(&Config{MaxAssignersPerMember: 2, MinExperiencedReviewers: 1}, BestEffortBackend)
	testEveryMemberHasAssignedReviewers(t, assigner)
}

func TestEveryMemberHasAtLeastOneExperiencedReviewerBestEffort(t *testing.T) {
	assigner := New(&Config{MaxAssignersPerMember: 2, MinExperiencedReviewers: 1}, BestEffortBackend)
	testEveryMemberHasAtLeastOneExperiencedReviewer(t, assigner)
}