/*
 * TextMagic API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package TextMagic

import (
	"time"
)

type MessageTemplate struct {
	// Template ID.
	Id int32 `json:"id"`
	// Template name.
	Name string `json:"name"`
	// Template text. May contain the tags. See [Custom fields list](http://docs.textmagictesting.com/#section/Custom-fields-list-(Merge-tags)).
	Content string `json:"content"`
	// Time when template was last modified.
	LastModified time.Time `json:"lastModified"`
}