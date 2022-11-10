CREATE SCHEMA IF NOT EXISTS posta_ppc;

USE posta_ppc;

DROP TABLE IF EXISTS patients;

CREATE TABLE patients (
    id_patient INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name_patient VARCHAR(100) NOT NULL,
    last_name_patient VARCHAR(100) NOT NULL,
    age_patient INT NOT NULL
);

INSERT INTO
    patients (name_patient, last_name_patient, age_patient)
    VALUES  ('arnold','chavez',23),
            ('kevin','burgos',21);

SELECT * FROM patients;