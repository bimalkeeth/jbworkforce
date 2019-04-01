package entityrelations

import (
	"github.com/jinzhu/gorm"
	ent "jbworkforce/masterdataservice/master_dataaccess/entities"
)

//----------------------------------------------------
// Contact Type
//----------------------------------------------------
func CreateHospital(database *gorm.DB) {

	if !database.HasTable(&ent.TableHospital{}) {

		database.CreateTable(&ent.TableHospital{})
		database.Model(&ent.TableHospital{}).AddUniqueIndex("idx_table_hospital_hospitalname", "hospitalname")
		database.Model(&ent.TableHospital{}).AddUniqueIndex("idx_table_hospital_hospitalabbr", "hospitalabbr")
		database.Model(&ent.TableHospital{}).AddForeignKey("organisationid", "table_organisation(id)", "CASCADE", "RESTRICT")
		database.Model(&ent.TableHospital{}).Association("CostCenters")
	}
}
