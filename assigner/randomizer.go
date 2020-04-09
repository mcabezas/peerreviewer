package assigner

import (
	"math/rand"
	"sync"

	"peerreviewer/team/member"
)

type Randomizer struct {
	allMembers    []member.Member
	remaindersPtr int
	mu            *sync.Mutex
}

func NewRandomizer(members []member.Member) *Randomizer {
	return &Randomizer{
		mu:            &sync.Mutex{},
		remaindersPtr: len(members),
		allMembers:    members,
	}
}

func (r *Randomizer) Next(skipMemberID string, ) *member.Member {

	//if there is no more remainders then starts a new cycle
	if r.remaindersPtr == 0 {
		r.remaindersPtr = len(r.allMembers)
	}

	// Getting random member
	randomIdx := rand.Int() % r.remaindersPtr

	// Moving the next to the end of slice
	// remainders pointer was decreased in order to dont take care
	// about the next
	r.mu.Lock()
	next := r.allMembers[randomIdx]
	r.allMembers[randomIdx] = r.allMembers[len(r.allMembers)-1]
	r.allMembers[len(r.allMembers)-1] = next
	r.remaindersPtr--
	r.mu.Unlock()

	if next.ID == skipMemberID {
		return r.Next(skipMemberID)
	}
	return &next
}
