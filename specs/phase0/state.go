package phase0

// Bitlist represents a bitlist with a maximum length.
type Bitlist []byte

// State represents the beacon chain state container.
type State struct {
	Config Config `json:"config"`

	LatestJustified Checkpoint `json:"latest_justified"`
	LatestFinalized Checkpoint `json:"latest_finalized"`

	HistoricalBlockHashes []Root `json:"historical_block_hashes" dynssz-max:"HISTORICAL_ROOTS_LIMIT" ssz-max:"262144"`
	JustifiedSlots        []bool `json:"justified_slots" dynssz-max:"HISTORICAL_ROOTS_LIMIT" ssz-max:"262144"`

	JustificationsRoots      []Root `json:"justifications_roots" dynssz-max:"HISTORICAL_ROOTS_LIMIT" ssz-max:"262144"`
	JustificationsValidators []byte `json:"justifications_validators" dynssz-max:"(HISTORICAL_ROOTS_LIMIT * VALIDATOR_REGISTRY_LIMIT) / 8" ssz-max:"134217728"`
}
