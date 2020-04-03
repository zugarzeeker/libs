package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBodyString_ShouldReturnString_WhenBodyExist(t *testing.T) {
	r := &Request{
		body: []byte("xxx"),
	}

	assert.Equal(t, "xxx", r.bodyString())
}

func TestMarshalBody_ShouldMarshalBody_ToBody(t *testing.T) {
	r := &Request{
		Body: "{}",
	}

	err := r.marshalBody()

	assert.NoError(t, err)
	assert.Equal(t, `{}`, string(r.body))
}
