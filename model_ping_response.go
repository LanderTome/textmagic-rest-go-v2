/*
 * TextMagic API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package TextMagic

type PingResponse struct {
	// Pong.
	Ping string `json:"ping"`
	// Current date time
	UtcDateTime string `json:"utcDateTime"`
}