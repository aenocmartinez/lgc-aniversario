DROP DATABASE IF EXISTS lgc_aniversario_db;
CREATE DATABASE lgc_aniversario_db;

USE lgc_aniversario_db;


-- DROP TABLE IF EXISTS formularios;
-- CREATE TABLE formularios (
-- 	id 					BIGINT AUTO_INCREMENT PRIMARY KEY,
-- 	nombre 				VARCHAR(255) NOT NULL,
-- 	documento 			VARCHAR(40) NOT NULL,
-- 	email 				VARCHAR(255) NOT NULL,
-- 	telefono 			VARCHAR(50) NOT NULL,
-- 	comprobante_pago 	VARCHAR(700) NOT NULL,
-- 	ciudad 				VARCHAR(100) DEFAULT '',
-- 	iglesia 			VARCHAR(255) DEFAULT '',
-- 	habeas_data 		BOOL DEFAULT true,
-- 	estado 				ENUM('Aprobada', 'PreAprobada', 'Anulada') DEFAULT 'PreAprobada',
-- 	asistencia			ENUM('Virtual', 'Presencial') DEFAULT 'Presencial',
-- 	fecha_registro 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP	
-- );


-- ALTER TABLE formularios ADD COLUMN medio_pago ENUM('Transferencia', 'Efectivo') DEFAULT 'Transferencia' AFTER habeas_data;

DROP TABLE IF EXISTS participantes;
DROP TABLE IF EXISTS inscripciones;
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

CREATE TABLE inscripciones (
    id 					BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    forma_pago 			ENUM('efectivo', 'transaccion', 'gratuito') NOT NULL,
    monto_pagado_cop 	INT NOT NULL,
    monto_pagado_usd 	INT NOT NULL,
    soporte_pago_url 	VARCHAR(255) NOT NULL,
    estado 				ENUM('Aprobada', 'PreAprobada', 'Rechazada') DEFAULT 'PreAprobada',
    created_at 			TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE inscripciones 
MODIFY COLUMN forma_pago ENUM('efectivo', 'transaccion', 'gratuito') NOT NULL;




CREATE TABLE participantes (
    id 						BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    inscripcion_id 			BIGINT UNSIGNED NOT NULL,
    nombre_completo 		VARCHAR(150) NOT NULL,
    numero_documento 		VARCHAR(30) NOT NULL,
    correo_electronico 		VARCHAR(100),
    telefono 				VARCHAR(20),
    modalidad 				ENUM('presencial', 'virtual') NOT NULL,
    dias_asistencia 		ENUM('viernes_y_domingo', 'sabado') DEFAULT NULL,
    iglesia 				VARCHAR(150) DEFAULT NULL,
    ciudad 					VARCHAR(100) DEFAULT NULL,
    autorizacion_datos 		BOOLEAN NOT NULL DEFAULT TRUE,
    created_at 				TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (inscripcion_id) REFERENCES inscripciones(id) ON DELETE CASCADE
);
