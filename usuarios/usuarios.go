package usuarios

import "fmt"

var clientes []string
var administradores []string

func RegistrarCliente(nombre string) {
	
	clientes = append(clientes, nombre)
	fmt.Println("Registro de nuevo cliente:", nombre)
	
}

func RegistrarAdministrador(nombre string) {
	administradores = append(administradores, nombre)
	fmt.Println("Registro de nuevo administrador", nombre)
	
}

func IniciarSesion(nombre string) {
	fmt.Println("\nSesión iniciada por:", nombre)
	fmt.Println()
}

func ConsultarClientes() {

	fmt.Println("\n CLIENTES")

	for i, cliente := range clientes {
		fmt.Printf("%d. %s\n", i+1, cliente)
	
	}
}

func ConsultarAdministradores() {

	fmt.Println("\n ADMINISTRADORES")

	for i, administrador := range administradores {
		fmt.Printf("%d. %s\n", i+1, administrador)
		
	}
}