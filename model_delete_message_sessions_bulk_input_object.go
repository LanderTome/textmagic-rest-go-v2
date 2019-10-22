/*
 * TextMagic API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package TextMagic

type DeleteMessageSessionsBulkInputObject struct {
	// Entity ID(s), separated by comma
	Ids string `json:"ids,omitempty"`
	// Default is 0 (false). If set to 1 all the entities will be removed.
	All int32 `json:"all,omitempty"`
}
