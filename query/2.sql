SELECT mhs.Nama, COUNT(abs.Tanggal) FROM mahasiswa mhs RIGHT JOIN kehadiran abs ON mhs.NIM = abs.NIM  GROUP BY abs.NIM, mhs.Nama;;