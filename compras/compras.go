package compras

import "fmt"


var carrito []string

func AgregarAlCarrito(producto string) {
	
	carrito = append(carrito, producto)
	
	fmt.Println("Producto vendido:", producto)
}

func EliminarDelCarrito() {

	if len(carrito) > 0 {
		carrito = carrito[:len(carrito)-1]
	}

	fmt.Println("Último producto eliminado")
}

func MostrarCarrito() {

	fmt.Println("\n=== CARRITO ===")

	for _, producto := range carrito {
		fmt.Println(producto)
	}
}

func GenerarPedido() {
	fmt.Println("\n*********************************")
	fmt.Println("Pedido generado correctamente")
}