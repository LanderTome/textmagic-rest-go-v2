/*
 * TextMagic API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package TextMagic

type RequestNewSubaccountTokenInputObject struct {
	// Sub-account ID.
	UserId int32 `json:"userId"`
	// Your account password.
	Password string `json:"password"`
	// Application name.
	AppName string `json:"appName,omitempty"`
}
