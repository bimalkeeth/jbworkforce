package entityrelations

import (
	"github.com/jinzhu/gorm"
	ent "jbworkforce/masterdataservice/master_dataaccess/entities"
)

//----------------------------------------------------
// Competency Type
//----------------------------------------------------
func CreateCompetencyType(database *gorm.DB) {
	if !database.HasTable(&ent.TableCompetencyType{}) {
		database.CreateTable(&ent.TableCompetencyType{})
		database.Model(&ent.TableCompetencyType{}).Association("EmployeeCompetency")
		database.Model(&ent.TableCompetencyType{}).AddUniqueIndex("idx_compentencytype_compentencyabbr", "compentencyabbr")
		database.Model(&ent.TableCompetencyType{}).AddUniqueIndex("idx_compentencytype_compentencyname", "compentencyname")

	}
}
