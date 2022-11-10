CREATE SCHEMA IF NOT EXISTS posta_ppc;

USE posta_ppc;

DROP TABLE IF EXISTS clients;

CREATE TABLE clients (
    id_cliente INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name_cliente VARCHAR(100) NOT NULL,
    last_name_cliente VARCHAR(100) NOT NULL,
    age_cliente INT NOT NULL
);

INSERT INTO
    clients (name_cliente, last_name_cliente, age_cliente)
    VALUES  ('arnold','chavez',23),
            ('kevin','burgos',21);

SELECT * FROM clients;