package models

import (
	"github.com/jmoiron/sqlx"
)

type StudioReader interface {
	Find(id int) (*Studio, error)
	FindMany(ids []int) ([]*Studio, error)
	// FindChildren(id int) ([]*Studio, error)
	// FindBySceneID(sceneID int) (*Studio, error)
	FindByName(name string, nocase bool) (*Studio, error)
	// Count() (int, error)
	All() ([]*Studio, error)
	// AllSlim() ([]*Studio, error)
	// Query(studioFilter *StudioFilterType, findFilter *FindFilterType) ([]*Studio, int)
	GetStudioImage(studioID int) ([]byte, error)
}

type StudioWriter interface {
	Create(newStudio Studio) (*Studio, error)
	Update(updatedStudio StudioPartial) (*Studio, error)
	UpdateFull(updatedStudio Studio) (*Studio, error)
	// Destroy(id string) error
	UpdateStudioImage(studioID int, image []byte) error
	// DestroyStudioImage(studioID int) error
}

type StudioReaderWriter interface {
	StudioReader
	StudioWriter
}

func NewStudioReaderWriter(tx *sqlx.Tx) StudioReaderWriter {
	return &studioReaderWriter{
		tx: tx,
		qb: NewStudioQueryBuilder(),
	}
}

type studioReaderWriter struct {
	tx *sqlx.Tx
	qb StudioQueryBuilder
}

func (t *studioReaderWriter) Find(id int) (*Studio, error) {
	return t.qb.Find(id, t.tx)
}

func (t *studioReaderWriter) FindMany(ids []int) ([]*Studio, error) {
	return t.qb.FindMany(ids)
}

func (t *studioReaderWriter) FindByName(name string, nocase bool) (*Studio, error) {
	return t.qb.FindByName(name, t.tx, nocase)
}

func (t *studioReaderWriter) All() ([]*Studio, error) {
	return t.qb.All()
}

func (t *studioReaderWriter) GetStudioImage(studioID int) ([]byte, error) {
	return t.qb.GetStudioImage(studioID, t.tx)
}

func (t *studioReaderWriter) Create(newStudio Studio) (*Studio, error) {
	return t.qb.Create(newStudio, t.tx)
}

func (t *studioReaderWriter) Update(updatedStudio StudioPartial) (*Studio, error) {
	return t.qb.Update(updatedStudio, t.tx)
}

func (t *studioReaderWriter) UpdateFull(updatedStudio Studio) (*Studio, error) {
	return t.qb.UpdateFull(updatedStudio, t.tx)
}

func (t *studioReaderWriter) UpdateStudioImage(studioID int, image []byte) error {
	return t.qb.UpdateStudioImage(studioID, image, t.tx)
}
