USE posta_ppc;

CREATE TABLE Paciente(
	ID_Paciente INT PRIMARY KEY auto_increment,
	Nombre_Paciente NVARCHAR(128),
	Apellido_Paciente NVARCHAR(128),
	Nick_Paciente NVARCHAR(100),
	Clave_Paciente NVARCHAR(100),
	Foto_Paciente LONGBLOB,
	Nacionalidad_Paciente NVARCHAR(128),
	DATE_Nac_Paciente NVARCHAR(255),
	Email_Paciente NVARCHAR(100),
	Telefono_Paciente NVARCHAR(45),
	Celular_Paciente NVARCHAR(45),
	Cedula_Paciente NVARCHAR(45),
	Desc_Paciente TEXT,
	Archivo_Paciente NVARCHAR(128),
	Estado_Paciente NVARCHAR(45),
	Tag_Paciente NVARCHAR(128)
);

CREATE TABLE Perfil(
	ID_Perfil INT PRIMARY KEY auto_increment,
	Nombre_Perfil NVARCHAR(100),
	Estado_Perfil INT,
	Atributos_Perfil NVARCHAR(100)
);

CREATE TABLE Area(
	ID_Area INT PRIMARY KEY auto_increment,
	Nombre_Area NVARCHAR(128),
	Ubicacion_Area NVARCHAR(128),
	Localidades_Area TEXT,
	Ciudad_Area NVARCHAR(128),
	Pais_Area NVARCHAR(128),
	Capacidad_Area INT,
	Tag_Area NVARCHAR(128),
	Desc_Area TEXT
);

CREATE TABLE Usuario(
	ID_Usuario INT PRIMARY KEY auto_increment,
	ID_Area INT,
	ID_Perfil INT,
	Nombre_Usuario NVARCHAR(128),
	Apellido_Usuairo NVARCHAR(100),
	Nick_Usuario NVARCHAR(128),
	Clave_Usuario NVARCHAR(254),
	Mail_Usuario NVARCHAR(128),
	Especialidad_Usuario NVARCHAR(100),
	Estado_Usuario SMALLINT,
	CONSTRAINT fk_Area FOREIGN KEY (ID_Area) REFERENCES Area (ID_Area),
	CONSTRAINT fk_Perfil FOREIGN KEY (ID_Perfil) REFERENCES Perfil (ID_Perfil)
);

CREATE TABLE Cola(
	ID_Cola INT PRIMARY KEY auto_increment,
	ID_Paciente INT,
	ID_Usuario INT,
	DATE_Inicio_Cola DATE,
	Hora_Inicio_Cola TIME,
	DATE_Creación_Cola DATE,
	Estado_Cola NVARCHAR(45),
	CONSTRAINT fk_Paciente FOREIGN KEY (ID_Paciente) REFERENCES Paciente (ID_Paciente),
	CONSTRAINT fk_Usuario FOREIGN KEY (ID_Usuario) REFERENCES Usuario (ID_Usuario)
);

CREATE TABLE Cita(
	ID_Cita INT PRIMARY KEY auto_increment,
	ID_Usuario INT,
	ID_Paciente INT,
	Titulo_Cita NVARCHAR(100),
	DATE_Inicio_Cita DATE,
	DATE_Fin_Cita DATE,
	Hora_Inicio_Cita TIME,
	Hora_Fin_Cita TIME,
	Descripcion_Cita TEXT,
	Identificador_Cita NVARCHAR(120),
	Hora_Llegada_Paciente_Cita TIME,
	Hora_Salida_Paciente_cita TIME,
	Motivo_Anulacion_Cita TEXT,
	Estado_Cita NVARCHAR(45),
	CONSTRAINT fk_Paciente1 FOREIGN KEY (ID_Paciente) REFERENCES Paciente (ID_Paciente),
	CONSTRAINT fk_Usuario1 FOREIGN KEY (ID_Usuario) REFERENCES Usuario (ID_Usuario)
);