package entityrelations

import (
	"github.com/jinzhu/gorm"
	ent "jbworkforce/masterdataservice/master_dataaccess/entities"
)

//----------------------------------------------------
// Allowance level
//----------------------------------------------------
func CreateClinicalStreamAttributes(database *gorm.DB) {
	if !database.HasTable(&ent.TableClinicalStreamAttribute{}) {
		database.CreateTable(&ent.TableClinicalStreamAttribute{})
		database.Model(&ent.TableClinicalStreamAttribute{}).AddForeignKey("clinicalstreamid", "table_clinicalstream(id)", "RESTRICT", "RESTRICT")

	}
}
