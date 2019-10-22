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

type UsersInbound struct {
	// Dedicated number ID.
	Id int32 `json:"id"`
	// Format for representation of time
	DisplayTimeFormat string `json:"displayTimeFormat,omitempty"`
	// Dedicated phone number.
	Phone string `json:"phone,omitempty"`
	User *User `json:"user"`
	// Time when the dedicated number was purchased.
	PurchasedAt time.Time `json:"purchasedAt"`
	// Dedicated number subscription expiration time.
	ExpireAt time.Time `json:"expireAt"`
	// Number status: *   **U** for Unused. No messages have been sent from (or received to) this number. *   **A** for Active. 
	Status string `json:"status"`
	Country *Country `json:"country"`
}
