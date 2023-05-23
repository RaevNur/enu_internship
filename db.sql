-- users definition

CREATE TABLE `users` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `username` varchar(12) NOT NULL,
  `firstname` varchar(255) NOT NULL,
  `lastname` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- sessions definition

CREATE TABLE `sessions` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `uuid` varchar(255) NOT NULL,
  `user_id` bigint NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE,
  CONSTRAINT `sessions_fk1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO users (username,firstname,lastname,email,password,created_at) VALUES
	 ('015763438092','John','Doe','johndoe@example.com','$2a$10$i6RPPHA95dB.F8KkiNQ1C.DSeZGdRMdc7BBPlkMPfg9nvSAqWTfxy','2023-05-23 11:09:57'),
	 ('987623948765','Jane','Smith','janesmith@example.com','$2a$10$i6RPPHA95dB.F8KkiNQ1C.DSeZGdRMdc7BBPlkMPfg9nvSAqWTfxy','2023-05-23 11:09:57'),
	 ('089675435098','Michael','Johnson','michaeljohnson@example.com','$2a$10$i6RPPHA95dB.F8KkiNQ1C.DSeZGdRMdc7BBPlkMPfg9nvSAqWTfxy','2023-05-23 11:09:57'),
	 ('239456897523','Emily','Williams','emilywilliams@example.com','$2a$10$i6RPPHA95dB.F8KkiNQ1C.DSeZGdRMdc7BBPlkMPfg9nvSAqWTfxy','2023-05-23 11:09:57'),
	 ('659278405187','David','Brown','davidbrown@example.com','$2a$10$i6RPPHA95dB.F8KkiNQ1C.DSeZGdRMdc7BBPlkMPfg9nvSAqWTfxy','2023-05-23 11:09:57'),
	 ('123457682903','Sarah','Jones','sarahjones@example.com','$2a$10$i6RPPHA95dB.F8KkiNQ1C.DSeZGdRMdc7BBPlkMPfg9nvSAqWTfxy','2023-05-23 11:09:57'),
	 ('432129750387','Christopher','Clark','christopherclark@example.com','$2a$10$i6RPPHA95dB.F8KkiNQ1C.DSeZGdRMdc7BBPlkMPfg9nvSAqWTfxy','2023-05-23 11:09:57'),
	 ('567840923179','Jessica','Lee','jessicalee@example.com','$2a$10$i6RPPHA95dB.F8KkiNQ1C.DSeZGdRMdc7BBPlkMPfg9nvSAqWTfxy','2023-05-23 11:09:57'),
	 ('098723456736','Matthew','Lewis','matthewlewis@example.com','$2a$10$i6RPPHA95dB.F8KkiNQ1C.DSeZGdRMdc7BBPlkMPfg9nvSAqWTfxy','2023-05-23 11:09:57'),
	 ('785643298764','Lauren','Miller','laurenmiller@example.com','$2a$10$i6RPPHA95dB.F8KkiNQ1C.DSeZGdRMdc7BBPlkMPfg9nvSAqWTfxy','2023-05-23 11:09:57'),
	 ('450976834215','Andrew','Wilson','andrewwilson@example.com','$2a$10$i6RPPHA95dB.F8KkiNQ1C.DSeZGdRMdc7BBPlkMPfg9nvSAqWTfxy','2023-05-23 11:09:57'),
	 ('564328978654','Olivia','Taylor','oliviataylor@example.com','$2a$10$i6RPPHA95dB.F8KkiNQ1C.DSeZGdRMdc7BBPlkMPfg9nvSAqWTfxy','2023-05-23 11:09:57'),
	 ('785634291540','Ryan','Anderson','ryananderson@example.com','$2a$10$i6RPPHA95dB.F8KkiNQ1C.DSeZGdRMdc7BBPlkMPfg9nvSAqWTfxy','2023-05-23 11:09:57'),
	 ('564327895647','Amanda','Thomas','amandathomas@example.com','$2a$10$i6RPPHA95dB.F8KkiNQ1C.DSeZGdRMdc7BBPlkMPfg9nvSAqWTfxy','2023-05-23 11:09:57'),
	 ('786534297635','James','White','jameswhite@example.com','$2a$10$i6RPPHA95dB.F8KkiNQ1C.DSeZGdRMdc7BBPlkMPfg9nvSAqWTfxy','2023-05-23 11:09:57'),
	 ('786543290876','Elizabeth','Martin','elizabethmartin@example.com','$2a$10$i6RPPHA95dB.F8KkiNQ1C.DSeZGdRMdc7BBPlkMPfg9nvSAqWTfxy','2023-05-23 11:09:57'),
	 ('123409875634','Daniel','Thompson','danielthompson@example.com','$2a$10$i6RPPHA95dB.F8KkiNQ1C.DSeZGdRMdc7BBPlkMPfg9nvSAqWTfxy','2023-05-23 11:09:57'),
	 ('567432098769','Sophia','Davis','sophiadavis@example.com','$2a$10$i6RPPHA95dB.F8KkiNQ1C.DSeZGdRMdc7BBPlkMPfg9nvSAqWTfxy','2023-05-23 11:09:57'),
	 ('786543209876','Jacob','Wilson','jacobwilson@example.com','$2a$10$i6RPPHA95dB.F8KkiNQ1C.DSeZGdRMdc7BBPlkMPfg9nvSAqWTfxy','2023-05-23 11:09:57'),
	 ('564327891532','Mia','Moore','miamoore@example.com','$2a$10$i6RPPHA95dB.F8KkiNQ1C.DSeZGdRMdc7BBPlkMPfg9nvSAqWTfxy','2023-05-23 11:09:57');

	 -- пароль для пользователей 123456