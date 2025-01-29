package repository

import (
	"github.com/YESUBZERO/consumer-service/internal/models"
	"gorm.io/gorm"
)

// ShipRepository maneja las operaciones en PostgreSQL
type ShipRepository struct {
	db *gorm.DB
}

func NewShipRepository(db *gorm.DB) *ShipRepository {
	return &ShipRepository{db: db}
}

// Verifica si un barco ya está almacenado en la BD
func (r *ShipRepository) ShipExists(imo int) bool {
	var count int64
	r.db.Model(&models.Ship{}).Where("imo = ?", imo).Count(&count)
	return count > 0
}

// Guarda el barco en la base de datos solo si viene de enriched-message
func (r *ShipRepository) SaveShip(ship models.Ship) error {
	return r.db.Create(&ship).Error
}
