/*
 * TextMagic API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package TextMagic

type UnsubscribeContactInputObject struct {
	// Contact phone number.
	Phone string `json:"phone,omitempty"`
	// If set to 1 incoming messages from this number will be blocked.
	BlockIncoming int32 `json:"blockIncoming,omitempty"`
}
