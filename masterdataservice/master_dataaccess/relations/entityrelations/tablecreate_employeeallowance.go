package entityrelations

import (
	"github.com/jinzhu/gorm"
	ent "jbworkforce/masterdataservice/master_dataaccess/entities"
)

//----------------------------------------------------
// EmployeeAllowance
//----------------------------------------------------
func CreateEmployeeAllowance(database *gorm.DB) {
	if !database.HasTable(&ent.TableEmployeeAllowance{}) {
		database.CreateTable(&ent.TableEmployeeAllowance{})
		database.Model(&ent.TableEmployeeAllowance{}).AddForeignKey("allowanceid", "table_allowance(id)", "CASCADE", "RESTRICT")
		database.Model(&ent.TableEmployeeAllowance{}).AddForeignKey("employeeid", "table_employee(id)", "CASCADE", "RESTRICT")
	}
}
