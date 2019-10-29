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

type BulkSession struct {
	// Bulk Session ID.
	Id int32 `json:"id"`
	// * **n** – bulk session is just created * **w** - work in progress * **f** - failed * **c** - completed with success * **s** - suspended 
	Status string `json:"status"`
	// Amount of messages already processed.
	ItemsProcessed int32 `json:"itemsProcessed"`
	// Total amount of messages to be processed.
	ItemsTotal int32 `json:"itemsTotal"`
	// Creation date and time of a Bulk Session.
	CreatedAt time.Time `json:"createdAt"`
	Session *MessageSession `json:"session"`
	// Message text of a Bulk Session.
	Text string `json:"text"`
}
