/*
 * Jumpserver API Docs
 *
 * Jumpserver Restful api docs
 *
 * API version: v1
 * Contact: support@fit2cloud.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

type Ticket struct {
	Id string `json:"id,omitempty"`
	User string `json:"user,omitempty"`
	UserDisplay string `json:"user_display,omitempty"`
	Title string `json:"title"`
	Body string `json:"body"`
	Assignees []string `json:"assignees"`
	AssigneesDisplay string `json:"assignees_display,omitempty"`
	Status string `json:"status,omitempty"`
	Action string `json:"action,omitempty"`
	DateCreated time.Time `json:"date_created,omitempty"`
	DateUpdated time.Time `json:"date_updated,omitempty"`
	Type_ string `json:"type,omitempty"`
	TypeDisplay string `json:"type_display,omitempty"`
	ActionDisplay string `json:"action_display,omitempty"`
}