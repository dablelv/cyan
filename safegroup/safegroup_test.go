package safegroup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroup_Go(t *testing.T) {
	var sg Group
	sg.Go(func() error {
		return nil
	})
	sg.Go(func() error {
		panic("this is a panic")
	})
	err := sg.Wait()
	assert.NotNil(t, err)
}
