package service

import (
	"web-desa/model"
	"web-desa/request"
)

type desaService struct {
	desaRepository         model.DesaRepository
	infoKegiatanRepository model.InfoKegiatanRepository
	umkmRepository         model.UmkmRepository
	wisataRepository       model.WisataRepository
}

func NewDesaService(desa model.DesaRepository, infoKegiatan model.InfoKegiatanRepository, umkm model.UmkmRepository, wisata model.WisataRepository) model.DesaService {
	return &desaService{
		desaRepository:         desa,
		infoKegiatanRepository: infoKegiatan,
		umkmRepository:         umkm,
		wisataRepository:       wisata,
	}
}

func (s *desaService) StoreDesa(req *request.DesaRequest) (*model.Desa, error) {
	desa := &model.Desa{
		ID:          1,
		TentangDesa: req.TentangDesa,
		Visi:        req.Visi,
		Misi:        req.Misi,
	}

	newDesa, err := s.desaRepository.Create(desa)
	if err != nil {
		return nil, err
	}

	return newDesa, nil
}

func (s *desaService) FetchDesa() (*model.Desa, []*model.InfoKegiatan, []*model.Umkm, []*model.Wisata, error) {
	desa, err := s.desaRepository.Fetch()
	if err != nil {
		return nil, nil, nil, nil, err
	}

	infoKegiatan, err := s.infoKegiatanRepository.GetLimitedInfoKegiatan(3)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	umkm, err := s.umkmRepository.GetLimitedUmkm(3)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	wisata, err := s.wisataRepository.GetLimitedWisata(4)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return desa, infoKegiatan, umkm, wisata, nil
}

func (s *desaService) EditDesa(id uint, req *request.DesaRequest) (*model.Desa, error) {
	newDesa, err := s.desaRepository.Update(&model.Desa{
		ID:          id,
		TentangDesa: req.TentangDesa,
		Visi:        req.Visi,
		Misi:        req.Misi,
	})

	if err != nil {
		return nil, err
	}

	return newDesa, nil
}

func (s *desaService) DestroyDesa() error {
	desa, _ := s.desaRepository.Fetch()

	_, err := s.desaRepository.Delete(desa)
	if err != nil {
		return err
	}

	return nil
}
