package requirements

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheck(t *testing.T) {
	noErrorFunc := func() error {return nil}
	errorFunc := func() error {return errors.New("ERROR!")}


	t.Run("Returns no error for no functions", func(t *testing.T) {
		err := Check()

		assert.NoError(t, err)
	})

	t.Run("Returns no error for errorless functions", func(t *testing.T) {
		noErrorCallable1 := NewCallable(noErrorFunc)
		noErrorCallable2 := NewCallable(noErrorFunc)
		noErrorCallable3 := NewCallable(noErrorFunc)

		err := Check(noErrorCallable1.Run, noErrorCallable2.Run, noErrorCallable3.Run)

		assert.NoError(t, err)
		assert.True(t, noErrorCallable1.WasCalled)
		assert.True(t, noErrorCallable2.WasCalled)
		assert.True(t, noErrorCallable3.WasCalled)
	})

	t.Run("Calls functions unless an error is returned", func(t *testing.T) {
		noErrorCallable := NewCallable(noErrorFunc)
		errorCallable := NewCallable(errorFunc)

		err := Check(noErrorCallable.Run, errorCallable.Run)

		assert.Error(t, err)
		assert.True(t, noErrorCallable.WasCalled)
		assert.True(t, errorCallable.WasCalled)
	})

	t.Run("Shortcuts after error", func(t *testing.T) {
		noErrorCallable := NewCallable(noErrorFunc)
		errorCallable := NewCallable(errorFunc)

		err := Check(errorCallable.Run, noErrorCallable.Run)

		assert.Error(t, err)
		assert.True(t, errorCallable.WasCalled)
		assert.False(t, noErrorCallable.WasCalled)
	})
}

type Callable struct {
	WasCalled bool
	fn func() error
}

func NewCallable(fn func() error) Callable{
	return Callable{
		fn:       fn,
	}
}

func (c* Callable) Run() error {
	c.WasCalled = true
	return c.fn()
}