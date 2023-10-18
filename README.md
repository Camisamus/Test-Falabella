# Falabella
 
1.- Primero se necesita crear la base de datos
Esta consiste de 2 tablas 

-- test.producto definition

CREATE TABLE `producto` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `SKU` varchar(15) NOT NULL,
  `NAME` varchar(50) NOT NULL,
  `BRAND` varchar(50) NOT NULL,
  `SIZE` varchar(50) NOT NULL,
  `PRICE` varchar(15) NOT NULL,
  `ESTADO` tinytext NOT NULL,
  PRIMARY KEY (`ID`),
  UNIQUE KEY `producto_un` (`SKU`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

-- test.imagen definition

CREATE TABLE `imagen` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `URL` text NOT NULL,
  `PRODUCTO` int(11) NOT NULL,
  `ESTADO` tinytext NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

2.- Segundo se requiere configurar el archivo "configuraciones.txt" el cual tiene 9 líneas
1   :18080      (Puerto donde levantaremos la app)
2           
3   root:       (Usuario de la BDD seguido de :)
4               (Clave de la BDD)
5   @tcp(       (Protocolo BDD)
6   localhost   (direccion de la BDD)
7   :3306       (Puerto de la BDD)
8   )/          (Cierre de cadena)
9   test        (Nombre de la BDD)

3.- Compilar El archivo main.go o debuggearlo según el método requerido

4.- Para hacer uso de los test unitarios Recordar modificar los valores necesarios para la ejecución
Según sea necesario, Por ejemplo cada vez que se crea un nuevo producto se requiere un SKU distinto

5.-Se inicia la página web como cualquier HTML



_____________________________________________________________________________________________________
La arquitectura usada es básicamente una separación de Handers, Funciones y Conexiones

Cada Handler responde a una URL disponible en la API 
Cada función resuelve una tarea en específico
Cada conexión Permite realizar un comando o conjunto de comandos a la base de datos con utilidad definida.

Esto se hace con este método para mantener un código ordenado y fácil de mantener.



Un detalle relevante es la decisión de eliminar los productos de forma lógica y no física 
(Marcarlos con una "D" en el estado) en un sistema más trabajado esto incluiría índices, y 
Facilitaría el seguimiento, la recuperación y trazabilidad de cambios y datos.






# Falabella
 
1.- First we need to create de database
This one requires 2 tables 

-- test.producto definition

CREATE TABLE `producto` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `SKU` varchar(15) NOT NULL,
  `NAME` varchar(50) NOT NULL,
  `BRAND` varchar(50) NOT NULL,
  `SIZE` varchar(50) NOT NULL,
  `PRICE` varchar(15) NOT NULL,
  `ESTADO` tinytext NOT NULL,
  PRIMARY KEY (`ID`),
  UNIQUE KEY `producto_un` (`SKU`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

-- test.imagen definition

CREATE TABLE `imagen` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `URL` text NOT NULL,
  `PRODUCTO` int(11) NOT NULL,
  `ESTADO` tinytext NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

2.- Second we must fill the "configuraciones.txt" file, it has 9 lines
1   :18080      (port to support app)
2           
3   root:       (DB user folowed by :)
4               (DB password)
5   @tcp(       (DB protocol)
6   localhost   (DB addres)
7   :3306       (DB port)
8   )/          (Closing of the connection Chain)
9   test        (DB name)

3.- Build or debug the main.go file according to requirement

4.- to use the unit tests remember to modify the required values for each execution if necessary, 
As example : each time you create a new product it is requires do use a brand new SKU

5.-The web page can be run as any HTML



_____________________________________________________________________________________________________
The architecture is divided between Handlers, functions and connections

Each Handler answers for a URL valid to de API
Each Function solves a specific task 
Each connection aloud an specific command or group of commands for the database 

This is for the code to be easy to read and maintain


An important detail is the decision to not physically delete products but use logic delete 
(By marking them with a D) in a finished system this with the usage of indexes would aloud 
easier traceability and recovery on their data.
