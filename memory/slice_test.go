package memory_test

import (
	"testing"

	"github.com/nolag/gocpu/memory"
)

func TestRunReadWriteTest(t *testing.T) {
	RunReadWriteTest(t, true, memory.NewSlice)
}

func TestNewMeorySliceCreatesMemoryWithCorrectSize(t *testing.T) {
	RunSizeTest(t, memory.NewSlice)
}

func TestSliceBoundsChceking(t *testing.T) {
	RunBoundsTests(t, memory.NewSlice)
}
