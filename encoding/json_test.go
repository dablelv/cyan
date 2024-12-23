package encoding

import (
	"testing"

	"github.com/dablelv/cyan/internal/utest"
)

func TestToIndentJSON(t *testing.T) {
	assert := utest.NewAssert(t, "ToIndentJSON")

	got, err := ToIndentJSON("foo")
	assert.IsNil(err)
	assert.Equal(`"foo"`, got)

	_, err = ToIndentJSON(complex(3, -5))
	assert.IsNotNil(err)
}
