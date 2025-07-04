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

type HospitalEmployeeListAPI interface {

	// CreateEmployeeListEntry Post /api/employee-list/:hospitalId/entries
	// Saves new entry into employee list
	CreateEmployeeListEntry(c *gin.Context)

	// DeleteEmployeeListEntry Delete /api/employee-list/:hospitalId/entries/:entryId
	// Deletes specific entry
	DeleteEmployeeListEntry(c *gin.Context)

	// GetEmployeeListEntries Get /api/employee-list/:hospitalId/entries
	// Provides the hospital employee list
	GetEmployeeListEntries(c *gin.Context)

	// GetEmployeeListEntry Get /api/employee-list/:hospitalId/entries/:entryId
	// Provides details about employee list entry
	GetEmployeeListEntry(c *gin.Context)

	// TransferEmployeeListEntry Post /api/employee-list/:hospitalId/entries/:entryId/transfer
	// Transfer an employee entry to another hospital
	TransferEmployeeListEntry(c *gin.Context)

	// UpdateEmployeeListEntry Put /api/employee-list/:hospitalId/entries/:entryId
	// Updates specific entry
	UpdateEmployeeListEntry(c *gin.Context)

	// GetPerformanceEntries Get /api/employee-list/:hospitalId/entries/:entryId/performances
	// Get all performance entries for an employee
	GetPerformanceEntries(c *gin.Context)

	// CreatePerformanceEntry Post /api/employee-list/:hospitalId/entries/:entryId/performances
	// Create a new performance entry for an employee
	CreatePerformanceEntry(c *gin.Context)

	// GetPerformanceEntry Get /api/employee-list/:hospitalId/entries/:entryId/performances/:performanceId
	// Get a specific performance entry
	GetPerformanceEntry(c *gin.Context)

	// UpdatePerformanceEntry Put /api/employee-list/:hospitalId/entries/:entryId/performances/:performanceId
	// Update a performance entry
	UpdatePerformanceEntry(c *gin.Context)

	// DeletePerformanceEntry Delete /api/employee-list/:hospitalId/entries/:entryId/performances/:performanceId
	// Delete a performance entry
	DeletePerformanceEntry(c *gin.Context)
}
