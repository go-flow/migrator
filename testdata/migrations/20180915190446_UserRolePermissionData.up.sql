
INSERT INTO `roles` (`id`, `name`, `description`, `is_active`, `created_at`, `updated_at`) VALUES ('1', 'Admin', 'Admin', '1', NOW(), NOW());
INSERT INTO `roles` (`id`, `name`, `description`, `is_active`, `created_at`, `updated_at`) VALUES ('2', 'Editor', 'Default role', '1', NOW(), NOW());


INSERT INTO `permissions` (`id`, `code`, `is_active`, `is_system`, `created_at`, `updated_at`) VALUES ('1', 'PERMISSIONS', '1', '1', NOW(), NOW());
INSERT INTO `permissions` (`id`, `code`, `is_active`, `is_system`, `created_at`, `updated_at`) VALUES ('2', 'ROLES', '1', '1', NOW(), NOW());
INSERT INTO `permissions` (`id`, `code`, `is_active`, `is_system`, `created_at`, `updated_at`) VALUES ('3', 'USERS', '1', '1', NOW(), NOW());
INSERT INTO `permissions` (`id`, `code`, `is_active`, `is_system`, `created_at`, `updated_at`) VALUES ('4', 'WEBSITES', '1', '1', NOW(), NOW());


INSERT INTO `role_permissions` (`role_id`,`permission_id`) VALUES ('1','1');
INSERT INTO `role_permissions` (`role_id`,`permission_id`) VALUES ('1','2');
INSERT INTO `role_permissions` (`role_id`,`permission_id`) VALUES ('1','3');
INSERT INTO `role_permissions` (`role_id`,`permission_id`) VALUES ('1','4');


INSERT INTO `role_permissions` (`role_id`,`permission_id`) VALUES ('2','4');




