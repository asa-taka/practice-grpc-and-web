/*
 * todo.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

type TodoRecord struct {

	Id int32 `json:"id,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	Title string `json:"title,omitempty"`

	Detail string `json:"detail,omitempty"`

	Deadline time.Time `json:"deadline,omitempty"`

	Status *RecordStatus `json:"status,omitempty"`
}