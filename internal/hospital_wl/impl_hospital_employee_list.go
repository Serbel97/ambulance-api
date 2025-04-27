package hospital_wl

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"slices"
)

type implHospitalEmployeeListAPI struct {
}

func NewHospitalEmployeeListApi() HospitalEmployeeListAPI {
	return &implHospitalEmployeeListAPI{}
}

func (o *implHospitalEmployeeListAPI) CreateEmployeeListEntry(c *gin.Context) {
	//c.AbortWithStatus(http.StatusNotImplemented)
	updateHospitalFunc(c, func(c *gin.Context, hospital *Hospital) (*Hospital, interface{}, int) {
		var entry EmployeeListEntry

		if err := c.ShouldBindJSON(&entry); err != nil {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Invalid request body",
				"error":   err.Error(),
			}, http.StatusBadRequest
		}

// 		if entry.Name == "" {
// 			return nil, gin.H{
// 				"status":  http.StatusBadRequest,
// 				"message": "Name is required",
// 			}, http.StatusBadRequest
// 		}

		if entry.Id == "" || entry.Id == "@new" {
			entry.Id = uuid.NewString()
		}

		conflictIndx := slices.IndexFunc(hospital.EmployeeList, func(employee EmployeeListEntry) bool {
			return entry.Id == employee.Id //|| entry.PatientId == employee.PatientId
		})

		if conflictIndx >= 0 {
			return nil, gin.H{
				"status":  http.StatusConflict,
				"message": "Entry already exists",
			}, http.StatusConflict
		}

		hospital.EmployeeList = append(hospital.EmployeeList, entry)
// 		hospital.reconcileEmployeeList()
		// entry was copied by value return reconciled value from the list
		entryIndx := slices.IndexFunc(hospital.EmployeeList, func(employee EmployeeListEntry) bool {
			return entry.Id == employee.Id
		})
		if entryIndx < 0 {
			return nil, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Failed to save entry",
			}, http.StatusInternalServerError
		}
		return hospital, hospital.EmployeeList[entryIndx], http.StatusOK
	})
}

func (o *implHospitalEmployeeListAPI) DeleteEmployeeListEntry(c *gin.Context) {
	//c.AbortWithStatus(http.StatusNotImplemented)
	updateHospitalFunc(c, func(c *gin.Context, hospital *Hospital) (*Hospital, interface{}, int) {
		entryId := c.Param("entryId")

		if entryId == "" {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Entry ID is required",
			}, http.StatusBadRequest
		}

		entryIndx := slices.IndexFunc(hospital.EmployeeList, func(employee EmployeeListEntry) bool {
			return entryId == employee.Id
		})

		if entryIndx < 0 {
			return nil, gin.H{
				"status":  http.StatusNotFound,
				"message": "Entry not found",
			}, http.StatusNotFound
		}

		hospital.EmployeeList = append(hospital.EmployeeList[:entryIndx], hospital.EmployeeList[entryIndx+1:]...)
// 		hospital.reconcileEmployeeList()
		return hospital, nil, http.StatusNoContent
	})
}

func (o *implHospitalEmployeeListAPI) GetEmployeeListEntries(c *gin.Context) {
	//c.AbortWithStatus(http.StatusNotImplemented)
	updateHospitalFunc(c, func(c *gin.Context, hospital *Hospital) (*Hospital, interface{}, int) {
		result := hospital.EmployeeList
		if result == nil {
			result = []EmployeeListEntry{}
		}
		// return nil hospital - no need to update it in db
		return nil, result, http.StatusOK
	})
}

func (o *implHospitalEmployeeListAPI) GetEmployeeListEntry(c *gin.Context) {
	//c.AbortWithStatus(http.StatusNotImplemented)
	updateHospitalFunc(c, func(c *gin.Context, hospital *Hospital) (*Hospital, interface{}, int) {
		entryId := c.Param("entryId")

		if entryId == "" {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Entry ID is required",
			}, http.StatusBadRequest
		}

		entryIndx := slices.IndexFunc(hospital.EmployeeList, func(employee EmployeeListEntry) bool {
			return entryId == employee.Id
		})

		if entryIndx < 0 {
			return nil, gin.H{
				"status":  http.StatusNotFound,
				"message": "Entry not found",
			}, http.StatusNotFound
		}

		// return nil hospital - no need to update it in db
		return nil, hospital.EmployeeList[entryIndx], http.StatusOK
	})
}

func (o *implHospitalEmployeeListAPI) UpdateEmployeeListEntry(c *gin.Context) {
	//c.AbortWithStatus(http.StatusNotImplemented)
	updateHospitalFunc(c, func(c *gin.Context, hospital *Hospital) (*Hospital, interface{}, int) {
		var entry EmployeeListEntry

		if err := c.ShouldBindJSON(&entry); err != nil {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Invalid request body",
				"error":   err.Error(),
			}, http.StatusBadRequest
		}

		entryId := c.Param("entryId")

		if entryId == "" {
		    return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Entry ID is required",
			}, http.StatusBadRequest
		}

		entryIndx := slices.IndexFunc(hospital.EmployeeList, func(employee EmployeeListEntry) bool {
			return entryId == employee.Id
		})

		if entryIndx < 0 {
			return nil, gin.H{
				"status":  http.StatusNotFound,
				"message": "Entry not found",
			}, http.StatusNotFound
		}

// 		if entry.PatientId != "" {
// 			hospital.EmployeeList[entryIndx].PatientId = entry.PatientId
// 		}

		if entry.Id != "" {
			hospital.EmployeeList[entryIndx].Id = entry.Id
		}

// 		if entry.EmployeeSince.After(time.Time{}) {
// 			hospital.EmployeeList[entryIndx].EmployeeSince = entry.EmployeeSince
// 		}

// 		if entry.EstimatedDurationMinutes > 0 {
// 			hospital.EmployeeList[entryIndx].EstimatedDurationMinutes = entry.EstimatedDurationMinutes
// 		}

// 		hospital.reconcileEmployeeList()
		return hospital, hospital.EmployeeList[entryIndx], http.StatusOK
	})
}
