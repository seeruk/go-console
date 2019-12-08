package specification_test

import (
	"testing"

	"github.com/seeruk/go-console/specification"
	"github.com/stretchr/testify/assert"
)

func TestParseArgumentSpecification(t *testing.T) {
	t.Run("should set the name", func(t *testing.T) {
		argument, err := specification.ParseArgumentSpecification("GALAXY_QUEST")
		assert.NoError(t, err)
		assert.Equal(t, "GALAXY_QUEST", argument.Name)
	})

	t.Run("should set whether or not the argument is required", func(t *testing.T) {
		argument, err := specification.ParseArgumentSpecification("GALAXY_QUEST")
		assert.NoError(t, err)
		assert.Equal(t, true, argument.Required)

		argument, err = specification.ParseArgumentSpecification("[MEMENTO]")
		assert.NoError(t, err)
		assert.Equal(t, false, argument.Required)
	})

	t.Run("should expect a close bracket if an opening one is given", func(t *testing.T) {
		_, err := specification.ParseArgumentSpecification("[MEMENTO")
		assert.Error(t, err)
	})

	t.Run("should not expect any whitespace", func(t *testing.T) {
		_, err := specification.ParseArgumentSpecification("GALAXY QUEST")
		assert.Error(t, err)
	})

	t.Run("should always expect an identifier", func(t *testing.T) {
		_, err := specification.ParseArgumentSpecification("")
		assert.Error(t, err)

		_, err = specification.ParseArgumentSpecification("[]")
		assert.Error(t, err)
	})
}
