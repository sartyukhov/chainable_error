package chainableerror_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	cherrors "github.com/sartyukhov/chainable_error"
)

func TestChainableError(t *testing.T) {
	errChild := errors.New("child")
	errBase := cherrors.New("base")
	errTopLevel := fmt.Errorf("top level error: %w", errBase.Wrap(errChild))

	assert.Equal(t, "top level error: base: child", errTopLevel.Error())

	assert.ErrorIs(t, errTopLevel, errChild)
	assert.ErrorIs(t, errTopLevel, errBase)

	receiverErr := &cherrors.ChainableError{}
	assert.ErrorAs(t, errTopLevel, &receiverErr)
}

func TestChainableErrorMultipleTimes(t *testing.T) {
	errChildZero := errors.New("child0")
	errChildOne := errors.New("child1")
	errBase := cherrors.New("base")
	errTopLevelOne := fmt.Errorf("top level error: %w", errBase.Wrap(errChildZero))
	// wrap other error to the same base error
	errTopLevelTwo := fmt.Errorf("top level error: %w", errBase.Wrap(errChildOne))

	assert.Equal(t, "top level error: base: child0", errTopLevelOne.Error())

	assert.ErrorIs(t, errTopLevelOne, errChildZero)
	assert.ErrorIs(t, errTopLevelOne, errBase)

	assert.Equal(t, "top level error: base: child1", errTopLevelTwo.Error())

	assert.ErrorIs(t, errTopLevelTwo, errChildOne)
	assert.ErrorIs(t, errTopLevelTwo, errBase)
}

func TestChainableErrorEmpty(t *testing.T) {
	errBase := cherrors.New("base")
	errTopLevel := fmt.Errorf("top level error: %w", errBase)

	assert.Equal(t, "top level error: base", errTopLevel.Error())

	assert.ErrorIs(t, errTopLevel, errBase)

	receiverErr := &cherrors.ChainableError{}
	assert.ErrorAs(t, errTopLevel, &receiverErr)
}
