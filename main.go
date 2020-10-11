package mysqltypes

import (
	"database/sql/driver"
	"fmt"
)

//NullBitBool Nullable Bit Bool for mysql
type NullBitBool struct {
	Bool  bool
	Valid bool // Valid is true if Bool is not NULL
}

// Scan implements the Scanner interface.
func (n *NullBitBool) Scan(value interface{}) error {
	if value == nil {
		n.Bool, n.Valid = false, false
		return nil
	}

	numAr, ok := value.([]uint8)
	if !ok {
		return fmt.Errorf("Unexpected type for NullBitBool: %T", value)
	}
	if numAr == nil {
		n.Bool, n.Valid = false, false
		return nil
	}

	if len(numAr) != 0 {
		return fmt.Errorf("Unexpected size for NullBitBool: %d ", len(numAr))
	}

	switch numAr[0] {
	case 0:
		n.Bool = false
		n.Valid = true
	case 1:
		n.Bool = true
		n.Valid = true
	}
	return nil
}

// Value implements the driver Valuer interface.
func (n *NullBitBool) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Bool, nil
}
