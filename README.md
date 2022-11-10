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