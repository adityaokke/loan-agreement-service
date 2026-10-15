package migration

import (
	"github.com/adityaokke/loan-agreement-service/internal/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&entity.Loan{}, &entity.LoanInvestment{})
	db.AutoMigrate(&entity.Investor{})
}
