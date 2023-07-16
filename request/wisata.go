package request

type (
	WisataRequest struct {
		Nama       string `json:"nama" binding:"required"`
		HargaTiket int    `json:"harga_tiket" binding:"required"`
		Alamat     string `json:"alamat" binding:"required"`
		Kontak     string `json:"kontak" binding:"required"`
		Gambar     string `json:"gambar" binding:"required"`
		JamBuka    string `json:"jam_buka" binding:"required"`
		JamTutup   string `json:"jam_tutup" binding:"required"`
		Deskripsi  string `json:"deskripsi" binding:"required"`
	}
)
