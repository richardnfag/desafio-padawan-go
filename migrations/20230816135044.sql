-- Create "currencies" table
CREATE TABLE `currencies` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  `code` varchar(191) NOT NULL,
  `symbol` longtext NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `code` (`code`),
  INDEX `idx_currencies_deleted_at` (`deleted_at`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "conversions" table
CREATE TABLE `conversions` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  `amount` double NULL,
  `from_currency_id` bigint unsigned NULL,
  `to_currency_id` bigint unsigned NULL,
  `rate` double NULL,
  `result` double NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_conversions_from` (`from_currency_id`),
  INDEX `fk_conversions_to` (`to_currency_id`),
  INDEX `idx_conversions_deleted_at` (`deleted_at`),
  CONSTRAINT `fk_conversions_from` FOREIGN KEY (`from_currency_id`) REFERENCES `currencies` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT `fk_conversions_to` FOREIGN KEY (`to_currency_id`) REFERENCES `currencies` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
