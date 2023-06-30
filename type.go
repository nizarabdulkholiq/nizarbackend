package nizarbackend

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mahasiswa struct {
	ID            primitive.ObjectID `bson:"id,omitempty" json:"id,omitempty"`
	NamaMahasiswa string             `bson:"nama_mahasiswa" json:"nama_mahasiswa"`
	NIM           string             `bson:"nim" json:"nim"`
}
type TagihanRegistrasi struct {
	ID              primitive.ObjectID `bson:"id,omitempty" json:"id,omitempty"`
	NamaMahasiswa   string             `bson:"nama_mahasiswa" json:"nama_mahasiswa"`
	NIM             string             `bson:"nim" json:"nim"`
	Semester        string             `bson:"semester" json:"semester"`
	BiayaRegistrasi string             `bson:"biaya_registrasi" json:"biaya_registrasi"`
}
type TagihanSPP struct {
	ID            primitive.ObjectID `bson:"id,omitempty" json:"id,omitempty"`
	NamaMahasiswa string             `bson:"namamahasiswa" json:"namamahasiswa"`
	NIM           string             `bson:"nim" json:"nim"`
	Semester      string             `bson:"semester" json:"semester"`
	BiayaSPP      string             `bson:"biaya_spp" json:"biaya_spp"`
}
type TagihanMakanan struct {
	ID            primitive.ObjectID `bson:"id,omitempty" json:"id,omitempty"`
	NamaMahasiswa string             `bson:"nama_mahasiswa" json:"nama_mahasiswa"`
	NIM           string             `bson:"nim" json:"nim"`
	Bulan         string             `bson:"bulan" json:"bulan"`
	Tahun         int                `bson:"tahun" json:"tahun"`
	BiayaMakanan  string             `bson:"biaya_makanan" json:"biaya_makanan"`
}

type TagihanPerpustakaan struct {
	ID                primitive.ObjectID `bson:"id,omitempty" json:"id,omitempty"`
	NamaMahasiswa     string             `bson:"nama_mahasiswa" json:"nama_mahasiswa"`
	NIM               string             `bson:"nim" json:"nim"`
	Tahun             int                `bson:"tahun" json:"tahun"`
	BiayaPerpustakaan string             `bson:"biaya_perpustakaan" json:"biaya_perpustakaan"`
}
