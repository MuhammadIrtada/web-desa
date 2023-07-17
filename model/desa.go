package model

import (
	"web-desa/request"

	"github.com/gin-gonic/gin"
)

type (
	Desa struct {
		ID          uint   `gorm:"primaryKey"`
		TentangDesa string `json:"tentang_desa"`
		Visi        string `json:"visi"`
		Misi        string `json:"misi"`
	}

	DesaRepository interface {
		Create(desa *Desa) (*Desa, error)
		Fetch() (*Desa, error)
		Update(desa *Desa) (*Desa, error)
		Delete(desa *Desa) (*Desa, error)
	}

	DesaService interface {
		StoreDesa(req *request.DesaRequest) (*Desa, error)
		FetchDesa() (*Desa, error)
		EditDesa(id uint, req *request.DesaRequest) (*Desa, error)
		DestroyDesa() error
	}

	DesaHandler interface {
		Mount(group *gin.RouterGroup)
	}
)