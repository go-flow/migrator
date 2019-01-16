CREATE TABLE `users` (
    `id` int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT, 
    `first_name` varchar(255), 
    `last_name` varchar(255), 
    `job_title` varchar(255), 
    `birth_date` timestamp NULL, 
    `email` varchar(255) NOT NULL UNIQUE, 
    `phone_number` varchar(255), 
    `country` varchar(255), 
    `state` varchar(255), 
    `city` varchar(255), 
    `zip_code` varchar(255), 
    `address` varchar(255), 
    `bio` text, 
    `profile_image` varchar(255), 
    `google_id` varchar(50), 
    `facebook_id` varchar(50), 
    `password_hash` varchar(255), 
    `password_token` varchar(255), 
    `is_active` tinyint(4) NOT NULL DEFAULT 0, 
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB;

CREATE TABLE `roles` (
    `id` int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT, 
    `name` varchar(255) NOT NULL UNIQUE, 
    `description` varchar(255) NOT NULL DEFAULT '',
    `is_active` tinyint(4) NOT NULL DEFAULT 1, 
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB;

CREATE TABLE `permissions` (
    `id` int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT, 
    `code` varchar(255) NOT NULL UNIQUE, 
    `description` varchar(255) NOT NULL DEFAULT '',
    `is_active` tinyint(4) NOT NULL DEFAULT 1, 
    `is_system` tinyint(4) NOT NULL DEFAULT 0, 
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB;

CREATE TABLE `user_roles` (
    `user_id` int(11) NOT NULL, 
    `role_id` int(11) NOT NULL, 
    PRIMARY KEY(`user_id`, `role_id`)
) ENGINE=InnoDB;

CREATE TABLE `role_permissions` (
    `role_id` int(11) NOT NULL, 
    `permission_id` int(11) NOT NULL, 
    PRIMARY KEY(`role_id`, `permission_id`)
) ENGINE=InnoDB;

CREATE INDEX `idx_user_roles_user_id` ON `user_roles`(`user_id`);
CREATE INDEX `idx_user_roles_role_id` ON `user_roles`(`role_id`);

CREATE INDEX `idx_role_permissions_role_id` ON `role_permissions`(`role_id`);
CREATE INDEX `idx_role_permissions_permission_id` ON `role_permissions`(`permission_id`);

ALTER TABLE `user_roles` ADD CONSTRAINT `fk_user_roles_user_id` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`);
ALTER TABLE `user_roles` ADD CONSTRAINT `fk_user_roles_role_id` FOREIGN KEY (`role_id`) REFERENCES `roles`(`id`);

ALTER TABLE `role_permissions` ADD CONSTRAINT `fk_role_permissions_role_id` FOREIGN KEY (`role_id`) REFERENCES `roles`(`id`);
ALTER TABLE `role_permissions` ADD CONSTRAINT `fk_role_permissions_permission_id` FOREIGN KEY (`permission_id`) REFERENCES `permissions`(`id`);

