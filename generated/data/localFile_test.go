package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-local-schema/v2/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestLocalFileSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.LocalFileSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
