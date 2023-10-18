package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Producto
type Producto struct {
	ID       int      `json:"identificador"`
	SKU      string   `json:"SKU"`
	Name     string   `json:"name"`
	Brand    string   `json:"brand"`
	Size     string   `json:"size"`
	Price    string   `json:"price"`
	Imagenes []Imagen `json:"imagenes"`
	Estado   string   `json:"estado"`
}

//Imagen para json 1.0
type Imagen struct {
	ID       int    `json:"identificador"`
	Producto int    `json:"producto"`
	Url      string `json:"url"`
	Estado   string `json:"estado"`
}

var configuracion []string // variables necesarias
var bd *sql.DB             // coneccion a base de datos

func init() {
	actualizarConfiguracion()
	go actualizadorPasivo()
	conectadb()
}
func main() {
	enturador := mux.NewRouter().StrictSlash(false)
	enableCORS(enturador)
	enturador.HandleFunc("/Create", crearHandler).Methods("PUT")
	enturador.HandleFunc("/Read", leerHandler).Methods("GET")
	enturador.HandleFunc("/Ask", preguntarHandler).Methods("POST")
	enturador.HandleFunc("/Update", modificarHandler).Methods("POST")
	enturador.HandleFunc("/Delete", borrarHandler).Methods("DELETE")

	servidor := http.Server{
		Addr:           configuracion[0],
		Handler:        enturador,
		ReadTimeout:    3 * time.Second,
		WriteTimeout:   3 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Escuchando Puerto" + configuracion[0])

	servidor.ListenAndServe()

}
func enableCORS(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	}).Methods(http.MethodOptions)
	router.Use(middlewareCors)
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			// and call next handler!
			next.ServeHTTP(w, req)
		})
}
func actualizarConfiguracion() {
	contenido, err := ioutil.ReadFile("configuracion.txt")
	if err != nil {
		return
	}
	var lineas []string
	if runtime.GOOS == "windows" {
		lineas = strings.Split(strings.Replace(string(contenido), "\r\n", "\n", -1), "\n") //versio windows
	} else {
		lineas = strings.Split(string(contenido), "\n") //version linux
	}
	configuracion = lineas
}
func actualizadorPasivo() {
	for {
		log.Println("Reovando Variables Generales")
		actualizarConfiguracion()
		time.Sleep(time.Hour * 24)
	}
}

func conectadb() {
	bda, err := sql.Open("mysql", configuracion[2]+configuracion[3]+configuracion[4]+configuracion[5]+configuracion[6]+configuracion[7]+configuracion[8])
	if err != nil {
		log.Println(err.Error())
	}
	bd = bda
}

