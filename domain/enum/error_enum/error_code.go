package enum

/**
* First Code indicate the scenario
* GNE = General Error
* LGE = Login Error
* RGE = Register Error
*
* Second Code indicate sub scenario
* TS = Test Scheduler
* HK = House Keeping
*
* Third Code indicate either Request / Response
* RQ = Request
* RS = Response
* SY = System
*
* Fourth Code Indicate Level of Error
* FE = Fatal Error
* NE = Normal Error
*
* Fifth Code Indicate the specific case
* 001 - XXX
 */

const (
	// General Error
	CodeInvalidJson         = "GNE-GN-RQ-FE-001"
	CodeInternalServerError = "GNE-GN-RQ-FE-002"
	CodeSystemError         = "GNE-GN-RQ-FE-003"
	CodeUnauthorized        = "GNE-GN-RQ-NE-003"

	CodeDatabaseError = "GNE-DB-SY-FE-001"
	CodeBadRequest    = "GNE-GN-RQ-FE-001"

	// Token Related
	CodeRefreshTokenExpired = "TKE-RF-RQ-NE-001"
	CodeInvalidAccessToken  = "TKE-AU-RQ-FE-002"
)
