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
	updateHospitalFunc(c, func(c *gin.Context, hospital *Hospital) (*Hospital, interface{}, int) {
		var entry EmployeeListEntry

		if err := c.ShouldBindJSON(&entry); err != nil {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Invalid request body",
				"error":   err.Error(),
			}, http.StatusBadRequest
		}

		if entry.Id == "" || entry.Id == "@new" {
			entry.Id = uuid.NewString()
		}

		conflictIndx := slices.IndexFunc(hospital.EmployeeList, func(employee EmployeeListEntry) bool {
			return entry.Id == employee.Id
		})

		if conflictIndx >= 0 {
			return nil, gin.H{
				"status":  http.StatusConflict,
				"message": "Entry already exists",
			}, http.StatusConflict
		}

		hospital.EmployeeList = append(hospital.EmployeeList, entry)
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
		return hospital, nil, http.StatusNoContent
	})
}

func (o *implHospitalEmployeeListAPI) GetEmployeeListEntries(c *gin.Context) {
	updateHospitalFunc(c, func(c *gin.Context, hospital *Hospital) (*Hospital, interface{}, int) {
		result := hospital.EmployeeList
		if result == nil {
			result = []EmployeeListEntry{}
		}
		return nil, result, http.StatusOK
	})
}

func (o *implHospitalEmployeeListAPI) GetEmployeeListEntry(c *gin.Context) {
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

		return nil, hospital.EmployeeList[entryIndx], http.StatusOK
	})
}

func (o *implHospitalEmployeeListAPI) UpdateEmployeeListEntry(c *gin.Context) {
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

		return hospital, hospital.EmployeeList[entryIndx], http.StatusOK
	})
}

func (o *implHospitalEmployeeListAPI) TransferEmployeeListEntry(c *gin.Context) {
	srcHospID := c.Param("hospitalId")
	entryID := c.Param("entryId")
	if srcHospID == "" || entryID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "hospitalId and entryId are required"})
		return
	}

	var req struct {
		TargetHospitalId string `json:"targetHospitalId"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.TargetHospitalId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Invalid or missing targetHospitalId", "error": err.Error()})
		return
	}

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

	srcHosp, err := dbSvc.FindDocument(c, srcHospID)
	if err != nil {
		if err == db_service.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "source hospital not found"})
		} else {
			c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway, "message": err.Error()})
		}
		return
	}

	idx := slices.IndexFunc(srcHosp.EmployeeList, func(e EmployeeListEntry) bool {
		return e.Id == entryID
	})
	if idx < 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "entry not found in source hospital"})
		return
	}
	entry := srcHosp.EmployeeList[idx]

	srcHosp.EmployeeList = append(srcHosp.EmployeeList[:idx], srcHosp.EmployeeList[idx+1:]...)
	if err := dbSvc.UpdateDocument(c, srcHospID, srcHosp); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway, "message": "failed to update source hospital", "error": err.Error()})
		return
	}

	dstHosp, err := dbSvc.FindDocument(c, req.TargetHospitalId)
	if err != nil {
		if err == db_service.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "target hospital not found"})
		} else {
			c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway, "message": err.Error()})
		}
		return
	}

	dstHosp.EmployeeList = append(dstHosp.EmployeeList, entry)
	if err := dbSvc.UpdateDocument(c, req.TargetHospitalId, dstHosp); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway, "message": "failed to update target hospital", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, entry)
}
