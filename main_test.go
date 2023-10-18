package main

import (
	"strconv"
	"testing"
)

//Tests Base
func TestValidarSKU(t *testing.T) {
	got := validarSKU("FAL-881952288")

	if !got {
		t.Errorf("got %q, wanted %q", "false", "true")
		return
	}
}
func TestCrear(t *testing.T) {
	auxP := Producto{
		SKU:      "FAL-881952288",
		Name:     "Test",
		Brand:    "Test",
		Size:     "XXL",
		Price:    "5999.99",
		Imagenes: []Imagen{},
		Estado:   "A",
	}
	auxI := Imagen{
		Producto: 0,
		Url:      "TEST.test.cl",
		Estado:   "A",
	}
	auxP.Imagenes = append(auxP.Imagenes, auxI)
	got, err := funcionCrear(auxP)
	if got.ID == 0 || err != nil {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(got.ID), "nil, >0")
		return
	}
}
func TestConsultaMasiva(t *testing.T) {
	got, err := funcionConsultaMasiva()
	if len(got) == 0 || err != nil {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(len(got)), "nil, >0")
		return
	}
}
func TestConsultaUnitaria(t *testing.T) {
	auxP := Producto{
		SKU: "FAL-881952288",
	}
	got, err := funcionConsultaUnitaria(auxP)
	if got.ID == 0 || err != nil {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(got.ID), "nil, >0")
		return
	}
}

func TestModificarProd(t *testing.T) {
	auxP := Producto{
		SKU:      "FAL-881952288",
		Name:     "Test",
		Brand:    "Test",
		Size:     "XXL",
		Price:    "5999.99",
		Imagenes: []Imagen{},
		Estado:   "A",
	}
	auxI := Imagen{
		Producto: 0,
		Url:      "TEST.test.cl",
		Estado:   "A",
	}
	auxP.Imagenes = append(auxP.Imagenes, auxI)
	got, err := funcionModificar(auxP)
	if got.Estado != "Producto modificado" || err != nil {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(got.ID), "nil, >0")
		return
	}
}
func TestEliminarProd(t *testing.T) {
	got, err := funcionEliminar(Producto{
		SKU: "FAL-881952288",
	})
	if got.Estado != "Producto Elimiinado" || err != nil {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(got.ID), "nil, >0")
		return
	}
}
func TestConsultaProducto(t *testing.T) {
	got, err := consultaProducto("A", "FAL-881952286")
	if got.ID == 0 || err != nil {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(got.ID), "nil, >0")
		return
	}
}
func TestConsultarImagenesProd(t *testing.T) {
	got, err := consultaImagenesProd("A", 2)
	if len(got) == 0 || err != nil {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(len(got)), "nil, >0")
		return
	}
}
func TestConsultarProductoss(t *testing.T) {
	got, err := consultaProductos("A")
	if len(got) == 0 || err != nil {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(len(got)), "nil, >0")
		return
	}
}
func TestConsultarImagenes(t *testing.T) {
	got, err := consultaImagenes("A")
	if len(got) == 0 || err != nil {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(len(got)), "nil, >0")
		return
	}
}
func TestIngresoProducto(t *testing.T) {
	aux, err := ingresoProducto(Producto{
		SKU:    "FAL-881952286",
		Name:   "TestProduct",
		Brand:  "tt2",
		Size:   "XXL",
		Price:  "29990",
		Estado: "A",
	})
	if err != nil || aux == 0 {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(aux), "nil, >0")
		return
	}
}

func TestIntesoImagenes(t *testing.T) {
	err := ingresoImagenes(Imagen{
		Producto: 5,
		Url:      "http://test.test.cl",
		Estado:   "A",
	})
	if err != nil {
		t.Errorf("got %q, wanted %q", err.Error(), "nil")
		return
	}
}
func TestModificarProducto(t *testing.T) {
	err := modificarProducto(Producto{
		ID:       0,
		SKU:      "FAL-881952284",
		Name:     "",
		Brand:    "",
		Size:     "",
		Price:    "",
		Imagenes: []Imagen{},
		Estado:   "A",
	})
	if err != nil {
		t.Errorf("got %q, wanted %q", err.Error(), "nil")
		return
	}
}

