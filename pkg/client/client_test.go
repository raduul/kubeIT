package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/client-go/tools/clientcmd"
)

func TestGetClient(t *testing.T) {
	_, err := clientcmd.BuildConfigFromFlags("", "/invalid/path")
	assert.Error(t, err)
}
