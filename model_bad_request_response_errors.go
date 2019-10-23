/*
 * TextMagic API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package TextMagic

// If it was a **POST** or **PUT** request (and the **message** returned is `Validation Failed`), this field may contain **errors **that describe the errors grouped by the input parameter name. 
type BadRequestResponseErrors struct {
	// Array of messages with errors related to the entire request. For example, you did not specify either the **text** or **templateId** when [sending the message](http://docs.textmagictesting.com/#tag/Outbound-Messages). 
	Common []string `json:"common,omitempty"`
	// Associative array. The keys are the POST/PUT parameters names and the values are arrays with error messages for these parameters. 
	Fields *interface{} `json:"fields,omitempty"`
}