func TestEliminarImagenes(t *testing.T) {
	err := eliminarImagenes(8)

	if err != nil {
		t.Errorf("got %q, wanted %q", err.Error(), "nil")
		return
	}
}
func TestFormatearPrecio(t *testing.T) {
	got := formatearPrecio("xxxx")
	want := "xx.xx"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
		return
	}
}

//Test Medios

func TestValidarSKU_2(t *testing.T) {
	//se espera que falle la validacion
	got := validarSKU("FAL-88195228k")
	if got {
		t.Errorf("got %q, wanted %q", "false", "true")
		return
	}
}

func TestCrear_2(t *testing.T) {
	//se espera que falle al intentar ingresar 2 veces el mismo SKU
	auxP := Producto{
		SKU:      "FAL-881952291",
		Name:     "Test",
		Brand:    "Test",
		Size:     "XXL",
		Price:    "5999.99",
		Imagenes: []Imagen{},
		Estado:   "A",
	}
	auxI := Imagen{
		Producto: 0,
		Url:      "TEST.test.cl",
		Estado:   "A",
	}
	auxP.Imagenes = append(auxP.Imagenes, auxI)
	got, err := funcionCrear(auxP)
	if got.ID == 0 || err != nil {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(got.ID), "nil, >0")
		return
	}
	got, err = funcionCrear(auxP)
	if got.ID != 0 || err.Error() != "Error 1062: Duplicate entry '"+auxP.SKU+"' for key 'producto_un'" {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(got.ID), "Error 1062: Duplicate entry '"+auxP.SKU+"' for key 'producto_un', 0")
		return
	}
}
func TestCrear_3(t *testing.T) {
	//se espera que falle al usar un SKU invalido
	auxP := Producto{
		SKU:      "FALl881952291",
		Name:     "Test",
		Brand:    "Test",
		Size:     "XXL",
		Price:    "5999.99",
		Imagenes: []Imagen{},
		Estado:   "A",
	}
	auxI := Imagen{
		Producto: 0,
		Url:      "TEST.test.cl",
		Estado:   "A",
	}
	auxP.Imagenes = append(auxP.Imagenes, auxI)
	got, err := funcionCrear(auxP)
	if got.ID != 0 || err.Error() != "SKU no valido" {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(got.ID), "nil, 0")
		return
	}
}
func TestCrear_4(t *testing.T) {
	//se espera que falle al intentar ingresar un producto sin imagen
	auxP := Producto{
		SKU:      "FAL-881952291",
		Name:     "Test",
		Brand:    "Test",
		Size:     "XXL",
		Price:    "5999.99",
		Imagenes: []Imagen{},
		Estado:   "A",
	}
	got, err := funcionCrear(auxP)
	if got.ID != 0 || err.Error() != "se requiere al menos una imagen del producto" {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(got.ID), "nil, 0")
		return
	}
}
func TestCrear_5(t *testing.T) {
	//se espera que falle al usar precio en formato incorrecto
	auxP := Producto{
		SKU:      "FAL-881952291",
		Name:     "Test",
		Brand:    "Test",
		Size:     "XXL",
		Price:    "5999,99",
		Imagenes: []Imagen{},
		Estado:   "A",
	}
	auxI := Imagen{
		Producto: 0,
		Url:      "TEST.test.cl",
		Estado:   "A",
	}
	auxP.Imagenes = append(auxP.Imagenes, auxI)
	got, err := funcionCrear(auxP)
	if got.ID != 0 || err.Error() != "el precio ingresado no tiene el formato correcto" {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(got.ID), "nil, 0")
		return
	}
}
func TestCrear_6(t *testing.T) {
	//se espera que falle al usar un precio muy bajo
	auxP := Producto{
		SKU:      "FAL-881952291",
		Name:     "Test",
		Brand:    "Test",
		Size:     "XXL",
		Price:    "99",
		Imagenes: []Imagen{},
		Estado:   "A",
	}
	auxI := Imagen{
		Producto: 0,
		Url:      "TEST.test.cl",
		Estado:   "A",
	}
	auxP.Imagenes = append(auxP.Imagenes, auxI)
	got, err := funcionCrear(auxP)
	if got.ID != 0 || err.Error() != "el rango de precios es de 1.00 a 99999999.00" {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(got.ID), "nil, 0")
		return
	}
}
func TestCrear_7(t *testing.T) {
	//se espera que falle al usar un precio muy alto
	auxP := Producto{
		SKU:      "FAL-881952291",
		Name:     "Test",
		Brand:    "Test",
		Size:     "XXL",
		Price:    "10000000000",
		Imagenes: []Imagen{},
		Estado:   "A",
	}
	auxI := Imagen{
		Producto: 0,
		Url:      "TEST.test.cl",
		Estado:   "A",
	}
	auxP.Imagenes = append(auxP.Imagenes, auxI)
	got, err := funcionCrear(auxP)
	if got.ID != 0 || err.Error() != "el rango de precios es de 1.00 a 99999999.00" {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(got.ID), "nil, 0")
		return
	}
}
func TestCrear_8(t *testing.T) {
	//se espera que funcione pero al recuperar la talla aparesca como "N/A"
	auxP := Producto{
		SKU:      "FAL-881952293",
		Name:     "Test",
		Brand:    "Test",
		Size:     "",
		Price:    "1000000",
		Imagenes: []Imagen{},
		Estado:   "A",
	}
	auxI := Imagen{
		Producto: 0,
		Url:      "TEST.test.cl",
		Estado:   "A",
	}
	auxP.Imagenes = append(auxP.Imagenes, auxI)
	got, err := funcionCrear(auxP)
	if got.ID == 0 || err != nil {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(got.ID), "nil, 0")
		return
	}
	got, err = consultaProducto("A", auxP.SKU)
	if got.ID == 0 || err != nil || got.Size != "N/A" {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(got.ID), "nil, 0")
		return
	}
}

