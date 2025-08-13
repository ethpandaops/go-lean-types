package phase0

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/pkg/errors"
)

// Signature is a signature.
type Signature [32]byte

// SignatureLength is the number of bytes in a signature.
const SignatureLength = 32

var (
	emptySignature = Signature{}
)

// IsZero returns true if the signature is zero.
func (s Signature) IsZero() bool {
	return bytes.Equal(s[:], emptySignature[:])
}

// String returns a string version of the structure.
func (s Signature) String() string {
	return fmt.Sprintf("%#x", s)
}

// Format formats the signature.
func (s Signature) Format(state fmt.State, v rune) {
	format := string(v)
	switch v {
	case 's':
		fmt.Fprint(state, s.String())
	case 'x', 'X':
		if state.Flag('#') {
			format = "#" + format
		}
		fmt.Fprintf(state, "%"+format, s[:])
	default:
		fmt.Fprintf(state, "%"+format, s[:])
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (s *Signature) UnmarshalJSON(input []byte) error {
	if len(input) == 0 {
		return errors.New("input missing")
	}

	if !bytes.HasPrefix(input, []byte{'"', '0', 'x'}) {
		return errors.New("invalid prefix")
	}
	if !bytes.HasSuffix(input, []byte{'"'}) {
		return errors.New("invalid suffix")
	}
	if len(input) != 1+2+SignatureLength*2+1 {
		return errors.New("incorrect length")
	}

	length, err := hex.Decode(s[:], input[3:3+SignatureLength*2])
	if err != nil {
		return errors.Wrapf(err, "invalid value %s", string(input[3:3+SignatureLength*2]))
	}

	if length != SignatureLength {
		return errors.New("incorrect length")
	}

	return nil
}

// MarshalJSON implements json.Marshaler.
func (s Signature) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%#x"`, s)), nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (s *Signature) UnmarshalYAML(input []byte) error {
	if len(input) == 0 {
		return errors.New("input missing")
	}

	if !bytes.HasPrefix(input, []byte{'\'', '0', 'x'}) {
		return errors.New("invalid prefix")
	}
	if !bytes.HasSuffix(input, []byte{'\''}) {
		return errors.New("invalid suffix")
	}
	if len(input) != 1+2+SignatureLength*2+1 {
		return errors.New("incorrect length")
	}

	length, err := hex.Decode(s[:], input[3:3+SignatureLength*2])
	if err != nil {
		return errors.Wrapf(err, "invalid value %s", string(input[3:3+SignatureLength*2]))
	}

	if length != SignatureLength {
		return errors.New("incorrect length")
	}

	return nil
}

// MarshalYAML implements yaml.Marshaler.
func (s Signature) MarshalYAML() ([]byte, error) {
	return []byte(fmt.Sprintf(`'%#x'`, s)), nil
}
