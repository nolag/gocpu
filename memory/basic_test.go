package memory

import (
	"testing"
)

func TestRunReadWriteTest(t *testing.T) {
	RunReadWriteTest(t, true, NewBasic)
}

func TestNewMeorySliceCreatesMemoryWithCorrectSize(t *testing.T) {
	RunSizeTest(t, NewBasic)
}

func TestBoundsChceking(t *testing.T) {
	RunBoundsTests(t, NewBasic)
}