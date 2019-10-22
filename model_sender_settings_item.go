/*
 * TextMagic API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package TextMagic

type SenderSettingsItem struct {
	// Two-letter ISO country code of the recipient phone number. 
	Country string `json:"country"`
	// Phone enabled for sending to specified country
	Phone string `json:"phone"`
}
