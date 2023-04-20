package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const localFile = `{
  "block": {
    "attributes": {
      "content": {
        "computed": true,
        "description": "Raw content of the file that was read, as UTF-8 encoded string. Files that do not contain UTF-8 text will have invalid UTF-8 sequences in ` + "`" + `content` + "`" + `\n  replaced with the Unicode replacement character. ",
        "description_kind": "plain",
        "type": "string"
      },
      "content_base64": {
        "computed": true,
        "description": "Base64 encoded version of the file content (use this when dealing with binary data).",
        "description_kind": "plain",
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
      "filename": {
        "description": "Path to the file that will be read. The data source will return an error if the file does not exist.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "The hexadecimal encoding of the SHA1 checksum of the file content.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Reads a file from the local filesystem.",
    "description_kind": "plain"
  },
  "version": 0
}`

func LocalFileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(localFile), &result)
	return &result
}
