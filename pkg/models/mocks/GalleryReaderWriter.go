// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	models "github.com/stashapp/stash/pkg/models"
	mock "github.com/stretchr/testify/mock"
)

// GalleryReaderWriter is an autogenerated mock type for the GalleryReaderWriter type
type GalleryReaderWriter struct {
	mock.Mock
}

// All provides a mock function with given fields:
func (_m *GalleryReaderWriter) All() ([]*models.Gallery, error) {
	ret := _m.Called()

	var r0 []*models.Gallery
	if rf, ok := ret.Get(0).(func() []*models.Gallery); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Gallery)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: newGallery
func (_m *GalleryReaderWriter) Create(newGallery models.Gallery) (*models.Gallery, error) {
	ret := _m.Called(newGallery)

	var r0 *models.Gallery
	if rf, ok := ret.Get(0).(func(models.Gallery) *models.Gallery); ok {
		r0 = rf(newGallery)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Gallery)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Gallery) error); ok {
		r1 = rf(newGallery)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByChecksum provides a mock function with given fields: checksum
func (_m *GalleryReaderWriter) FindByChecksum(checksum string) (*models.Gallery, error) {
	ret := _m.Called(checksum)

	var r0 *models.Gallery
	if rf, ok := ret.Get(0).(func(string) *models.Gallery); ok {
		r0 = rf(checksum)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Gallery)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(checksum)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByImageID provides a mock function with given fields: imageID
func (_m *GalleryReaderWriter) FindByImageID(imageID int) ([]*models.Gallery, error) {
	ret := _m.Called(imageID)

	var r0 []*models.Gallery
	if rf, ok := ret.Get(0).(func(int) []*models.Gallery); ok {
		r0 = rf(imageID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Gallery)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(imageID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByPath provides a mock function with given fields: path
func (_m *GalleryReaderWriter) FindByPath(path string) (*models.Gallery, error) {
	ret := _m.Called(path)

	var r0 *models.Gallery
	if rf, ok := ret.Get(0).(func(string) *models.Gallery); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Gallery)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindBySceneID provides a mock function with given fields: sceneID
func (_m *GalleryReaderWriter) FindBySceneID(sceneID int) (*models.Gallery, error) {
	ret := _m.Called(sceneID)

	var r0 *models.Gallery
	if rf, ok := ret.Get(0).(func(int) *models.Gallery); ok {
		r0 = rf(sceneID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Gallery)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(sceneID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindMany provides a mock function with given fields: ids
func (_m *GalleryReaderWriter) FindMany(ids []int) ([]*models.Gallery, error) {
	ret := _m.Called(ids)

	var r0 []*models.Gallery
	if rf, ok := ret.Get(0).(func([]int) []*models.Gallery); ok {
		r0 = rf(ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Gallery)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]int) error); ok {
		r1 = rf(ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: updatedGallery
func (_m *GalleryReaderWriter) Update(updatedGallery models.Gallery) (*models.Gallery, error) {
	ret := _m.Called(updatedGallery)

	var r0 *models.Gallery
	if rf, ok := ret.Get(0).(func(models.Gallery) *models.Gallery); ok {
		r0 = rf(updatedGallery)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Gallery)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Gallery) error); ok {
		r1 = rf(updatedGallery)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