func TestCrear_9(t *testing.T) {
	//se espera que falle al usar marca con nombre muy corto
	auxP := Producto{
		SKU:      "FAL-881952291",
		Name:     "Test",
		Brand:    "Te",
		Size:     "XXL",
		Price:    "10000000",
		Imagenes: []Imagen{},
		Estado:   "A",
	}
	auxI := Imagen{
		Producto: 0,
		Url:      "TEST.test.cl",
		Estado:   "A",
	}
	auxP.Imagenes = append(auxP.Imagenes, auxI)
	got, err := funcionCrear(auxP)
	if got.ID != 0 || err.Error() != "incluir una marca con nombre de largo entre 3 y 50" {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(got.ID), "nil, 0")
		return
	}
}
func TestCrear_10(t *testing.T) {
	//se espera que falle al usar marca con nombre muy largo
	auxP := Producto{
		SKU:      "FAL-881952291",
		Name:     "Test",
		Brand:    "Test-Test-Test-Test-Test-Test-Test-Test-Test-Test-Test",
		Size:     "XXL",
		Price:    "10000000",
		Imagenes: []Imagen{},
		Estado:   "A",
	}
	auxI := Imagen{
		Producto: 0,
		Url:      "TEST.test.cl",
		Estado:   "A",
	}
	auxP.Imagenes = append(auxP.Imagenes, auxI)
	got, err := funcionCrear(auxP)
	if got.ID != 0 || err.Error() != "incluir una marca con nombre de largo entre 3 y 50" {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(got.ID), "nil, 0")
		return
	}
}
func TestCrear_11(t *testing.T) {
	//se espera que falle al usar nombre muy corto
	auxP := Producto{
		SKU:      "FAL-881952291",
		Name:     "Te",
		Brand:    "Test",
		Size:     "XXL",
		Price:    "10000000",
		Imagenes: []Imagen{},
		Estado:   "A",
	}
	auxI := Imagen{
		Producto: 0,
		Url:      "TEST.test.cl",
		Estado:   "A",
	}
	auxP.Imagenes = append(auxP.Imagenes, auxI)
	got, err := funcionCrear(auxP)
	if got.ID != 0 || err.Error() != "incluir un nombre o descripcion de largo entre 3 y 50" {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(got.ID), "incluir un nombre o descripcion de largo entre 3 y 50, 0")
		return
	}
}
func TestCrear_12(t *testing.T) {
	//se espera que falle al usar nombre muy largo
	auxP := Producto{
		SKU:      "FAL-881952291",
		Name:     "Test-Test-Test-Test-Test-Test-Test-Test-Test-Test-Test",
		Brand:    "Tes",
		Size:     "XXL",
		Price:    "10000000",
		Imagenes: []Imagen{},
		Estado:   "A",
	}
	auxI := Imagen{
		Producto: 0,
		Url:      "TEST.test.cl",
		Estado:   "A",
	}
	auxP.Imagenes = append(auxP.Imagenes, auxI)
	got, err := funcionCrear(auxP)
	if got.ID != 0 || err.Error() != "incluir un nombre o descripcion de largo entre 3 y 50" {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(got.ID), "incluir un nombre o descripcion de largo entre 3 y 50, 0")
		return
	}
}

