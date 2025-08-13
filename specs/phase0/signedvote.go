package phase0

type SignedVote struct {
	Vote      Vote      `json:"vote"`
	Signature Signature `json:"signature"`
}
