package entityrelations

import (
	"github.com/jinzhu/gorm"
	ent "jbworkforce/masterdataservice/master_dataaccess/entities"
)

//----------------------------------------------------
// ClientSkillType
//----------------------------------------------------
func CreateClientSkillType(database *gorm.DB) {
	if !database.HasTable(&ent.TableClientSkillType{}) {

		database.CreateTable(&ent.TableClientSkillType{})
		database.Model(&ent.TableClientSkillType{}).AddForeignKey("clientskillgroupid", "table_clientskillgroup(id)", "RESTRICT", "RESTRICT")
	}
}
