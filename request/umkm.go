package request

type (
	UmkmRequest struct {
		Nama       string `json:"nama" binding:"required"`
		Alamat     string `json:"alamat" binding:"required"`
		Kontak     string `json:"kontak" binding:"required"`
		Gambar     string `json:"gambar" binding:"required"`
		Deskripsi  string `json:"deskripsi" binding:"required"`
	}
)
