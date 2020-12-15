CREATE TABLE `merchants` (
  `id` varbinary(255) NOT NULL,
  `name` varchar(255) UNIQUE NOT NULL,
  `user_id` int DEFAULT NULL,
  `category_id` int DEFAULT NULL,
  `business_started_date` date DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `business_address` varchar(255) UNIQUE DEFAULT NULL,
  `omzet` double UNIQUE DEFAULT NULL,
  `profit` double DEFAULT NULL,
  `business_problem` text DEFAULT NULL,
  `employee_count` int DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;