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

type CommandExecution struct {
	Id string `json:"id,omitempty"`
	Hosts []string `json:"hosts"`
	RunAs string `json:"run_as"`
	Command string `json:"command"`
	Result string `json:"result,omitempty"`
	LogUrl string `json:"log_url,omitempty"`
	IsFinished bool `json:"is_finished,omitempty"`
	DateCreated time.Time `json:"date_created,omitempty"`
	DateFinished time.Time `json:"date_finished,omitempty"`
}