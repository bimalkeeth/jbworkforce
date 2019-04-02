package entityrelations

import (
	"github.com/jinzhu/gorm"
	ent "jbworkforce/masterdataservice/master_dataaccess/entities"
)

func CreateEmployee(database *gorm.DB) {

	if !database.HasTable(&ent.TableEmployee{}) {
		database.CreateTable(&ent.TableEmployee{})
		database.Model(&ent.TableEmployee{}).AddUniqueIndex("idx_table_employee_empno", "empno")
		database.Model(&ent.TableEmployee{}).Association("EmployeeDetail")
		database.Model(&ent.TableEmployee{}).Association("EmployeeContacts")
		database.Model(&ent.TableEmployee{}).Association("EmployeeAddress")
		database.Model(&ent.TableEmployee{}).Association("EmployeeAgencies")
		database.Model(&ent.TableEmployee{}).Association("EmployeeClinicalStream")
		database.Model(&ent.TableEmployee{}).Association("EmployeeAllowance")
		database.Model(&ent.TableEmployee{}).Association("EmployeeCompetency")
		database.Model(&ent.TableEmployee{}).Association("EmployeeTransitions")
		database.Model(&ent.TableEmployee{}).Association("EmployeeDepartmentPreference")
		database.Model(&ent.TableEmployee{}).Association("EmployeeDiscrepancies")
		database.Model(&ent.TableEmployee{}).Association("EmployeeEmpType")
		database.Model(&ent.TableEmployee{}).Association("EmployeeGroupEmp")
		database.Model(&ent.TableEmployee{}).Association("Gender")
	}
}
