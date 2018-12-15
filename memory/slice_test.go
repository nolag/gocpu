package memory_test

import (
	"testing"

	"github.com/nolag/gocpu/memory"
	"github.com/nolag/gocpu/memory/testhelper"
)

func TestRunReadWriteTest(t *testing.T) {
	testhelper.RunReadWriteTest(t, true, memory.NewSlice)
}

func TestNewMeorySliceCreatesMemoryWithCorrectSize(t *testing.T) {
	testhelper.RunSizeTest(t, memory.NewSlice)
}

func TestSliceBoundsChceking(t *testing.T) {
	testhelper.RunBoundsTests(t, memory.NewSlice)
}
