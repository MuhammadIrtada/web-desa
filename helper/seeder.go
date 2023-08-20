package helper

import (
	"web-desa/config"
	"web-desa/model"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error){
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes),err
}

func SeederRefresh(cfg config.Config) {

	cfg.Database().Migrator().DropTable(&model.User{})
	cfg.Database().Migrator().DropTable(&model.Desa{})
	cfg.Database().Migrator().DropTable(&model.InfoKegiatan{})
	cfg.Database().Migrator().DropTable(&model.Umkm{})
	cfg.Database().Migrator().DropTable(&model.Wisata{})

	hashAdmin1, _ := HashPassword("Admin1**.")
	
	users := []model.User{
		{Username: "Admin1", Password: hashAdmin1},
	}

	desa := model.Desa{
		ID: 1,
		TentangDesa: "Desa sumberpucung meruakan sebuah desa yang terletak di indonesia provinsi jawa timur kabupaten malang kecamatan sumberpucung. Desa ini terbagi menjadi tiga dusun yaitu dusun suko, krajan, pakel. Di desa ini memiliki umkm unggulan seperti kolam pemancingan, krupuk pasir, tahu putih, dan blangkon. Desa Sumberpucung merupakan salah satu desa yang terletak di Kecamatan Sumberpucung, Kabupaten Malang. Desa Sumberpucung sendiri memiliki luas 463 KM² dengan 3117 rumah di dalamnya yang dibatasi oleh Desa Jatiguwi dan Desa Karangkates. Penduduk yang terdapat di Desa Sumberpucung sebesar 10.584 jiwa yang terdiri dari 5133 laki-laki dan 5451 perempuan. Salah satu produk unggulan dari desa Sumberpucung adalah sektor perikanan.",
		Visi: "1. Meningkatkan akses dan kualitas pendidikan, kesehatan dan layanan dasar lainnya bagi semua warga. 2. Mewujudkan kota produktif dan berdaya saing berbasis ekonomi kreatif, keberlanjutan dan keterpaduan. 3. Mewujudkan kota yang rukun dan toleran berazaskan keberagaman dan keberpihakan terhadap masyarakat rentan dan gender. 4. Memastikan kepuasan masyarakat atas layanan pemerintah yang tertib hukum, profesional dan akuntabel.",
		Misi: "Hakekat Bermartabat: Perwujudan dan Implementasi dari Kewajiban dan Tanggung Jawab Manusia Sebagai Khalifah, Kepada Masyarakat yang Dipimpin. Bermartabat Merujuk Pada Sebuah Nilai Harga Diri Kemanusiaan, yang Memiliki Arti Kemuliaan. Baldatun Thoyibatun Wa Robbun Ghofur: Tercipta Situasi, Kondisi, Tatanan Dan Karakter Yang Mulia Bagi Kota Malang Beserta Segenap Masyarakatnya",
	}

	infoKegiatans := []model.InfoKegiatan{
		{Judul: "Acara Penerimaan Mahasiswa Membangun Desa 2023 Universitas Brawijaya", Gambar: "https://wefjcxfoexcvmthzivgu.supabase.co/storage/v1/object/public/web-desa/informasi-kegiatan/acara pembukaan mmd.png", Tanggal: "03 Juli 2023", Deskripsi: "Desa Sumberpucung kali ini kedatangan tamu dari Mahasiswa Universitas Brawijaya. Dalam agenda pelaksanaan Mahasiswa Membangun Desa 2023 Universitas Brawijaya. Pelaksanaan MMD diawali dengan acara pembukaan yang dilakukan di Balai Desa Sumberpucung. Pak Hadi selaku Kepala Desa Sumberpucung ikut serta hadir dalam acara pembukaan ini."},
		{Judul: "Kunjungan Mayor pada Kegiatan Bedah Rumah di Desa Sumberpucung", Gambar: "https://wefjcxfoexcvmthzivgu.supabase.co/storage/v1/object/public/web-desa/informasi-kegiatan/bedah rumah.jpg", Tanggal: "04 Juli 2023", Deskripsi: "Kunjungan Mayor ke Desa Sumberpucung untuk meninjau langsung progres kegiatan bedah rumah."},
		{Judul: "Kegiatan Posyandu Balita Bulan Juli", Gambar: "https://wefjcxfoexcvmthzivgu.supabase.co/storage/v1/object/public/web-desa/informasi-kegiatan/posyandu.png", Tanggal: "11 Juli 2023", Deskripsi: "Kegiatan posyandu balita yang dilkukan secara rutin setiap bulan. Namun, ada yang berbeda pada pelaksanaan posyandu kali ini karena ibu-ibu posyandu ditemani oleh mahasiswa Universitas Brawijaya."},
		{Judul: "Rapat Persiapan Bersih Desa", Gambar: "https://wefjcxfoexcvmthzivgu.supabase.co/storage/v1/object/public/web-desa/informasi-kegiatan/rapat bersih desa.jpg", Tanggal: "11 Juli 2023", Deskripsi: "Untuk mempersiapakan acara Bersih Desa Sumberpucung 2023. Kepala Desa beserta jajarannya mengadakan rapat bersama Ketua RW dan RT. Membahas apa saya yang nantinya diperlukan saat pelaksanaan Bersih Desa."},
	}

	umkms := []model.Umkm{
		{Nama: "Ikha Flowers", Alamat: "Jl. Airlangga, Dusun Suko, Sumberpucung", Kontak: "Instagram: @ikha_flowers", Gambar: "https://wefjcxfoexcvmthzivgu.supabase.co/storage/v1/object/public/web-desa/umkm/umkm.jpg", Deskripsi: "Di Ikha Flowers kalian bisa melihat berbagai macam bunga buatan. Di sini kalian bisa memilih dari bunga yang berukuran kecil, sedang, hingga besar. Bentuk dan warna bisa dibuat sesuai yang diinginkan."},
		{Nama: "Tahu Bapak Muri", Alamat: "Jl. Nusantara, Dusun Suko, Sumberpucung", Kontak: "-", Gambar: "https://wefjcxfoexcvmthzivgu.supabase.co/storage/v1/object/public/web-desa/umkm/tahu.jpg", Deskripsi: "Bapak Muri merupakan pemilik usaha tahu. Tahu ini berasal dari olahan kedelai yang nantinya digoreng sehingga menghasilkan produk yaitu tahu goreng coklat."},
		{Nama: "Kerupuk Pasir Pak Edy", Alamat: "Jl. Kertanegara, Dusun Suko, Sumberpucung", Kontak: "-", Gambar: "https://wefjcxfoexcvmthzivgu.supabase.co/storage/v1/object/public/web-desa/umkm/kerupuk pasir.jpg", Deskripsi: "Kerupuk pasir pak Edy menjual berbagai jenis kerupuk pasir. Seperti kerupuk bawang, kerupuk upil, dan masih banyak lagi."},
	}

	wisatas := []model.Wisata{
		{Nama: "Kedung Gopit", HargaTiket: 5000, Alamat: "Dusun Krajan, Sumberpucung, Kec. Sumberpucung", Kontak: "-", Gambar: "https://wefjcxfoexcvmthzivgu.supabase.co/storage/v1/object/public/web-desa/wisata/kedung-gopit-kolam.jpg", JamBuka: "09:00", JamTutup: "16:00", Deskripsi: "Wisata Kedung Gopit merupakan wisata alam yang menyajikan keindahan alam dari aliran sungai. Tidak hanya itu disini juga terdapat kolam untuk anak-anak dan orang dewasa. Di sisi lain terdapat gazebo yang dapat digunakan untuk bersantai dengan keluarga"},
	}

	cfg.Database().Create(&desa)

	for _, user := range users {
		cfg.Database().Create(&user)
	}

	for _, infoKegiatan := range infoKegiatans {
		cfg.Database().Create(&infoKegiatan)
	}

	for _, umkm := range umkms {
		cfg.Database().Create(&umkm)
	}

	for _, wisata := range wisatas {
		cfg.Database().Create(&wisata)
	}

}
