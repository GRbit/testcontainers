package testcontainers

import (
	"github.com/imdario/mergo"
	"github.com/romnnn/testcontainers-go"
)

// MergeRequest ...
func MergeRequest(c *testcontainers.ContainerRequest, override *testcontainers.ContainerRequest) {
	if err := mergo.Merge(c, override, mergo.WithOverride); err != nil {
		panic(err)
	}
}
