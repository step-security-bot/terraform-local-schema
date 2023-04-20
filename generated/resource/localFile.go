package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const localFile = `{
  "block": {
    "attributes": {
      "content": {
        "description": "Content to store in the file, expected to be a UTF-8 encoded string.\n Conflicts with ` + "`" + `sensitive_content` + "`" + `, ` + "`" + `content_base64` + "`" + ` and ` + "`" + `source` + "`" + `.\n Exactly one of these four arguments must be specified.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "content_base64": {
        "description": "Content to store in the file, expected to be binary encoded as base64 string.\n Conflicts with ` + "`" + `content` + "`" + `, ` + "`" + `sensitive_content` + "`" + ` and ` + "`" + `source` + "`" + `.\n Exactly one of these four arguments must be specified.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "content_base64sha256": {
        "computed": true,
        "description": "Base64 encoded SHA256 checksum of file content.",
        "description_kind": "plain",
        "type": "string"
      },
      "content_base64sha512": {
        "computed": true,
        "description": "Base64 encoded SHA512 checksum of file content.",
        "description_kind": "plain",
        "type": "string"
      },
      "content_md5": {
        "computed": true,
        "description": "MD5 checksum of file content.",
        "description_kind": "plain",
        "type": "string"
      },
      "content_sha1": {
        "computed": true,
        "description": "SHA1 checksum of file content.",
        "description_kind": "plain",
        "type": "string"
      },
      "content_sha256": {
        "computed": true,
        "description": "SHA256 checksum of file content.",
        "description_kind": "plain",
        "type": "string"
      },
      "content_sha512": {
        "computed": true,
        "description": "SHA512 checksum of file content.",
        "description_kind": "plain",
        "type": "string"
      },
      "directory_permission": {
        "computed": true,
        "description": "Permissions to set for directories created (before umask), expressed as string in\n [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).\n Default value is ` + "`" + `\"0777\"` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "file_permission": {
        "computed": true,
        "description": "Permissions to set for the output file (before umask), expressed as string in\n [numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).\n Default value is ` + "`" + `\"0777\"` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "filename": {
        "description": "The path to the file that will be created.\n Missing parent directories will be created.\n If the file already exists, it will be overridden with the given content.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "The hexadecimal encoding of the SHA1 checksum of the file content.",
        "description_kind": "plain",
        "type": "string"
      },
      "sensitive_content": {
        "deprecated": true,
        "description": "Sensitive content to store in the file, expected to be an UTF-8 encoded string.\n Will not be displayed in diffs.\n Conflicts with ` + "`" + `content` + "`" + `, ` + "`" + `content_base64` + "`" + ` and ` + "`" + `source` + "`" + `.\n Exactly one of these four arguments must be specified.\n If in need to use _sensitive_ content, please use the [` + "`" + `local_sensitive_file` + "`" + `](./sensitive_file.html)\n resource instead.",
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "source": {
        "description": "Path to file to use as source for the one we are creating.\n Conflicts with ` + "`" + `content` + "`" + `, ` + "`" + `sensitive_content` + "`" + ` and ` + "`" + `content_base64` + "`" + `.\n Exactly one of these four arguments must be specified.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Generates a local file with the given content.",
    "description_kind": "plain"
  },
  "version": 0
}`

func LocalFileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(localFile), &result)
	return &result
}
