package generated

import (
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-local-schema/v2/generated/data"
	"github.com/lonegunmanb/terraform-local-schema/v2/generated/resource"
)

var Resources map[string]*tfjson.Schema
var DataSources map[string]*tfjson.Schema

func init() {
	resources := make(map[string]*tfjson.Schema, 0)
	dataSources := make(map[string]*tfjson.Schema, 0)  
	resources["local_file"] = resource.LocalFileSchema()  
	resources["local_sensitive_file"] = resource.LocalSensitiveFileSchema()  
	dataSources["local_file"] = data.LocalFileSchema()  
	dataSources["local_sensitive_file"] = data.LocalSensitiveFileSchema()  
	Resources = resources
	DataSources = dataSources
}