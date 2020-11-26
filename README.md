# Mikro 	


## About

Mikro adalah sebuah layanan pendampingan dan pelatihan untuk UMKM di Indonesia yang menjembatani Super Admin, Coach dan UMKM untuk berinteraksi dalam aplikasi tersebut. Repo ini adalah khusus backend.



## Prerequisite

- Golang minimal V1.15 https://golang.org/

- MySQL Database https://www.mysql.com/downloads/

  

## Technology Stack	

- Golang V1.12 or up
- MySQL Database



## Library

- Echo Labstack HTTP Client
- GORM (Database ORM)
- Google Wire (Dependency Injection)
- JWT-GO
- Viper Config (TODO)
- Playgorund Validator (TODO)
- GoSwagger (TODO)
- Mockery (TODO)
- Testify (TODO)
- LogRush



## Software Architecture

Pada repository ini, dibangun menggunakan Hexagonal / Clean Architecture dimana setiap layer di bagi-bagi menjadi beberapa layer inti yaitu :

- domain
- repository
- service
- api

Untuk lebih jelasnya mengenai Clean Architecture bisa dibaca pada artikel ini mengenai kenapa kita harus menggunakan pendekatan ini :

- https://threedots.tech/post/introducing-clean-architecture/

- https://medium.com/swlh/clean-architecture-java-spring-fea51e26e00

  

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

- Clone repository ini

- Kerjakan fitur yang akan dikontribusikan dalam repository ini

- buatlah branch baru mengenai apa yang mau anda kerjakan dengan prefix **feature/** , misalkan anda akan mengerjakan fitur Forgot Password maka kerjakan di branch **feature/forgot-password**

- Push latest code ke branch tersebut

- Buka Repository ini di github kemudian create new **Pull Request **, target branch **main**, dan assign ke saya

  



## Community Contributions Guide

Untuk lebih lanjut silahkan hubungi teten.nugraha18@gmail.com



## Contributor

- Teten Nugraha
- Rifki Rahadian

