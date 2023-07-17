package request

type (
	InfoKegiatanRequest struct {
		Judul     string `json:"judul" binding:"required"`
		Gambar    string `json:"gambar" binding:"required"`
		Tanggal   string `json:"tanggal" binding:"required"`
		Deskripsi string `json:"deskripsi" binding:"required"`
	}
)