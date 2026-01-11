package seeders

import "gorm.io/gorm"

type Registry struct {
	db *gorm.DB
}

// Run implements ISeedRegistry.
func (r *Registry) Run() {
	RunRoleSeeder(r.db)
	RunUserSeeder(r.db)
}

type ISeedRegistry interface {
	Run()
}

func NewSeedRegistry(db *gorm.DB) ISeedRegistry {
	return &Registry{
		db: db,
	}
}
