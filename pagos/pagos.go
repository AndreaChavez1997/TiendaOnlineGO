package pagos

import "fmt"

var pagosRegistrados = make(map[string]float64)


func RegistrarPago(cliente string, valor float64) {
fmt.Println("\n=== PAGOS ===")
	pagosRegistrados[cliente] = valor

	fmt.Println("\nPago registrado")
}

func ConsultarPagos() {

	

	for cliente, valor := range pagosRegistrados {
		fmt.Println(cliente, "=> $", valor)
	}
}

func ConsultarEnvio() {
	fmt.Println("Estado del envío: En proceso")
}