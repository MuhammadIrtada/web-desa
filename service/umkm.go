package service

import (
	"web-desa/model"
	"web-desa/request"
)

type umkmService struct {
	umkmRepository model.UmkmRepository
}


func NewumkmService(umkm model.UmkmRepository) model.UmkmService {
	return &umkmService{umkmRepository: umkm}
}

// DestroyUmkm implements model.UmkmService
func (u *umkmService) DestroyUmkm(id uint) error {
	umkm, _ := u.umkmRepository.FindByID(id)
	
	err := u.umkmRepository.Delete(umkm)
	if err != nil {
		return err
	}
	return nil
}

// EditUmkm implements model.UmkmService
func (u *umkmService) EditUmkm(id uint, req *request.UmkmRequest) (*model.Umkm, error) {
	newUmkm, err := u.umkmRepository.UpdateByID(&model.Umkm{
		ID:        id,
		Nama:      req.Nama,
		Alamat:    req.Alamat,
		Kontak:    req.Kontak,
		Gambar:    req.Gambar,
		Deskripsi: req.Deskripsi,
	})

	if err != nil {
		return nil, err
	}

	return newUmkm, err
}

// FetchUmkm implements model.UmkmService
func (u *umkmService) FetchUmkm() ([]*model.Umkm, error) {
	umkms, err := u.umkmRepository.Fetch()
	if err != nil {
		return nil, err
	}

	return umkms, err
}

// GetByID implements model.UmkmService
func (u *umkmService) GetByID(id uint) (*model.Umkm, error) {
	matkul, err := u.umkmRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return matkul, err
}

// StoreUmkm implements model.UmkmService
func (u *umkmService) StoreUmkm(req *request.UmkmRequest) (*model.Umkm, error) {
	umkm := &model.Umkm{
		Nama:      req.Nama,
		Alamat:    req.Alamat,
		Kontak:    req.Kontak,
		Gambar:    req.Gambar,
		Deskripsi: req.Deskripsi,
	}

	newUmkm, err := u.umkmRepository.Create(umkm)
	if err != nil {
		return nil, err
	}

	return newUmkm, nil
}