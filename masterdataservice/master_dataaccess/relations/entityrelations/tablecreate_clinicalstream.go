package entityrelations

import (
	"github.com/jinzhu/gorm"
	ent "jbworkforce/masterdataservice/master_dataaccess/entities"
)

//----------------------------------------------------
// Allowance level
//----------------------------------------------------
func CreateClinicalStream(database *gorm.DB) {
	if !database.HasTable(&ent.TableClinicalStream{}) {
		database.CreateTable(&ent.TableClinicalStream{})
		database.Model(&ent.TableClinicalStream{}).AddForeignKey("healthgroupid", "table_healthgroup(id)", "RESTRICT", "RESTRICT")
		database.Model(&ent.TableClinicalStream{}).AddForeignKey("activitytypeid", "table_activitytypes(id)", "RESTRICT", "RESTRICT")
		database.Model(&ent.TableClinicalStream{}).Association("ClinicalStreamAttributes")
	}
}
