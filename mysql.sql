-- CREATED_AT = 21-02-2021 01:00:00
-- create table `blogs` with id using uuid
DROP TABLE IF EXISTS `blogs`;
CREATE TABLE `blogs` (
  `id` char(36) NOT NULL,
  `title` varchar(100) NOT NULL,
  `body` varchar(100) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;