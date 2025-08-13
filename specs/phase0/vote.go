package phase0

type Vote struct {
	ValidatorId uint64     `json:"validator_id"`
	Slot        Slot       `json:"slot"`
	Head        Checkpoint `json:"head"`
	Target      Checkpoint `json:"target"`
	Source      Checkpoint `json:"source"`
}
