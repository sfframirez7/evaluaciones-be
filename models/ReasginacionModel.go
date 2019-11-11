package models

import (
	"time"

	_ "github.com/jinzhu/gorm"
)

type Reasignacion struct {
	IdAsignacion      int       `gorm:"column:idAsignacion;AUTO_INCREMENT"`
	IdColaborador     string    `gorm:"column:idColaborador"`
	IdSubAreaResta    int       `gorm:"column:idSubAreaResta"`
	IdSubAreaSuma     int       `gorm:"column:idSubAreaSuma"`
	IdUsuarioModifico string    `gorm:"column:idUsuarioModifico"`
	FechaModificacion time.Time `gorm:"column:FechaModificacion"`
	Activo            bool      `gorm:"column:Activo;DEFAULT:1"`
}

type ReasignacionDetallada struct {
	Reasignacion

	SubAreaOrigen     string `gorm:"column:SubAreaOrigen"`
	SubAreaDestino    string `gorm:"column:SubAreaDestino"`
	NombreColaborador string `gorm:"column:NombreColaborador"`
	Modifica          string `gorm:"column:Modifica"`
}

func (Reasignacion) TableName() string {
	return "dbReasignaciones"
}
