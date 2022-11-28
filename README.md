# PPC_G1_2022-2_Back-end  

**Nombre del Proyecto**: Back-end del proyecto   
SISTEMA DE GESTIÓN WEB DE CITAS VIRTUALES PARA LA ATENCIÓN EFICIENTE DE PACIENTES USANDO EL CONCEPTO DE CONCURRENCIA EN UNA CLÍNICA PRIVADA EN PERÚ   
**Curso**: Programación Paralela y Concurrente   
**Grupo**: Grupo 1, ciclo 2022-2

## Integrantes del Proyecto

- Arango Quispe, Esmeralda
- Chavez Burgos, Luiz Arnold
- Veramendi Santibañez, Dax Daco
- Villegas Panca, Fernando Miguel
- Zarate Villa, Jhennyfer Nayeli


# INSTRUCCIONES PARA CORRER EL PROYECTO   

## 1. Requerimientos de Software
- IDE: GoLand (recomendado para este proyecto)
- Go SDK 1.18.1
- MySQL 8.*

## 2. Editar el Archivo .env.example

1. Renombrar el archivo .env.example -> .env  
2. Agregar las cedenciales de su proyecto al formato establecido
```
# Server Config
PORT=

# Database Config
DB_USER=
DB_PASSWD=
DB_ADDR=
DB_PORT=
DB_NAME=
```

- Nota: en caso que la base de datos sea local:
  - BD_USER=root
  - BD_PASSWD=my_past_word
  - BD_ADDR=127.0.0.1
  - DB_PORT=3306
  - DB_NAME=schema_name

# 3. End Points

## 3.1. *Pacientes*   

### Obtener Todos los pacientes

```GET```
````http://127.0.0.1:4000/paciente````

### Obtener el Paciente segun {idPacient}

```GET```
````http://127.0.0.1:4000/paciente/{idPacient}````   

### Agregar Paciente

```POST```
````http://127.0.0.1:4000/paciente````   

### Editar Paciente segun {idPacient}

```PUT```
````http://127.0.0.1:4000/paciente/{idPacient}````

### Eliminar Paciente segun {idPacient}

```DELETE```
````http://127.0.0.1:4000/paciente/{idPacient}````   

## 3.2. *Areas*

### Obtener Todos los Areas

```GET```
````http://127.0.0.1:4000/area````

### Obtener el Area segun {idArea}

```GET```
````http://127.0.0.1:4000/area/{idArea}````

### Agregar Area

```POST```
````http://127.0.0.1:4000/area````

### Editar Area segun {idArea}

```PUT```
````http://127.0.0.1:4000/area/{idArea}````

### Eliminar Area segun {idArea}

```DELETE```
````http://127.0.0.1:4000/area/{idArea}````   

## 3.3. *Usuarios*

### Obtener Todos los Usuarios

```GET```
````http://127.0.0.1:4000/user````

### Obtener el Usuario segun {idUser}

```GET```
````http://127.0.0.1:4000/user/{idUser}````

### Agregar Usuario

```POST```
````http://127.0.0.1:4000/user````

### Editar Usuario segun {idUser}

```PUT```
````http://127.0.0.1:4000/user/{idUser}````

### Eliminar Usuario segun {idUser}

```DELETE```
````http://127.0.0.1:4000/user/{idUser}````   
