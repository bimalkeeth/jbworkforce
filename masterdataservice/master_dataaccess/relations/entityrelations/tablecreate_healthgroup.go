package entityrelations

import (
	"github.com/jinzhu/gorm"
	ent "jbworkforce/masterdataservice/master_dataaccess/entities"
)

//----------------------------------------------------
// Allowance level
//----------------------------------------------------
func CreateHealthGroup(database *gorm.DB) {
	if !database.HasTable(&ent.TableHealthGroup{}) {
		database.CreateTable(&ent.TableHealthGroup{})
		database.Model(&ent.TableHealthGroup{}).Association("ClinicalStreams")
	}
}