func TestModificar_2(t *testing.T) {
	//se espera que funcione al hacer varias modificaciones
	auxP := Producto{
		SKU:      "FAL-881952291",
		Name:     "Test",
		Brand:    "Test",
		Size:     "XXL",
		Price:    "5999.99",
		Imagenes: []Imagen{},
		Estado:   "A",
	}
	auxI := Imagen{
		Producto: 0,
		Url:      "TEST.test.cl",
		Estado:   "A",
	}
	auxP.Imagenes = append(auxP.Imagenes, auxI)
	got, err := funcionModificar(auxP)
	if got.Estado != "Producto modificado" || err != nil {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(got.ID), "nil, >0")
		return
	}
	got, err = funcionModificar(auxP)
	if got.Estado != "Producto modificado" || err != nil {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(got.ID), "Error 1062: Duplicate entry '"+auxP.SKU+"' for key 'producto_un', 0")
		return
	}
}
func TestModificar_3(t *testing.T) {
	//se espera que falle al usar un SKU invalido
	auxP := Producto{
		SKU:      "FALl881952291",
		Name:     "Test",
		Brand:    "Test",
		Size:     "XXL",
		Price:    "5999.99",
		Imagenes: []Imagen{},
		Estado:   "A",
	}
	auxI := Imagen{
		Producto: 0,
		Url:      "TEST.test.cl",
		Estado:   "A",
	}
	auxP.Imagenes = append(auxP.Imagenes, auxI)
	got, err := funcionModificar(auxP)
	if got.Estado == "Producto modificado" || err.Error() != "SKU no valido" {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(got.ID), "nil, 0")
		return
	}
}
func TestModificar_4(t *testing.T) {
	//se espera que falle al intentar ingresar un producto sin imagen
	auxP := Producto{
		SKU:      "FAL-881952291",
		Name:     "Test",
		Brand:    "Test",
		Size:     "XXL",
		Price:    "5999.99",
		Imagenes: []Imagen{},
		Estado:   "A",
	}
	got, err := funcionModificar(auxP)
	if got.Estado == "Producto modificado" || err.Error() != "se requiere al menos una imagen del producto" {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(got.ID), "nil, 0")
		return
	}
}
func TestModificar_5(t *testing.T) {
	//se espera que falle al usar precio en formato incorrecto
	auxP := Producto{
		SKU:      "FAL-881952291",
		Name:     "Test",
		Brand:    "Test",
		Size:     "XXL",
		Price:    "5999,99",
		Imagenes: []Imagen{},
		Estado:   "A",
	}
	auxI := Imagen{
		Producto: 0,
		Url:      "TEST.test.cl",
		Estado:   "A",
	}
	auxP.Imagenes = append(auxP.Imagenes, auxI)
	got, err := funcionModificar(auxP)
	if got.Estado == "Producto modificado" || err.Error() != "el precio ingresado no tiene el formato correcto" {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(got.ID), "nil, 0")
		return
	}
}
func TestModificar_6(t *testing.T) {
	//se espera que falle al usar un precio muy bajo
	auxP := Producto{
		SKU:      "FAL-881952291",
		Name:     "Test",
		Brand:    "Test",
		Size:     "XXL",
		Price:    "99",
		Imagenes: []Imagen{},
		Estado:   "A",
	}
	auxI := Imagen{
		Producto: 0,
		Url:      "TEST.test.cl",
		Estado:   "A",
	}
	auxP.Imagenes = append(auxP.Imagenes, auxI)
	got, err := funcionModificar(auxP)
	if got.Estado == "Producto modificado" || err.Error() != "el rango de precios es de 1.00 a 99999999.00" {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(got.ID), "nil, 0")
		return
	}
}
func TestModificar_7(t *testing.T) {
	//se espera que falle al usar un precio muy alto
	auxP := Producto{
		SKU:      "FAL-881952291",
		Name:     "Test",
		Brand:    "Test",
		Size:     "XXL",
		Price:    "10000000000",
		Imagenes: []Imagen{},
		Estado:   "A",
	}
	auxI := Imagen{
		Producto: 0,
		Url:      "TEST.test.cl",
		Estado:   "A",
	}
	auxP.Imagenes = append(auxP.Imagenes, auxI)
	got, err := funcionModificar(auxP)
	if got.Estado == "Producto modificado" || err.Error() != "el rango de precios es de 1.00 a 99999999.00" {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(got.ID), "nil, 0")
		return
	}
}
func TestModificar_8(t *testing.T) {
	//se espera que funcione pero al recuperar la talla aparesca como "N/A"
	auxP := Producto{
		SKU:      "FAL-881952293",
		Name:     "Test",
		Brand:    "Test",
		Size:     "",
		Price:    "1000000",
		Imagenes: []Imagen{},
		Estado:   "A",
	}
	auxI := Imagen{
		Producto: 0,
		Url:      "TEST.test.cl",
		Estado:   "A",
	}
	auxP.Imagenes = append(auxP.Imagenes, auxI)
	got, err := funcionModificar(auxP)
	if got.Estado != "Producto modificado" || err != nil {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(got.ID), "nil, 0")
		return
	}
	got, err = consultaProducto("A", auxP.SKU)
	if got.ID == 0 || err != nil || got.Size != "N/A" {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(got.ID), "nil, 0")
		return
	}
}

