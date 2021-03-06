/*
 * Application Manager API
 *
 * Application Manager APIs to control Apache Flink jobs
 *
 * API version: 2.1.0
 * Contact: platform@ververica.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package appmanagerapi

import (
	"time"
)

type DeploymentStatusRunning struct {
	Conditions     []DeploymentCondition `json:"conditions,omitempty"`
	JobId          string                `json:"jobId,omitempty"`
	TransitionTime *time.Time            `json:"transitionTime,omitempty"`
}
