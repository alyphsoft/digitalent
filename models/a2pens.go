package models

type A2pens struct {
	Nomor         uint `gorm:"primaryKey; autoIncrement:true`
	Tahun         string
	Cabang_taspen string
	Nama          string
	Nip           string
	Is_cetak      bool
}
