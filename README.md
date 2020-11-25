# Mikro 	


## About

Mikro adalah sebuah layanan pendampingan dan pelatihan untuk UMKM di Indonesia yang menjembatani Super Admin, Coach dan UMKM untuk berinteraksi dalam aplikasi tersebut. Repo ini adalah khusus backend.



## Technology Stack	

- Golang V1.12 or up
- MySQL Database



## Library

- Echo Labstack HTTP Client
- GORM
- Google Wire (Dependency Injection)
- JWT-GO
- Viper Config (TODO)
- Playgorund Validator (TODO)
- GoSwagger (TODO)
- Mockery (TODO)
- Testify (TODO)
- LogRush (TODO)



## Software Architecture

Pada repository ini, dibangun menggunakan Hexagonal / Clean Architecture dimana setiap layer di bagi-bagi menjadi beberapa layer inti yaitu :

- domain

- repository

- service

- api

  

## Unit Test

TODO



## How to use it ?

```g
git clone https://github.com/teten-nugraha/mikro-backend.git
```

Open your IDE,  dan buka terminal

```
go mod download
```

Untuk sementara buat sebuah database dengan nama **mikro_backend** dan sesuaikan **username** dan **password** dengan MySQL Local kalian di file **db/db_config.go** . kemudian running aplikasi.

```
go run main.go
```



## Bagaimana Cara Berkontribusi ?

Lakukan langkah-langkah berikut untuk dapat berkontrobusi dalam repo ini :

- Fork repository ini
- Kerjakan fitur yang akan dikontribusikan dalam repository anda hasil dari fork diatas
- Merge ke master repository Anda
- Buka repo ini https://github.com/teten-nugraha/mikro-backend.git kemudian create pull request



## Community Contributions Guide

Untuk lebih lanjut silahkan hubungi teten.nugraha18@gmail.com

