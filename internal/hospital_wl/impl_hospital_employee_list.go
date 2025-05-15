package hospital_wl

import (
	"net/http"

	"github.com/Serbel97/ambulance-api/internal/db_service"
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
		if entry.Name != "" {
			hospital.EmployeeList[entryIndx].Name = entry.Name
		}
		if entry.Role.Value != "" {
			hospital.EmployeeList[entryIndx].Role.Value = entry.Role.Value
		}
		if entry.Role.Code != "" {
			hospital.EmployeeList[entryIndx].Role.Code = entry.Role.Code
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

// TransferEmployeeListEntry moves an employee from one hospital to another
// POST /api/employee-list/:hospitalId/entries/:entryId/transfer
func (o *implHospitalEmployeeListAPI) TransferEmployeeListEntry(c *gin.Context) {
	// 1) parse path params
	srcHospID := c.Param("hospitalId")
	entryID := c.Param("entryId")
	if srcHospID == "" || entryID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "hospitalId and entryId are required"})
		return
	}

	// 2) bind JSON { "targetHospitalId": "..." }
	var req struct {
		TargetHospitalId string `json:"targetHospitalId"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.TargetHospitalId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Invalid or missing targetHospitalId", "error": err.Error()})
		return
	}

	// 3) grab the db service from context
	val, ok := c.Get("db_service")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "db_service not found"})
		return
	}
	dbSvc, ok := val.(db_service.DbService[Hospital])
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "invalid db_service type"})
		return
	}

	// 4) load source hospital
	srcHosp, err := dbSvc.FindDocument(c, srcHospID)
	if err != nil {
		if err == db_service.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "source hospital not found"})
		} else {
			c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway, "message": err.Error()})
		}
		return
	}

	// 5) locate the entry in source
	idx := slices.IndexFunc(srcHosp.EmployeeList, func(e EmployeeListEntry) bool {
		return e.Id == entryID
	})
	if idx < 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "entry not found in source hospital"})
		return
	}
	entry := srcHosp.EmployeeList[idx]

	// 6) remove from source list
	srcHosp.EmployeeList = append(srcHosp.EmployeeList[:idx], srcHosp.EmployeeList[idx+1:]...)
	if err := dbSvc.UpdateDocument(c, srcHospID, srcHosp); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway, "message": "failed to update source hospital", "error": err.Error()})
		return
	}

	// 7) load destination hospital
	dstHosp, err := dbSvc.FindDocument(c, req.TargetHospitalId)
	if err != nil {
		if err == db_service.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "target hospital not found"})
		} else {
			c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway, "message": err.Error()})
		}
		return
	}

	// 8) append to destination list
	dstHosp.EmployeeList = append(dstHosp.EmployeeList, entry)
	if err := dbSvc.UpdateDocument(c, req.TargetHospitalId, dstHosp); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway, "message": "failed to update target hospital", "error": err.Error()})
		return
	}

	// 9) respond with the transferred entry
	c.JSON(http.StatusOK, entry)
}