func TestModificar_9(t *testing.T) {
	//se espera que falle al usar marca con nombre muy corto
	auxP := Producto{
		SKU:      "FAL-881952291",
		Name:     "Test",
		Brand:    "Te",
		Size:     "XXL",
		Price:    "10000000",
		Imagenes: []Imagen{},
		Estado:   "A",
	}
	auxI := Imagen{
		Producto: 0,
		Url:      "TEST.test.cl",
		Estado:   "A",
	}
	auxP.Imagenes = append(auxP.Imagenes, auxI)
	got, err := funcionModificar(auxP)
	if got.Estado == "Producto modificado" || err.Error() != "incluir una marca con nombre de largo entre 3 y 50" {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(got.ID), "nil, 0")
		return
	}
}
func TestModificar_10(t *testing.T) {
	//se espera que falle al usar marca con nombre muy largo
	auxP := Producto{
		SKU:      "FAL-881952291",
		Name:     "Test",
		Brand:    "Test-Test-Test-Test-Test-Test-Test-Test-Test-Test-Test",
		Size:     "XXL",
		Price:    "10000000",
		Imagenes: []Imagen{},
		Estado:   "A",
	}
	auxI := Imagen{
		Producto: 0,
		Url:      "TEST.test.cl",
		Estado:   "A",
	}
	auxP.Imagenes = append(auxP.Imagenes, auxI)
	got, err := funcionModificar(auxP)
	if got.Estado == "Producto modificado" || err.Error() != "incluir una marca con nombre de largo entre 3 y 50" {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(got.ID), "nil, 0")
		return
	}
}
func TestModificar_11(t *testing.T) {
	//se espera que falle al usar nombre muy corto
	auxP := Producto{
		SKU:      "FAL-881952291",
		Name:     "Te",
		Brand:    "Test",
		Size:     "XXL",
		Price:    "10000000",
		Imagenes: []Imagen{},
		Estado:   "A",
	}
	auxI := Imagen{
		Producto: 0,
		Url:      "TEST.test.cl",
		Estado:   "A",
	}
	auxP.Imagenes = append(auxP.Imagenes, auxI)
	got, err := funcionModificar(auxP)
	if got.Estado == "Producto modificado" || err.Error() != "incluir un nombre o descripcion de largo entre 3 y 50" {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(got.ID), "incluir un nombre o descripcion de largo entre 3 y 50, 0")
		return
	}
}
func TestModificar_12(t *testing.T) {
	//se espera que falle al usar nombre muy largo
	auxP := Producto{
		SKU:      "FAL-881952291",
		Name:     "Test-Test-Test-Test-Test-Test-Test-Test-Test-Test-Test",
		Brand:    "Tes",
		Size:     "XXL",
		Price:    "10000000",
		Imagenes: []Imagen{},
		Estado:   "A",
	}
	auxI := Imagen{
		Producto: 0,
		Url:      "TEST.test.cl",
		Estado:   "A",
	}
	auxP.Imagenes = append(auxP.Imagenes, auxI)
	got, err := funcionModificar(auxP)
	if got.Estado == "Producto modificado" || err.Error() != "incluir un nombre o descripcion de largo entre 3 y 50" {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(got.ID), "incluir un nombre o descripcion de largo entre 3 y 50, 0")
		return
	}
}

