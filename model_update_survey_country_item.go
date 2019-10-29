/*
 * TextMagic API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package TextMagic

type UpdateSurveyCountryItem struct {
	// The 2-letter ISO country code
	Country string `json:"country"`
	// User inbound phone ID
	UserInboundId int32 `json:"userInboundId"`
}
