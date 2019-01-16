CREATE TABLE `websites` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(120) NOT NULL,
  `description` VARCHAR(255) NULL,
  `dev_url` VARCHAR(255),
  `prod_url` VARCHAR(255),
  `created_by` INT NOT NULL,
  `is_active` tinyint(4) NOT NULL DEFAULT 0, 
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `created_at` (`created_at`),
  KEY `updated_at` (`updated_at`),
  
  UNIQUE INDEX `idx_name_unique` (`name` ASC),
  
  FOREIGN KEY (`created_by`) REFERENCES users(`id`)
  )ENGINE=InnoDB;
