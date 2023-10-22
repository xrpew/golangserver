DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` varchar(30) unsigned NOT NULL AUTO_INCREMENT,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT NOW() NOT
);
