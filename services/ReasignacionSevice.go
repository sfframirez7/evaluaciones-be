package services

import (
	"time"

	"../config"
	"../models"
)

func NewReasignacion(Modelo models.Reasignacion) (models.Reasignacion, error) {

	Modelo.IdAsignacion = 0
	Modelo.FechaModificacion = time.Now()

	db := config.ConnectDB()
	defer db.Close()

	db.Create(&Modelo)

	return Modelo, nil
}

func GetReasignaciones() ([]models.ReasignacionDetallada, error) {

	var result []models.ReasignacionDetallada

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("exec usp_GetReasignaciones").Scan(&result)

	return result, nil
}

func DeleteReasignacion(IdAsignacion string) (bool, error) {

	var result []models.Reasignacion

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("exec usp_RemoveAsignacion ?", IdAsignacion).Scan(&result)

	return true, nil
}
