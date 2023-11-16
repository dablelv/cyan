package str

import (
	"testing"

	"github.com/dablelv/cyan/internal"
)

func TestCapitalize(t *testing.T) {
	assert := internal.NewAssert(t, "TestCapitalize")

	// Test 1: Checks empty string.
	assert.Equal("", Capitalize(""))

	// Test 2: Checks if function works with 1 parameter, and default whitespace delimiter.
	in := "test is going.well.thank.you.for inquiring"
	assert.Equal("Test Is Going.well.thank.you.for Inquiring", Capitalize(in))

	// Test 3: Checks if function works with both parameters, with param 2 containing whitespace and '.'
	delimiters := []rune{' ', '.'}
	assert.Equal("Test Is Going.Well.Thank.You.For Inquiring", Capitalize(in, delimiters...))
}

func TestCapitalizeFully(t *testing.T) {
	assert := internal.NewAssert(t, "CapitalizeFully")

	assert.Equal("", CapitalizeFully(""))

	in := "tEsT iS goiNG.wELL.tHaNk.yOU.for inqUIrING"

	assert.Equal("Test Is Going.well.thank.you.for Inquiring", CapitalizeFully(in))

	delimiters := []rune{' ', '.'}
	assert.Equal("Test Is Going.Well.Thank.You.For Inquiring", CapitalizeFully(in, delimiters...))
}

func TestUncapitalize(t *testing.T) {
	assert := internal.NewAssert(t, "Uncapitalize")

	assert.Equal("", Uncapitalize(""))

	in := "This Is A.Test"

	assert.Equal("this is a.Test", Uncapitalize(in))

	delimiters := []rune{' ', '.'}
	assert.Equal("this is a.test", Uncapitalize(in, delimiters...))
}

func TestInitials(t *testing.T) {
	assert := internal.NewAssert(t, "Initials")

	assert.Equal("", Initials(""))

	in := "John Doe.Ray"
	assert.Equal("JD", Initials(in))

	delimiters := []rune{' ', '.'}
	assert.Equal("JDR", Initials(in, delimiters...))
}
