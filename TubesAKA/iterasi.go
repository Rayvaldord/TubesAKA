package main
import (
	"fmt"
	"time"
)

type City struct {
	Name     string
	Distance float64
}

type TabArray [100]City

func main() {
	start := time.Now() // Mulai pengukuran waktu

	// Jumlah kota yang ada dalam data
	var n int = 6
	// Data kota dan jaraknya
	var cities TabArray = TabArray{
		{"Jakarta", 385},
		{"Bandung", 240},
		{"Surabaya", 467},
		{"Medan", 2326},
		{"Yogyakarta", 164},
		{"Denpasar", 829},
		{"Purworejo", 111},
		{"Malang", 469},
		{"Bekasi", 367},
		{"Jayapura", 4860},
		{"Semarang", 182},
		{"Banyuwangi", 708},
		{"Purbalingga", 18},
		{"Tegal", 101},
		{"Kebumen", 69},
		{"Surakarta", 216},
		{"Cilacap", 49},
		{"Banjarnegara", 40},
		{"Wonosobo", 87},
		{"Lombok", 969},
		{"Bogor", 361},
		{"Makassar", 1288},
		{"Pekanbaru", 1686},
		{"Manado", 2698},
		{"Pekalongan", 118},
		{"Magelang", 136},
		
	}

	// Menampilkan data awal sebelum sorting
	fmt.Println("Data Kota Sebelum Diurutkan:")
	for i := 0; i < n; i++ {
		fmt.Printf("%s: %.2f km\n", cities[i].Name, cities[i].Distance)
	}

	// Sorting menggunakan prosedur khusus
	sortCities(&cities, n)

	// Menampilkan hasil setelah sorting
	fmt.Println("------------------------------------")
	fmt.Println("Setelah Diurutkan Berdasarkan Jarak Terdekat Dari Purwokerto:")
	for i := 0; i < n; i++ {
		fmt.Printf("%s: %.2f km\n", cities[i].Name, cities[i].Distance)
	}

	// Akhir pengukuran waktu
	end := time.Now()
	fmt.Printf("Waktu eksekusi: %v\n", end.Sub(start))
}

func sortCities(cities *TabArray, n int) {
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if cities[j].Distance > cities[j+1].Distance {
				// Tukar posisi
				cities[j], cities[j+1] = cities[j+1], cities[j]
			}
		}
	}
}