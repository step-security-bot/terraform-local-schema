package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-local-schema/v2/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestLocalSensitiveFileSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.LocalSensitiveFileSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