//Test Avanzados

func TestAvanzado1(t *testing.T) {
	//crear una solicitud, consultarla, modificarla, consultarla, eliminarla, consultarla, modificarla, consultarla
	auxP := Producto{
		SKU:      "FAL-881952104",
		Name:     "TestAvanzado",
		Brand:    "Test",
		Size:     "XXL",
		Price:    "5999.99",
		Imagenes: []Imagen{},
		Estado:   "A",
	}
	auxI := Imagen{
		Producto: 0,
		Url:      "TEST.test.cl",
		Estado:   "A",
	}
	auxP.Imagenes = append(auxP.Imagenes, auxI)
	got, err := funcionCrear(auxP)
	if got.ID == 0 || err != nil || got.Estado != "Producto ingresado" {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(got.ID), "nil, >0")
		return
	}
	got, err = consultaProducto("A", auxP.SKU)
	if got.ID == 0 || err != nil {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(got.ID), "nil, >0")
		return
	}
	auxP.Brand = "Brand.test"
	got, err = funcionModificar(auxP)
	if got.Estado != "Producto modificado" || err != nil {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(got.ID), "nil, >0")
		return
	}

	got, err = consultaProducto("A", auxP.SKU)
	if got.ID == 0 || err != nil || got.Brand != "Brand.test" {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(got.ID), "nil, >0")
		return
	}

	got, err = funcionEliminar(Producto{SKU: auxP.SKU})
	if got.Estado != "Producto Elimiinado" || err != nil {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(got.ID), "nil, >0")
		return
	}
	got, err = consultaProducto("A", auxP.SKU)
	if got.ID != 0 || err != nil {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(got.ID), "nil, >0")
		return
	}
	auxP.Price = "5990"
	got, err = funcionModificar(auxP)
	if got.Estado != "Producto modificado" || err != nil {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(got.ID), "nil, >0")
		return
	}

	got, err = consultaProducto("A", auxP.SKU)
	if got.ID == 0 || err != nil || got.Price != "59.90" {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(got.ID), "nil, >0")
		return
	}
}

func TestAvanzado2(t *testing.T) {
	//Consultar todas las solicitudes que no son parte de la data de prueba y eliminarlas
	got, err := funcionConsultaMasiva()
	if len(got) == 0 || err != nil {
		t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(len(got)), "nil, >0")
		return
	}
	for _, i2 := range got {
		if i2.SKU != "FAL-8406270" && i2.SKU != "FAL-881952283" && i2.SKU != "FAL-881898502" {
			got2, err := funcionEliminar(Producto{SKU: i2.SKU})
			if got2.Estado != "Producto Elimiinado" || err != nil {
				t.Errorf("got %q, wanted %q", err.Error()+" , "+strconv.Itoa(got2.ID), "nil, >0")
				return
			}
		}
	}

}
