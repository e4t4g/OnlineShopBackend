/*
 * Backend for Online Shop
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	ID     uuid.UUID
	UserID uuid.UUID `json:"userID,omitempty"`
	// Date     string `json:"date,omitempty"`
	Items    []Item `json:"products,omitempty"`
	ExpireAt time.Time
}
