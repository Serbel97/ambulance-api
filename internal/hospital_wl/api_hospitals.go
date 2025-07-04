/*
 * Employee List Api
 *
 * Hospital Employee Administration for Web-In-Cloud system
 *
 * API version: 1.0.0
 * Contact: xbelake@stuba.sk
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package hospital_wl

import (
	"github.com/gin-gonic/gin"
)

type HospitalsAPI interface {

	// CreateHospital Post /api/hospital
	// Saves new hospital definition
	CreateHospital(c *gin.Context)

	// DeleteHospital Delete /api/hospital/:hospitalId
	// Deletes specific hospital
	DeleteHospital(c *gin.Context)

	// GetHospital Get /api/hospital
	// Provides the hospital list
	GetHospital(c *gin.Context)
}
