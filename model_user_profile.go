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

type UserProfile struct {
	Id string `json:"id,omitempty"`
	Username string `json:"username"`
	Name string `json:"name"`
	Role string `json:"role,omitempty"`
	Email string `json:"email"`
}
