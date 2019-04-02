package entityrelations

import (
	"github.com/jinzhu/gorm"
	ent "jbworkforce/masterdataservice/master_dataaccess/entities"
)

//----------------------------------------------------
// Employee Address
//----------------------------------------------------
func CreateEmployeeAddress(database *gorm.DB) {
	if !database.HasTable(&ent.TableEmployeeAddress{}) {
		database.CreateTable(&ent.TableEmployeeAddress{})
		database.Model(&ent.TableEmployeeAddress{}).AddForeignKey("addressid", "table_address(id)", "CASCADE", "RESTRICT")
		database.Model(&ent.TableEmployeeAddress{}).AddForeignKey("employeeid", "table_employee(id)", "CASCADE", "RESTRICT")
	}
}
