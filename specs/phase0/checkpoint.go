package phase0

type Checkpoint struct {
	Root Root `json:"root"`
	Slot Slot `json:"slot"`
}
