SELECT a.NIM, a.Nama,b.Tanggal FROM mahasiswa a RIGHT JOIN kehadiran b ON a.NIM = b.NIM WHERE b.Tanggal = "2020-06-15" OR b.Tanggal = "2020-06-17" ORDER BY b.Tanggal DESC ;

