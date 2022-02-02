package jewelery

var daftarBarang = map[string]map[string]int{
	"Ruby": {
		"harga":  1000000,
		"jumlah": 5,
	},

	"Permata": {
		"harga":  3000000,
		"jumlah": 1,
	},
	"Topaz": {
		"harga":  1250000,
		"jumlah": 3,
	},
}

func countMaxPendapatan() int {
	var pendapatanTotal int
	for _, v := range daftarBarang {
		countjml := v["jumlah"] / 2

		pendapatanPerpasang := countjml * v["harga"]

		pendapatanTotal += pendapatanPerpasang
	}

	return pendapatanTotal
}
