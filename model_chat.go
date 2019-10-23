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

type Chat struct {
	// Chat ID.
	Id int32 `json:"id"`
	OriginalId int32 `json:"originalId"`
	// Chat partner phone number.
	Phone string `json:"phone"`
	Contact *Contact `json:"contact"`
	// If this field has a value then it means that chat phone number has been unsubscribed from you and this value is a ID of a Unsubscribed contact entity. See [Get all unsubscribed contacts](http://docs.textmagictesting.com/#operation/getUnsubscribers).
	UnsubscribedContactId int32 `json:"unsubscribedContactId"`
	// Total unread incoming messages.
	Unread int32 `json:"unread"`
	// Time when last incoming message arrived at this chat.
	UpdatedAt time.Time `json:"updatedAt"`
	// Chat status:   * **a** - Active   * **c** - Closed   * **d** - Deleted 
	Status string `json:"status"`
	// Indicates when chat is muted.
	Mute int32 `json:"mute"`
	// The last message content of a chat.
	LastMessage string `json:"lastMessage"`
	// Last message type: * **ci** - incoming call * **co** - outgoing call * **i** - incoming message * **o** - outgoing message 
	Direction string `json:"direction"`
	// If filled then value will be used as a sender number for all outgoing messages of a chat.
	From string `json:"from"`
	// Date and time until chat will be mutted.
	MutedUntil time.Time `json:"mutedUntil"`
	// Time left till chat will be unmutted (seconds).
	TimeLeftMute int32 `json:"timeLeftMute"`
	Country *Country `json:"country"`
}