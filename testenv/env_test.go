package testenv_test

import (
	"testing"

	"os"

	"github.com/goph/xtesting/testenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadEnvFromFile(t *testing.T) {
	err := testenv.LoadEnvFromFile("testdata/.env")
	require.NoError(t, err)
	assert.Equal(t, "test_value", os.Getenv("TEST_VARIABLE"))
}
