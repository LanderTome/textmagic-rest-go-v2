/*
 * TextMagic API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package TextMagic

type CreateListInputObject struct {
	// List name.
	Name string `json:"name"`
	// Should the new list be **shared** among all the sub-accounts?
	Shared bool `json:"shared,omitempty"`
	// Is the list favorited? Default is false.
	Favorited bool `json:"favorited,omitempty"`
	// Is the list default for new contacts (web only)?
	IsDefault bool `json:"isDefault,omitempty"`
}
