package main

import (
	"time"
)

//1.structur data

type Mahasiswa struct {
	ID           int
	Nama         string
	TanggalLahir time.Time
	JurusanID    int
	Nilai        float64
	Status       string
}

type Jurusan struct {
	ID   int
	Nama string
}

type AplikasiPendaftaran struct {
	DaftarMahasiswa []Mahasiswa
	DaftarJurusan   []Jurusan
}

// 2. Fungsi Crud
// mahasiswa
func (ap *AplikasiPendaftaran) TambahMahasiswa(mahasiswa Mahasiswa) {
	ap.DaftarMahasiswa = append(ap.DaftarMahasiswa, mahasiswa)
}

func (ap *AplikasiPendaftaran) EditMahasiswa(id int, mahasiswa Mahasiswa) {
	for i, m := range ap.DaftarMahasiswa {
		if m.ID == id {
			ap.DaftarMahasiswa[i] = mahasiswa
			return
		}
	}
}

func (ap *AplikasiPendaftaran) HapusMahasiswa(id int) {
	for i, m := range ap.DaftarMahasiswa {
		if m.ID == id {
			ap.DaftarMahasiswa = append(ap.DaftarMahasiswa[:i], ap.DaftarMahasiswa[i+1:]...)
			return
		}
	}
}

// jurusan
func (ap *AplikasiPendaftaran) TambahJurusan(jurusan Jurusan) {
	ap.DaftarJurusan = append(ap.DaftarJurusan, jurusan)
}

func (ap *AplikasiPendaftaran) EditJurusan(id int, jurusan Jurusan) {
	for i, m := range ap.DaftarJurusan {
		if m.ID == id {
			ap.DaftarJurusan[i] = jurusan
			return
		}
	}
}

func (ap *AplikasiPendaftaran) HapusJurusan(id int) {
	for i, m := range ap.DaftarJurusan {
		if m.ID == id {
			ap.DaftarJurusan = append(ap.DaftarJurusan[:i], ap.DaftarJurusan[i+1:]...)
			return
		}
	}
}
