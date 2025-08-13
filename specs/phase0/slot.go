package phase0

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

// Slot is a slot number.
type Slot uint64

// UnmarshalJSON implements json.Unmarshaler.
func (s *Slot) UnmarshalJSON(input []byte) error {
	if len(input) == 0 {
		return errors.New("input missing")
	}
	if len(input) < 3 {
		return errors.New("input malformed")
	}
	if !bytes.HasPrefix(input, []byte{'"'}) {
		return errors.New("invalid prefix")
	}
	if !bytes.HasSuffix(input, []byte{'"'}) {
		return errors.New("invalid suffix")
	}

	val, err := strconv.ParseUint(string(input[1:len(input)-1]), 10, 64)
	if err != nil {
		return errors.Wrapf(err, "invalid value %s", string(input[1:len(input)-1]))
	}
	*s = Slot(val)

	return nil
}

// MarshalJSON implements json.Marshaler.
func (s Slot) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%d"`, s)), nil
}
