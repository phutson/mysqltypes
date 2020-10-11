package mysqltypes

import (
	"testing"
)

func TestNullBitBool(t *testing.T) {

	t.Run("NilHandling", func(t *testing.T) {
		var itmUnderTest NullBitBool
		err := itmUnderTest.Scan(nil)
		if err != nil {
			t.Errorf("Should have retruned a nil")
		}
	})
	t.Run("InvalidTyoe", func(t *testing.T) {
		var itmUnderTest NullBitBool
		invalidType := "0"

		err := itmUnderTest.Scan(invalidType)
		if err == nil {
			t.Error("Should of thrown an error")
		}
		if err.Error() != "Unexpected type for NullBitBool: string" {
			t.Errorf("Should of throun an error %s", err.Error())
		}
	})
	t.Run("TestingValidTypeAndValueOfTrue", func(t *testing.T) {
		var itmUnderTest NullBitBool
		var bitBool []uint8
		bitBool = append(bitBool, 1)
		err := itmUnderTest.Scan(bitBool)
		t.Logf("bitbool 0 is %d", bitBool[0])
		if err != nil {
			t.Errorf("Threw an error when it shouldn't have %s", err)
		}
		if !itmUnderTest.Bool {
			t.Error("Should have been a true but it was a false")
		}
	})
	t.Run("TestingSliceLength", func(t *testing.T) {
		var itmUnderTest NullBitBool
		var bitBool []uint8
		bitBool = append(bitBool, 1)
		bitBool = append(bitBool, 0)
		bitBool = append(bitBool, 1)
		err := itmUnderTest.Scan(bitBool)
		if err == nil {
			t.Errorf("It should have thrown an error ")
		}
	})
	t.Run("TestingValidTypeAndValueOfFalse", func(t *testing.T) {
		var itmUnderTest NullBitBool
		var bitBool []uint8
		bitBool = append(bitBool, 0)
		err := itmUnderTest.Scan(bitBool)
		if err != nil {
			t.Errorf("Threw an error when it shouldn't have %s", err)
		}
		if itmUnderTest.Bool {
			t.Error("Should have been a true but it was a false")
		}
	})
	t.Run("TestingValidTypeInvalidValue", func(t *testing.T) {
		var itmUnderTest NullBitBool
		var bitBool []uint8
		bitBool = append(bitBool, 5)
		err := itmUnderTest.Scan(bitBool)
		t.Logf("bitbool 0 is %d", bitBool[0])
		if err == nil {
			t.Errorf("Should have thrown an error %s", err)
		}
	})
	t.Run("TestingValue", func(t *testing.T) {
		var itmUnderTest NullBitBool
		var bitBool []uint8
		bitBool = append(bitBool, 1)
		err := itmUnderTest.Scan(bitBool)
		if err != nil {
			t.Errorf("Threw an error when it shouldn't have %s", err)
		}
		value, err := itmUnderTest.Value()
		outBool, ok := value.(bool)
		if !ok {
			t.Error("returned the wrong type")
		}
		if !outBool {
			t.Error("Should have been a true but it was a false")
		}
	})

}
