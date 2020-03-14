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

type Deployment struct {
	ApiVersion string              `json:"apiVersion,omitempty"`
	Kind       string              `json:"kind,omitempty"`
	Metadata   *DeploymentMetadata `json:"metadata,omitempty"`
	Spec       *DeploymentSpec     `json:"spec,omitempty"`
	Status     *DeploymentStatus   `json:"status,omitempty"`
}