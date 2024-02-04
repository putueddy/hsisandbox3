package tugas2

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

func GetSemester(bulan int) int {
	if bulan >= 1 && bulan <= 6 {
		return 1
	}
	return 2
}

func GeneratorNIK(gender string, tahun int, jumlah_yang_digenerate int) ([]string, error) {
	// NIK 12 digit
	// 2 digit pertama adalah AR a.k.a Abdullah Roy
	// 1 digit berikutnya N: 'male' a.k.a ikhwan; atau T: 'female' a.k.a akhwat
	// 2 digit berikutnya adalah tahun 2 digit terakhir
	// 1 digit berikutnya adalah 1: (bulan 1-6); atau 2: (bulan 7-12) a.k.a semester
	// 1 digit berikutnya adalah tanda '-' a.k.a strip
	// 5 digit terakhir adalah nomor urut

	if gender != "female" && gender != "male" {
		return nil, fmt.Errorf("parameter gender tidak valid")
	}

	if tahun < 0 {
		return nil, fmt.Errorf("parameter tahun tidak valid")
	}

	if jumlah_yang_digenerate < 0 {
		return nil, fmt.Errorf("parameter jumlah_yang_digenerate tidak valid")
	}

	prefix := "ARN"

	if gender == "female" {
		prefix = "ART"
	}

	nomorUrut := 1 + rand.Intn(99999-1)

	var niks []string

	for i := nomorUrut; i < nomorUrut+jumlah_yang_digenerate; i++ {
		if i > 99999 {
			break
		}
		nik := fmt.Sprintf("%s%d%d-%05d", prefix, tahun%100, GetSemester(int(time.Now().Month())), i)
		niks = append(niks, nik)
	}

	return niks, nil
}

func NikBerikut(nikSebelum string, jumlah_yang_digenerate int) ([]string, error) {
	if len(nikSebelum) != 12 {
		return nil, fmt.Errorf("parameter NIK tidak valid")
	}

	if jumlah_yang_digenerate < 0 {
		return nil, fmt.Errorf("parameter jumlah_yang_digenerate tidak valid")
	}

	// Ekstraksi 3 digit pertama
	prefix := nikSebelum[:3]
	if prefix != "ARN" && prefix != "ART" {
		return nil, fmt.Errorf("parameter NIK[:3] tidak valid")
	}

	// Ekstraksi 2 digit berikutnya, a.k.a tahun
	tahun, err := strconv.Atoi(nikSebelum[3:5])
	if err != nil {
		return nil, fmt.Errorf("parameter NIK[3:5] tidak valid")
	}

	// Ekstraksi 1 digit berikunya, a.k.a semester
	semester, err := strconv.Atoi(nikSebelum[5:6])
	if err != nil && semester >= 1 && semester <= 2 {
		return nil, fmt.Errorf("parameter NIK[5:6] tidak valid")
	}

	// Ekstraksi 5 digit terakhir, a.k.a nomor urut
	nomorUrut := nikSebelum[7:]
	nomor, err := strconv.Atoi(nomorUrut)
	if err != nil {
		return nil, fmt.Errorf("parameter NIK[7:] tidak valid")
	}

	var niks []string

	for i := nomor; i < nomor+jumlah_yang_digenerate; i++ {
		if i > 99999 {
			break
		}
		nik := fmt.Sprintf("%s%d%d-%05d", prefix, tahun, semester, i)
		niks = append(niks, nik)
	}

	return niks, nil
}

func KelompokHalaqah(niks []string) []string {
	// Interpretasi:
	// - Kelompokkan NIK berdasarkan gender, angkatan
	// - Urutkan NIK dalam kelompok tahun berdasarkan 5 digit terakhir
	// - Gabungkan kembali dengan struktur:
	//  1. NIK admin angkatan sesudahnya adalah ikhwan atau akhwat dengan nip paling kecil
	//  2. NIK anggota setelah admin

	// Buat kelompok sesuai gender
	genders := make(map[string][]string)
	for _, nik := range niks {
		gender := nik[:3]
		genders[gender] = append(genders[gender], nik)
	}

	// Urutkan NIK per tahun berdasarkan 5 digit terakhir
	for _, niks := range genders {
		sort.Slice(niks, func(i, j int) bool {
			return niks[i][3:6] < niks[j][3:6]
		})
	}

	// Gabungkan kembali dengan struktur yang diminta
	var result []string
	for _, niks := range genders {
		result = append(result, niks[0]) // admin

		for i := 1; i < len(niks); i++ {
			result = append(result, niks[i]) // anggota
		}
	}

	return result
}
