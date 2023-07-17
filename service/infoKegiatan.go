package service

import (
	"web-desa/model"
	"web-desa/request"
)

type infoKegiatanService struct {
	infoKegiatanRepository model.InfoKegiatanRepository
}

func NewInfoKegiatanService(infoKegiatan model.InfoKegiatanRepository) model.InfoKegiatanService {
	return &infoKegiatanService{infoKegiatanRepository: infoKegiatan}
}

func (s *infoKegiatanService) StoreInfoKegiatan(req *request.InfoKegiatanRequest) (*model.InfoKegiatan, error) {
	infoKegiatan := &model.InfoKegiatan{
		Judul: req.Judul,
		Gambar: req.Gambar,
		Tanggal: req.Tanggal,
		Deskripsi: req.Deskripsi,
	}

	newInfoKegiatan, err := s.infoKegiatanRepository.Create(infoKegiatan)
	if err != nil {
		return nil, err
	}

	return newInfoKegiatan, nil
}

func (s *infoKegiatanService) EditInfoKegiatan(id uint, req *request.InfoKegiatanRequest) (*model.InfoKegiatan, error) {
	newDesa, err := s.infoKegiatanRepository.UpdateByID(&model.InfoKegiatan{
		ID: id,
		Judul: req.Judul,
		Gambar: req.Gambar,
		Tanggal: req.Tanggal,
		Deskripsi: req.Deskripsi,
	})

	if err != nil {
		return nil, err
	}

	return newDesa, nil
}

func (s *infoKegiatanService) GetByID(id uint) (*model.InfoKegiatan, error) {
	infoKegiatan, err := s.infoKegiatanRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return infoKegiatan, err
}

func (s *infoKegiatanService) DestroyInfoKegiatan(id uint) error {
	infoKegiatan, _ := s.infoKegiatanRepository.FindByID(id)

	_, err := s.infoKegiatanRepository.Delete(infoKegiatan)
	if err != nil {
		return err
	}

	return nil
}

func (s *infoKegiatanService) FetchInfoKegiatan() ([]*model.InfoKegiatan, error) {
	infoKegiatans, err := s.infoKegiatanRepository.Fetch()
	if err != nil {
		return nil, err
	}

	return infoKegiatans, nil
}