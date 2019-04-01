package entityrelations

import (
	"github.com/jinzhu/gorm"
	ent "jbworkforce/masterdataservice/master_dataaccess/entities"
)

//----------------------------------------------------
// Contact Type
//----------------------------------------------------
func CreateCostCentre(database *gorm.DB) {

	if !database.HasTable(&ent.TableCostCentre{}) {

		database.CreateTable(&ent.TableCostCentre{})
		database.Model(&ent.TableCostCentre{}).AddUniqueIndex("idx_table_costcentre_costcentrename", "costcentrename")
		database.Model(&ent.TableCostCentre{}).AddUniqueIndex("idx_table_costcentre_costcentreabbr", "costcentreabbr")
		database.Model(&ent.TableCostCentre{}).AddForeignKey("hospitalid", "table_hospital(id)", "CASCADE", "RESTRICT")
		database.Model(&ent.TableContactType{}).Association("SubDepartments")
		database.Model(&ent.TableContactType{}).Association("EmployeeEmpTypes")
	}
}
