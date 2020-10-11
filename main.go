package mysqltypes

import (
	"database/sql/driver"
	"fmt"
	"log"
)

//NullBitBool Nullable Bit Bool for mysql.
// This is a nullable struct that can be used whern a bit(1) type in
// mysql is nullable
type NullBitBool struct {
	Bool  bool
	Valid bool
}

// Scan implements the Scanner interface.
func (n *NullBitBool) Scan(value interface{}) error {
	if value == nil {
		n.Bool, n.Valid = false, false
		return nil
	}

	numAr, ok := value.([]uint8)
	if !ok {
		//The returns the type that was given so that the user can use the correct type
		return fmt.Errorf("Unexpected type for NullBitBool: %T", value)
	}

	if numAr == nil {
		n.Bool, n.Valid = false, false
		return nil
	}

	//if this occurs you should use the NullBitBoolArray type
	if len(numAr) != 1 {
		return fmt.Errorf("Unexpected size for NullBitBool: %d ", len(numAr))
	}

	log.Printf("numberar is %d", numAr[0])
	switch numAr[0] {
	case 0:
		n.Bool = false
		n.Valid = true
		break
	case 1:
		n.Bool = true
		n.Valid = true
		break
	default:
		return fmt.Errorf("Unexpected value in field %d", numAr[0])
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
