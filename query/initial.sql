CREATE DATABASE kawan_lama;
use kawan_lama;
CREATE TABLE kehadiran (id int primary key auto_increment,Tanggal date not null,NIM varchar(10) not null) engine=InnoDB;
CREATE TABLE peserta (id int primary key auto_increment,NIM varchar(10) not null,Nama varchar(100) not null) engine=InnoDB;
INSERT INTO kehadiran (Tanggal, NIM) VALUES("2020-06-15", "A001"), ("2020-06-18", "A003"),("2020-06-16", "A001"), ("2020-06-15", "A002"), ("2020-06-17", "A001"), ("2020-06-17", "A002"),("2020-06-15", "A003"), ("2020-06-16", "A002"), ("2020-06-17", "A004"), ("2020-06-16", "A003");
INSERT INTO peserta (NIM, Nama) VALUES("A001", "Mozarella"), ("A002", "Emmental"), ("A003", "Gouda"), ("A004", "Chevre");