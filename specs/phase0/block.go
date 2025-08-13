package phase0

type Block struct {
	Slot      Slot   `json:"slot"`
	Parent    Root   `json:"parent"`
	Votes     []Vote `json:"votes" dynssz-max:"VALIDATOR_REGISTRY_LIMIT" ssz-max:"4096"`
	StateRoot Root   `json:"state_root"`
}
