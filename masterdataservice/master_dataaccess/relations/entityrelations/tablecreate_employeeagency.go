package entityrelations

import (
	"github.com/jinzhu/gorm"
	ent "jbworkforce/masterdataservice/master_dataaccess/entities"
)

//----------------------------------------------------
// Employee Address
//----------------------------------------------------
func CreateEmployeeAgency(database *gorm.DB) {
	if !database.HasTable(&ent.TableEmployeeAgency{}) {
		database.CreateTable(&ent.TableEmployeeAgency{})
		database.Model(&ent.TableEmployeeAgency{}).AddForeignKey("agencyid", "table_agency(id)", "CASCADE", "RESTRICT")
		database.Model(&ent.TableEmployeeAgency{}).AddForeignKey("employeeid", "table_employee(id)", "CASCADE", "RESTRICT")
	}
}
