DROP DATABASE IF EXISTS lgc_aniversario_db;
CREATE DATABASE lgc_aniversario_db;

USE lgc_aniversario_db;


DROP TABLE IF EXISTS formularios;
CREATE TABLE formularios (
	id 					BIGINT AUTO_INCREMENT PRIMARY KEY,
	nombre 				VARCHAR(255) NOT NULL,
	documento 			VARCHAR(40) NOT NULL,
	email 				VARCHAR(255) NOT NULL,
	telefono 			VARCHAR(50) NOT NULL,
	comprobante_pago 	VARCHAR(700) NOT NULL,
	ciudad 				VARCHAR(100) DEFAULT '',
	iglesia 			VARCHAR(255) DEFAULT '',
	habeas_data 		BOOL DEFAULT true,
	estado 				ENUM('Pendiente', 'Validado') DEFAULT 'Pendiente',
	asistencia			ENUM('Virtual', 'Presencial') DEFAULT 'Presencial',
	fecha_registro 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP	


	-- UNIQUE(documento) 
);


DROP TABLE IF EXISTS usuarios;
CREATE TABLE usuarios (
	id 				BIGINT AUTO_INCREMENT PRIMARY KEY,
	nombre 			VARCHAR(255) NOT NULL,
	email 			VARCHAR(255) NOT NULL,
	password 		VARCHAR(255) NOT NULL,
	session_token 	VARCHAR(500),

	UNIQUE(email),
	UNIQUE(session_token)
);


INSERT INTO usuarios (nombre, email, password) VALUES ('Abimelec Enoc Martinez Robles', 'aenoc.martinez@gmail.com', '$2a$10$DGU1xtIJgFcsnsGSWY.7Ren2jPar1lE7Hlgoe7scnMlNbV1UFeCne');