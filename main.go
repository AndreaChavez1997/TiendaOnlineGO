package main

import (
		"fmt"
		"github.com/AndreaChavez1997/AprendizajeAutonomo1ACH/usuarios"
		"github.com/AndreaChavez1997/AprendizajeAutonomo1ACH/productos"
		"github.com/AndreaChavez1997/AprendizajeAutonomo1ACH/compras"
		"github.com/AndreaChavez1997/AprendizajeAutonomo1ACH/pagos"
		
	
)


func main() {

	fmt.Println("\n===== SISTEMA E-COMMERCE =====")

	// VARIABLE
	tienda := "TIENDA ONLINE ANDY"

	fmt.Println(tienda)

	// ARRAY
	categorias := [3]string{
		"Blusas",
		"Pantalones",
		"Zapatos",
	}

	fmt.Println("\n=== CATEGORÍAS ===")

	for _, categoria := range categorias {
		fmt.Println(categoria)
	}

	// USUARIOS

	fmt.Println("\n=== NUEVO USUARIO ===")
	usuarios.RegistrarAdministrador("Andrea")
	usuarios.RegistrarAdministrador("Erik")

	usuarios.RegistrarCliente("Valery")
	usuarios.RegistrarCliente("Farid")

	usuarios.IniciarSesion("Andrea")

	usuarios.ConsultarAdministradores()
	usuarios.ConsultarClientes()

	// PRODUCTOS
	fmt.Println("\n=== INVENTARIO ===")
	
	productos.RegistrarProducto("Blusa Formal")
	productos.RegistrarProducto("Mom Jean")
	productos.RegistrarProducto("Zapatos Casuales")

	productos.ConsultarProductos()

	// COMPRAS
	fmt.Println("\n=== VENTA ===")
	compras.AgregarAlCarrito("Blusa Formal")
	compras.AgregarAlCarrito("Mom Jean")

	compras.MostrarCarrito()

	compras.GenerarPedido()

	// PAGOS
	
	pagos.RegistrarPago("Andrea", 75.30)

	pagos.ConsultarPagos()

	pagos.ConsultarEnvio()
}