func crearHandler(w http.ResponseWriter, r *http.Request) {

	intput := Producto{}

	w.Header().Set("contenido-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&intput)
	if err != nil {
		log.Println("Errores : " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write(nil)
	}
	output, err := funcionCrear(intput)
	if err != nil {
		log.Println("Errores : " + err.Error())
		w.WriteHeader(http.StatusFailedDependency)
		respuesta, err := json.Marshal(`{"error":"` + err.Error() + `"}`)
		if err != nil {
			log.Println("Errores : " + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(nil)
			return
		}
		w.Write(respuesta)
		return
	}
	respuesta, err := json.Marshal(output)
	if err != nil {
		log.Println("Errores : " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(respuesta)
}
func leerHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("contenido-Type", "application/json")
	output, err := funcionConsultaMasiva()
	if err != nil {
		log.Println("Errores : " + err.Error())
		w.WriteHeader(http.StatusFailedDependency)
		respuesta, err := json.Marshal(`{"error":"` + err.Error() + `"}`)
		if err != nil {
			log.Println("Errores : " + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(nil)
			return
		}
		w.Write(respuesta)
		return
	}
	respuesta, err := json.Marshal(output)
	if err != nil {
		log.Println("Errores : " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)

		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(respuesta)
}
func preguntarHandler(w http.ResponseWriter, r *http.Request) {

	intput := Producto{}

	w.Header().Set("contenido-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&intput)
	if err != nil {
		log.Println("Errores : " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write(nil)
	}
	output, err := funcionConsultaUnitaria(intput)
	if err != nil {
		log.Println("Errores : " + err.Error())
		w.WriteHeader(http.StatusFailedDependency)
		respuesta, err := json.Marshal(`{"error":"` + err.Error() + `"}`)
		if err != nil {
			log.Println("Errores : " + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(nil)
			return
		}
		w.Write(respuesta)
		return
	}
	respuesta, err := json.Marshal(output)
	if err != nil {
		log.Println("Errores : " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(respuesta)
}
func modificarHandler(w http.ResponseWriter, r *http.Request) {

	intput := Producto{}

	w.Header().Set("contenido-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&intput)
	if err != nil {
		log.Println("Errores : " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write(nil)
	}
	output, err := funcionModificar(intput)
	if err != nil {
		log.Println("Errores : " + err.Error())
		w.WriteHeader(http.StatusFailedDependency)
		respuesta, err := json.Marshal(`{"error":"` + err.Error() + `"}`)
		if err != nil {
			log.Println("Errores : " + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(nil)
			return
		}
		w.Write(respuesta)
		return
	}
	respuesta, err := json.Marshal(output)
	if err != nil {
		log.Println("Errores : " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)

		w.Write(nil)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(respuesta)
}
func borrarHandler(w http.ResponseWriter, r *http.Request) {

	intput := Producto{}

	w.Header().Set("contenido-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&intput)
	if err != nil {
		log.Println("Errores : " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write(nil)
	}
	output, err := funcionEliminar(intput)
	if err != nil {
		log.Println("Errores : " + err.Error())
		w.WriteHeader(http.StatusFailedDependency)
		respuesta, err := json.Marshal(`{"error":"` + err.Error() + `"}`)
		if err != nil {
			log.Println("Errores : " + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(nil)
			return
		}
		w.Write(respuesta)
		return
	}
	respuesta, err := json.Marshal(output)
	if err != nil {
		log.Println("Errores : " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(respuesta)
}
func validarSKU(SKU string) bool {
	if len(SKU) < 6 {
		return false
	}
	aux1 := SKU[0:4]
	aux2 := SKU[4:]
	if aux1 != "FAL-" {
		return false
	}
	aux3, err := strconv.Atoi(aux2)
	if err != nil {
		return false
	}
	if aux3 < 1000000 || aux3 > 999999999 {
		return false
	}
	return true
}

func funcionCrear(prod Producto) (Producto, error) {
	outprod := Producto{}
	if !validarSKU(prod.SKU) {
		miError := errors.New("SKU no valido")
		return outprod, miError
	}
	if len(prod.Imagenes) <= 0 {
		miError := errors.New("se requiere al menos una imagen del producto")
		return outprod, miError
	}
	precio, err := strconv.Atoi(strings.Replace(prod.Price, ".", "", -1))
	if err != nil {
		miError := errors.New("el precio ingresado no tiene el formato correcto")
		return outprod, miError
	}
	if precio < 100 || precio > 9999999900 {
		miError := errors.New("el rango de precios es de 1.00 a 99999999.00")
		return outprod, miError
	}
	if prod.Size == "" {
		prod.Size = "N/A"
	}
	auxLenBrand := len(prod.Brand)
	if auxLenBrand > 50 || auxLenBrand < 3 {
		miError := errors.New("incluir una marca con nombre de largo entre 3 y 50")
		return outprod, miError
	}
	auxLenName := len(prod.Name)
	if auxLenName > 50 || auxLenName < 3 {
		miError := errors.New("incluir un nombre o descripcion de largo entre 3 y 50")
		return outprod, miError
	}
	aux0, err := consultaProducto("D", prod.SKU)
	if err != nil {
		return outprod, err
	}
	if aux0.Name != "" {
		prod.Estado = "A"
		err := modificarProducto(prod)
		if err != nil {
			return outprod, err
		}
		for _, i2 := range prod.Imagenes {
			i2.Producto = aux0.ID
			i2.Estado = "A"
			err = ingresoImagenes(i2)
			if err != nil {
				return outprod, err
			}
		}
		return outprod, nil
	}
	prod.Estado = "A"
	aux, err := ingresoProducto(prod)
	if err != nil {
		return outprod, err
	}
	for _, i2 := range prod.Imagenes {
		i2.Producto = aux
		i2.Estado = "A"
		err = ingresoImagenes(i2)
		if err != nil {
			return outprod, err
		}
	}
	outprod = prod
	outprod.ID = aux
	outprod.Estado = "Producto ingresado"
	return outprod, nil
}
func funcionConsultaMasiva() ([]Producto, error) {
	outprod := []Producto{}
	aux, err := consultaProductos("A")
	if err != nil {
		return outprod, err
	}
	aux2, err := consultaImagenes("A")
	if err != nil {
		return outprod, err
	}
	aux3 := map[int]Producto{}
	for _, j2 := range aux2 {
		for i, i2 := range aux {
			if i2.ID == j2.Producto {
				i2.Imagenes = append(aux3[i].Imagenes, j2)
				aux3[i] = i2
				break
			}
		}
	}
	for _, h2 := range aux3 {
		outprod = append(outprod, h2)
	}
	return outprod, nil
}
func funcionConsultaUnitaria(prod Producto) (Producto, error) {
	outprod := Producto{}
	if !validarSKU(prod.SKU) {
		miError := errors.New("SKU no valido")
		return outprod, miError
	}
	aux, err := consultaProducto("A", prod.SKU)
	if err != nil {
		return outprod, err
	}
	if aux.ID == 0 {
		miError := errors.New("SKU no encontrado")
		return outprod, miError
	}
	aux2, err := consultaImagenesProd("A", aux.ID)
	if err != nil {
		return outprod, err
	}
	aux.Imagenes = append(aux.Imagenes, aux2...)
	outprod = aux
	return outprod, nil
}

func funcionModificar(prod Producto) (Producto, error) {
	outprod := Producto{}
	if !validarSKU(prod.SKU) {
		miError := errors.New("SKU no valido")
		return outprod, miError
	}
	if len(prod.Imagenes) <= 0 {
		miError := errors.New("se requiere al menos una imagen del producto")
		return outprod, miError
	}
	precio, err := strconv.Atoi(strings.Replace(prod.Price, ".", "", -1))
	if err != nil {
		miError := errors.New("el precio ingresado no tiene el formato correcto")
		return outprod, miError
	}
	if precio < 100 || precio > 9999999900 {
		miError := errors.New("el rango de precios es de 1.00 a 99999999.00")
		return outprod, miError
	}
	if prod.Size == "" {
		prod.Size = "N/A"
	}
	auxLenBrand := len(prod.Brand)
	if auxLenBrand > 50 || auxLenBrand < 3 {
		miError := errors.New("incluir una marca con nombre de largo entre 3 y 50")
		return outprod, miError
	}
	auxLenName := len(prod.Name)
	if auxLenName > 50 || auxLenName < 3 {
		miError := errors.New("incluir un nombre o descripcion de largo entre 3 y 50")
		return outprod, miError
	}
	err = modificarProducto(prod)
	if err != nil {
		return outprod, err
	}
	aux, err := consultaProducto("A", prod.SKU)
	if err != nil {
		return outprod, err
	}
	err = eliminarImagenes(aux.ID)
	if err != nil {
		return outprod, err
	}
	for _, j2 := range prod.Imagenes {
		j2.Producto = aux.ID
		j2.Estado = "A"
		err = ingresoImagenes(j2)
		if err != nil {
			return outprod, err
		}
	}
	outprod.Estado = "Producto modificado"
	return outprod, nil
}
func funcionEliminar(prod Producto) (Producto, error) {
	outprod := Producto{}
	if !validarSKU(prod.SKU) {
		miError := errors.New("SKU no valido")
		return outprod, miError
	}
	prod.Estado = "D"
	err := modificarProducto(prod)
	if err != nil {
		return outprod, err
	}
	p, err := consultaProducto("D", prod.SKU)
	if err != nil {
		return outprod, err
	}
	eliminarImagenes(p.ID)
	outprod.Estado = "Producto Elimiinado"
	return outprod, nil
}

func consultaProducto(estado string, SKU string) (Productos Producto, err error) {
	producto := Producto{}
	db1, err := bd.Begin()
	if err != nil {
		log.Println(err.Error())
		return producto, err
	}
	query, err := db1.Prepare("SELECT * FROM PRODUCTO WHERE ESTADO = ? AND SKU = ?")
	if err != nil {
		log.Println(err.Error())
		db1.Rollback()
		return producto, err
	}
	tab1, err := query.Query(estado, SKU) //db1.Query(, )
	if err != nil {
		log.Println(err.Error())
		db1.Rollback()
		return producto, err
	}
	defer tab1.Close()
	for tab1.Next() {
		err = tab1.Scan(&producto.ID, &producto.SKU, &producto.Name, &producto.Brand, &producto.Size, &producto.Price, &producto.Estado)
		if err != nil {
			log.Println(err.Error())
			db1.Rollback()
			return producto, err
		}
		producto.Price = formatearPrecio(producto.Price)
	}
	err = db1.Commit()

	if err != nil {
		log.Println(err.Error())
		return producto, err
	}
	return producto, nil

}
func consultaImagenesProd(estado string, producto int) (Productos []Imagen, err error) {
	imagenes := []Imagen{}
	db1, err := bd.Begin()
	if err != nil {
		log.Println(err.Error())
		return imagenes, err
	}
	query, err := db1.Prepare("SELECT * FROM IMAGEN WHERE ESTADO = ? AND PRODUCTO = ?")
	if err != nil {
		log.Println(err.Error())
		db1.Rollback()
		return imagenes, err
	}
	tab1, err := query.Query(estado, producto) //db1.Query(, )
	if err != nil {
		log.Println(err.Error())
		db1.Rollback()
		return imagenes, err
	}
	defer tab1.Close()
	for tab1.Next() {
		imagen := Imagen{}
		err = tab1.Scan(&imagen.ID, &imagen.Url, &imagen.Producto, &imagen.Estado)
		if err != nil {
			log.Println(err.Error())
			db1.Rollback()
			return imagenes, err
		}
		imagenes = append(imagenes, imagen)
	}
	err = db1.Commit()

	if err != nil {
		log.Println(err.Error())
		return imagenes, err
	}
	return imagenes, nil

}

func consultaProductos(estado string) (Productos []Producto, err error) {
	productos := []Producto{}
	db1, err := bd.Begin()
	if err != nil {
		log.Println(err.Error())
		return productos, err
	}
	query, err := db1.Prepare("SELECT * FROM PRODUCTO WHERE ESTADO = ? ")
	if err != nil {
		log.Println(err.Error())
		db1.Rollback()
		return productos, err
	}
	tab1, err := query.Query(estado) //db1.Query(, )
	if err != nil {
		log.Println(err.Error())
		db1.Rollback()
		return productos, err
	}
	defer tab1.Close()
	for tab1.Next() {
		producto := Producto{}
		err = tab1.Scan(&producto.ID, &producto.SKU, &producto.Name, &producto.Brand, &producto.Size, &producto.Price, &producto.Estado)
		if err != nil {
			log.Println(err.Error())
			db1.Rollback()
			return productos, err
		}
		producto.Price = formatearPrecio(producto.Price)
		productos = append(productos, producto)
	}
	err = db1.Commit()

	if err != nil {
		log.Println(err.Error())
		return productos, err
	}
	return productos, nil

}
func consultaImagenes(estado string) (Productos []Imagen, err error) {
	imagenes := []Imagen{}
	db1, err := bd.Begin()
	if err != nil {
		log.Println(err.Error())
		return imagenes, err
	}
	query, err := db1.Prepare("SELECT * FROM IMAGEN WHERE ESTADO = ?")
	if err != nil {
		log.Println(err.Error())
		db1.Rollback()
		return imagenes, err
	}
	tab1, err := query.Query(estado) //db1.Query(, )
	if err != nil {
		log.Println(err.Error())
		db1.Rollback()
		return imagenes, err
	}
	defer tab1.Close()
	for tab1.Next() {
		imagen := Imagen{}
		err = tab1.Scan(&imagen.ID, &imagen.Url, &imagen.Producto, &imagen.Estado)
		if err != nil {
			log.Println(err.Error())
			db1.Rollback()
			return imagenes, err
		}
		imagenes = append(imagenes, imagen)
	}
	err = db1.Commit()

	if err != nil {
		log.Println(err.Error())
		return imagenes, err
	}
	return imagenes, nil

}

func ingresoProducto(producto Producto) (id int, err error) {
	db1, err := bd.Begin()
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	_, err = db1.Query("INSERT INTO PRODUCTO (SKU,NAME,BRAND,SIZE,PRICE,ESTADO) VALUES (?,?,?,?,?,?)", producto.SKU, producto.Name, producto.Brand, producto.Size, producto.Price, producto.Estado)
	if err != nil {
		log.Println(err.Error())
		db1.Rollback()
		return 0, err
	}
	aux := 0
	tab, err := db1.Query("select LAST_INSERT_ID()")
	if err != nil {
		log.Println(err.Error())
		db1.Rollback()
		return 0, err
	}
	defer tab.Close()
	for tab.Next() {
		err = tab.Scan(&aux)
		if err != nil {
			log.Println(err.Error())
			db1.Rollback()
			return 0, err
		}
	}
	err = db1.Commit()
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}
	return aux, err

}
func ingresoImagenes(imagen Imagen) (err error) {
	db1, err := bd.Begin()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	_, err = db1.Query("INSERT INTO IMAGEN (URL,PRODUCTO,ESTADO) VALUES (?,?,?)", imagen.Url, imagen.Producto, imagen.Estado)
	if err != nil {
		log.Println(err.Error())
		db1.Rollback()
		return err
	}
	err = db1.Commit()

	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil

}

func modificarProducto(producto Producto) (err error) {
	db1, err := bd.Begin()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	query, err := db1.Prepare("SELECT * FROM PRODUCTO WHERE SKU = ? ")
	if err != nil {
		log.Println(err.Error())
		db1.Rollback()
		return err
	}
	tab1, err := query.Query(producto.SKU) //db1.Query(, )
	if err != nil {
		log.Println(err.Error())
		db1.Rollback()
		return err
	}
	productoBDD := Producto{}
	for tab1.Next() {
		err = tab1.Scan(&productoBDD.ID, &productoBDD.SKU, &productoBDD.Name, &productoBDD.Brand, &productoBDD.Size, &productoBDD.Price, &productoBDD.Estado)
		if err != nil {
			log.Println(err.Error())
			db1.Rollback()
			return err
		}
	}
	defer tab1.Close()

	if producto.SKU != "" {
		productoBDD.SKU = producto.SKU
	}
	if producto.Name != "" {
		productoBDD.Name = producto.Name
	}
	if producto.Brand != "" {
		productoBDD.Brand = producto.Brand
	}
	if producto.Size != "" {
		productoBDD.Size = producto.Size
	}
	if producto.Price != "" {
		productoBDD.Price = producto.Price
	}
	if producto.Estado != "" {
		productoBDD.Estado = producto.Estado
	}

	_, err = db1.Query("UPDATE PRODUCTO SET SKU = ?,NAME = ?, BRAND = ?, SIZE = ?,PRICE = ?, ESTADO = ? WHERE ID = ?", productoBDD.SKU, productoBDD.Name, productoBDD.Brand, productoBDD.Size, productoBDD.Price, productoBDD.Estado, productoBDD.ID)
	if err != nil {
		log.Println(err.Error())
		db1.Rollback()
		return err
	}
	err = db1.Commit()

	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil

}
func eliminarImagenes(imagen int) (err error) {
	db1, err := bd.Begin()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	_, err = db1.Query("DELETE FROM IMAGEN WHERE PRODUCTO = ?", imagen)
	if err != nil {
		log.Println(err.Error())
		db1.Rollback()
		return err
	}
	err = db1.Commit()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
func formatearPrecio(precio string) string {
	auxLenPrecio := len(precio)
	return precio[:auxLenPrecio-2] + "." + precio[auxLenPrecio-2:]
}
