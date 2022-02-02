package refactor

/* Keselahannya terletak pada return dari function tsbt,
yang direturn bukan variable totla melainkan hasil yang dimana variable hasil valuenya tidak berubah
dibawah ini adalah perbaikan functionnya
*/
func perkalianSederhana(j, k int) (hasil int) {
	for j > 0 {
		hasil += k
		j--
	}

	return
}
