// Code generated by "github.com/kade-chen/library"
// DO NOT EDIT

package vertex

import (
	"bytes"
	"fmt"
	"strings"
)

// ParseGRANT_MODELFromString Parse GRANT_MODEL from string
func ParseGRANT_MODELFromString(str string) (GRANT_MODEL, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := GRANT_MODEL_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown GRANT_MODEL: %s", str)
	}

	return GRANT_MODEL(v), nil
}

// Equal type compare
func (t GRANT_MODEL) Equal(target GRANT_MODEL) bool {
	return t == target
}

// IsIn todo
func (t GRANT_MODEL) IsIn(targets ...GRANT_MODEL) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t GRANT_MODEL) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *GRANT_MODEL) UnmarshalJSON(b []byte) error {
	ins, err := ParseGRANT_MODELFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}
