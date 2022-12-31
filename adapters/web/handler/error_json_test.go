package handler

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHandler_jsonError(t *testing.T) {
	test := "Hello Json"
	result := jsonError(test)
	require.Equal(t, []byte(`{"message":"Hello Json"}`), result)
}
