package productos

import "fmt"


var listaProductos []string

func RegistrarProducto(nombre string) {
	
	listaProductos = append(listaProductos, nombre)
	
	fmt.Println("Producto registrado:", nombre)
}

func ConsultarProductos() {
	fmt.Println("\n=== CATÁLOGO DE PRODUCTOS ===")

	for i, producto := range listaProductos {
		fmt.Printf("%d. %s\n", i+1, producto)
	}
}

func ActualizarProducto(pos int, nuevo string) {
	if pos >= 0 && pos < len(listaProductos) {
		listaProductos[pos] = nuevo
	}
}

func EliminarProducto() {
	if len(listaProductos) > 0 {
		listaProductos = listaProductos[:len(listaProductos)-1]
	}
}