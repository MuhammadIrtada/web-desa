package response

import (
	"web-desa/model"
)

type (
	DesaResponse struct {
		ID           uint                 `gorm:"primaryKey"`
		TentangDesa  string               `json:"tentang_desa"`
		Visi         string               `json:"visi"`
		Misi         string               `json:"misi"`
		InfoKegiatan []*model.InfoKegiatan `json:"info_kegiatan"`
		Umkm         []*model.Umkm         `json:"umkm"`
		Wisata       []*model.Wisata       `json:"wisata"`
	}
)

func ConvertToDesaResponse(d model.Desa, i []*model.InfoKegiatan, u []*model.Umkm, w []*model.Wisata) DesaResponse {
	return DesaResponse{
		ID:           d.ID,
		TentangDesa:  d.TentangDesa,
		Visi:         d.Visi,
		Misi:         d.Misi,
		InfoKegiatan: i,
		Umkm:         u,
		Wisata:       w,
	}
}
