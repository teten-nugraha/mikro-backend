-- mikro_backend.users definition

CREATE TABLE `users` (
  `id` varbinary(255) NOT NULL,
  `username` varchar(255) UNIQUE NOT NULL,
  `password` varchar(255) DEFAULT NULL,
  `fullname` varchar(255) DEFAULT NULL,
  `role` varchar(20) DEFAULT NULL,
  `is_active` tinyint(1) DEFAULT NULL,
  `email` varchar(255) UNIQUE DEFAULT NULL,
  `phone` varchar(255) UNIQUE DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- mikro_backend.kategoris definition

CREATE TABLE `kategoris` (
  `id` varbinary(255) NOT NULL,
  `kategori` varchar(30) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_kategoris_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO mikro_backend.users (id,username,password,fullname,`role`,is_active,email,phone,created_at,updated_at,deleted_at) VALUES
(0x66313036623562392D363061382D343337652D623264362D616433386264333032376539,'gorm','$2a$10$.qrJ1LqsyYFlf.vCJNtT0.nwkHWV97DAhI.nJCC8gfxyCa1uH24de','Go ORM','MIKRO_USER',1,'gorm@mail.com','081234586','2020-12-03 09:22:23.0','2020-12-03 09:22:23.0',NULL)
;