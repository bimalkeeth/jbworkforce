package entityrelations

import (
	"github.com/jinzhu/gorm"
	ent "jbworkforce/masterdataservice/master_dataaccess/entities"
)

//----------------------------------------------------
// Allowance level
//----------------------------------------------------
func CreateCompetencyCompletionStatus(database *gorm.DB) {
	if !database.HasTable(&ent.TableCompetencyCompletionStatus{}) {
		database.CreateTable(&ent.TableCompetencyCompletionStatus{})
	}
}
