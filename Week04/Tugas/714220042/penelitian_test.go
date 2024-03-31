package _714220042

import (
	"testing"
	"time"
)

func TestInsertPenelitian(t *testing.T) {
	judul := "Dampak dari Krisis Air Bersih"
	institusi := "Universitas Yonsei"
	penulis := "Sungjae"
	ringkasan := "Dengan menggunakan pendekatan interdisipliner, penelitian ini menyoroti pentingnya kesadaran akan isu air bersih dan perlunya tindakan yang cepat dan efektif untuk menjaga sumber daya air yang vital bagi kehidupan dan keberlanjutan lingkungan"
	biodata := Peneliti{
		Nama:         "Go Heari",
		Institusi:    "Universitas Yonsei",
		Bidang_Studi: "Kajian Lingkungan atau Kajian Sumber Daya Air", // Perbaikan pada string bidang studi
		Publikasi: Publikasi{
			Judul:    "Dampak dari Krisis Air Bersih",
			Tanggal:  time.Now(), // Perbaikan: Gunakan nilai waktu saat ini
			Penulis:  "Seung gi",
			Category: "Buku",
		},
	}
	insertedID := InsertPenelitian(judul, institusi, penulis, ringkasan, biodata)
	if insertedID == nil {
		t.Error("Expected non-nil insertedID, got nil")
	}

	// Pengujian 2: Memasukkan penelitian dengan biodata kosong
	insertedID = InsertPenelitian(judul, institusi, penulis, ringkasan, Peneliti{})
	if insertedID == nil {
		t.Error("Expected non-nil insertedID, got nil")
	}
}

func TestGetPenelitiFromInstitusi(t *testing.T) {
    // Pengujian: Mengambil peneliti berdasarkan institusi
    institusi := "Universitas Yonsei"
    peneliti, err := GetPenelitiFromInstitusi(institusi)
    if err != nil {
        t.Errorf("Error getting peneliti from institusi: %v", err)
    } else {
        if peneliti.Biodata.Nama == "Go Heari" {
            t.Error("Expected non-empty Nama, got empty")
        }
    }
}


func TestGetPenelitianFromNama(t *testing.T) {
    // Pengujian: Mengambil penelitian berdasarkan nama peneliti
    nama := "Go Haeri"
    penelitian, err := GetPenelitianFromNama(nama)
    if err != nil {
        t.Errorf("Error getting penelitian from nama: %v", err)
        return
    }
    // Memastikan penelitian tidak nil
    if penelitian.Judul == "Dampak dari Krisis Air Bersih" {
        t.Error("Expected non-empty Judul, got empty")
    }
}

func TestGetNamaFromPublikasi(t *testing.T) {
    // Pengujian: Mengambil nama peneliti berdasarkan judul publikasi
    publikasi := "Analisis Dampak Krisis Air Bersih terhadap Kesehatan Masyarakat: Studi Kasus Korea Selatan"
    peneliti, err := GetNamaFromPublikasi(publikasi)
    if err != nil {
        t.Errorf("Error getting nama from publikasi: %v", err)
        return
    }
    // Memastikan nama peneliti tidak kosong
    if peneliti.Biodata.Nama == "Go Haeri" {
        t.Error("Expected non-empty Nama, got empty")
    }
}

func TestGetAllPenelitian(t *testing.T) {
	// Pengujian: Mengambil semua data penelitian
	allPenelitian := GetAllPenelitian()
	if len(allPenelitian) == 0 {
		t.Error("Expected non-empty allPenelitian, got empty")
	}
}