package participation

import "github.com/gohornet/hornet/pkg/model/milestone"

// EventsResponse defines the response of a GET RouteParticipationEvents REST API call.
type EventsResponse struct {
	// The hex encoded IDs of the found events.
	EventIDs []string `json:"eventIds"`
}

// CreateEventResponse defines the response of a POST RouteParticipationEvents REST API call.
type CreateEventResponse struct {
	// The hex encoded ID of the created participation event.
	EventID string `json:"eventId"`
}

// TrackedParticipation holds the information for each tracked participation.
type TrackedParticipation struct {
	MessageID           string          `json:"messageId"`
	Amount              uint64          `json:"amount"`
	StartMilestoneIndex milestone.Index `json:"startMilestoneIndex"`
	EndMilestoneIndex   milestone.Index `json:"endMilestoneIndex"`
}

// OutputStatusResponse defines the response of a GET RouteOutputStatus REST API call.
type OutputStatusResponse struct {
	Participations map[string]*TrackedParticipation `json:"participations"`
}

// AddressReward holds the amount and token symbol for a certain reward.
type AddressReward struct {
	Amount         uint64 `json:"amount"`
	Symbol         string `json:"symbol"`
	MinimumReached bool   `json:"minimumReached"`
}

// AddressRewardsResponse defines the response of a GET RouteAddressBech32Status or RouteAddressEd25519Status REST API call.
type AddressRewardsResponse struct {
	Rewards map[string]*AddressReward `json:"rewards"`
}

// RewardsResponse defines the response of a GET RouteAdminRewards REST API call and contains the rewards for each address.
type RewardsResponse struct {
	Symbol         string            `json:"symbol"`
	MilestoneIndex milestone.Index   `json:"milestoneIndex"`
	TotalRewards   uint64            `json:"totalRewards"`
	Checksum       string            `json:"checksum"`
	Rewards        map[string]uint64 `json:"rewards"`
}

// ParticipationsResponse defines the response of a GET RouteAdminActiveParticipations or RouteAdminPastParticipations REST API call.
type ParticipationsResponse struct {
	Participations map[string]*TrackedParticipation `json:"participations"`
}