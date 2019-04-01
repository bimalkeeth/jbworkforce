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
		database.Model(&ent.TableDivision{}).AddUniqueIndex("idx_table_department_departmentname", "departmentname")
		database.Model(&ent.TableDivision{}).AddUniqueIndex("idx_table_division_divisionabbr", "divisionabbr")
		database.Model(&ent.TableDivision{}).AddForeignKey("campusid", "table_campus(id)", "CASCADE", "RESTRICT")
		database.Model(&ent.TableDivision{}).Association("Campus")
	}
}
