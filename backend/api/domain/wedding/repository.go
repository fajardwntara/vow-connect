package wedding

import "gorm.io/gorm"

type weddingRepository struct {
	db *gorm.DB
}

func NewWeddingRepository(db *gorm.DB) *weddingRepository {
	return &weddingRepository{db}
}

// Create implements WeddingRepository.
func (w *weddingRepository) Create(wd *Wedding) error {
	panic("unimplemented")
}

// Delete implements WeddingRepository.
func (w *weddingRepository) Delete(id uint) error {
	panic("unimplemented")
}

// Update implements WeddingRepository.
func (w *weddingRepository) Update(wd *Wedding) error {
	panic("unimplemented")
}
