package worker

import (
	"testing"

	"github.com/go-load-test/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_GetTimingFrom(t *testing.T) {
	handle := mocks.JSHandle{
		Children: map[string]mocks.JSHandle{
			"foo": mocks.JSHandle{Value: "150", Children: nil},
		},
	}

	timing, err := getTimingFrom(handle, "foo")

	assert.Equal(t, nil, err, "should not return an error")
	assert.Equal(t, 150, timing, "should return a timing")
}
