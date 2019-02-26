package entityrelations

import (
	"github.com/jinzhu/gorm"
	ent "jbworkforce/masterdataservice/master_dataaccess/entities"
)

//----------------------------------------------------
// ClientSkillTypes
//----------------------------------------------------
func CreateClientSkillGroup(database *gorm.DB) {
	if !database.HasTable(&ent.TableClientSkillGroup{}) {

		database.CreateTable(&ent.TableClientSkillGroup{})
		database.Model(&ent.TableClientSkillGroup{}).Association("ClientSkillTypes")
	}
}
