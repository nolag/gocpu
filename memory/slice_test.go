package memory

import "testing"

func TestRunReadWriteTest(t *testing.T) {
	RunReadWriteTest(t, true, NewSlice)
}

func TestNewMeorySliceCreatesMemoryWithCorrectSize(t *testing.T) {
	RunSizeTest(t, NewSlice)
}

func TestSliceBoundsChceking(t *testing.T) {
	RunBoundsTests(t, NewSlice)
}
