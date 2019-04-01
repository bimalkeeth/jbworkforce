package entityrelations

import (
	"github.com/jinzhu/gorm"
	ent "jbworkforce/masterdataservice/master_dataaccess/entities"
)

//----------------------------------------------------
// Contact Type
//----------------------------------------------------
func CreateDepartment(database *gorm.DB) {

	if !database.HasTable(&ent.TableDepartment{}) {

		database.CreateTable(&ent.TableDepartment{})
		database.Model(&ent.TableDepartment{}).AddUniqueIndex("idx_table_department_departmentname", "departmentname")
		database.Model(&ent.TableDepartment{}).AddUniqueIndex("idx_table_department_departmentabbr", "departmentabbr")
		database.Model(&ent.TableDepartment{}).AddForeignKey("divisionid", "table_division(id)", "CASCADE", "RESTRICT")
		database.Model(&ent.TableDepartment{}).Association("SubDepartments")
	}
}
