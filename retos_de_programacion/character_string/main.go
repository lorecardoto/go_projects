/*
 * EJERCICIO:
 * Muestra ejemplos de todas las operaciones que puedes realizar con cadenas de caracteres
 * en tu lenguaje. Algunas de esas operaciones podrían ser (busca todas las que puedas):
 * - Acceso a caracteres específicos, subcadenas, longitud, concatenación, repetición,
 *   recorrido, conversión a mayúsculas y minúsculas, reemplazo, división, unión,
 *   interpolación, verificación...
 *
 * DIFICULTAD EXTRA (opcional):
 * Crea un programa que analice dos palabras diferentes y realice comprobaciones
 * para descubrir si son:
 * - Palíndromos
 * - Anagramas
 * - Isogramas
 */
package main

import (
	"fmt"
	"strings"
)

var cadena string
var cadena2 string

 func main() {

	for {
		fmt.Println("*******CADENA DE CARACTERES******")
        fmt.Println("Seleccione una operación:")
        fmt.Println("1. Insertar cadena de caracteres")
        fmt.Println("2. Ingresar segunda cadena de caracteres")
        fmt.Println("3. Operaciones basicas con cadena de caracteres")
        fmt.Println("4. Operaciones basicas con la segunda cadena de caracteres")
        fmt.Println("5. Palindromos")
        fmt.Println("6. Anagramas")
		fmt.Println("7. Isogramas")
        fmt.Println("8. Salir")
		fmt.Println("********************************")
		var option int
		fmt.Scan(&option)
	
		switch option {
		case 1:
			insertarCadena()
		case 2:
			insertarSegundaCadena()
		case 3:
			operacionesBasicas()
		case 4:
			operacionesBasicas2()
		case 5:
			palindromos()
		case 6:
			anagramas()
		case 7:
			isogramas()
		case 8:
			fmt.Println("Hasta luego")
			return
		default:
			fmt.Println("Opción no válida")
		}
	
	}

 }

 func insertarCadena() {
	fmt.Print("Ingrese una cadena de caracteres: ")
	fmt.Scan(&cadena)
	 }

func insertarSegundaCadena() {
	fmt.Print("Ingrese otra cadena de caracteres: ")
	fmt.Scan(&cadena2)
}

func operacionesBasicas() {
	if len (cadena) == 0 {
		fmt.Println("No se ha ingresado una cadena de caracteres")
		return
	}
	fmt.Println("El primer caracter de la cadena es:", string(cadena[0]))
	fmt.Println("El último caracter de la cadena es:", string(cadena[len(cadena)-1]))
	fmt.Println("El caracter en la posición 5 es:", string(cadena[5]))
	fmt.Println("La concatenación de las dos cadenas es:", cadena + cadena2)
	fmt.Println("La longitud de la cadena es:", len(cadena))
	cadenaMayusculas := strings.ToUpper(cadena)
	fmt.Println("La cadena en mayúsculas es:", cadenaMayusculas)
	cadenaMinusculas := strings.ToLower(cadena)
	fmt.Println("La cadena en minúsculas es:", cadenaMinusculas)
	fmt.Println("La cadena repetida 3 veces es:", cadena2)

}	

func operacionesBasicas2() {
	if len(cadena2) == 0 {
		fmt.Println("No se ha ingresado una cadena de caracteres")
		return
	}
	fmt.Println("El primer caracter de la cadena es:", string(cadena2[0]))
	fmt.Println("El último caracter de la cadena es:", string(cadena2[len(cadena2)-1]))
	fmt.Println("El caracter en la posición 5 es:", string(cadena2[5]))
	fmt.Println("La concatenación de las dos cadenas es:", cadena + cadena2)
	fmt.Println("La longitud de la cadena es:", len(cadena2))
	cadenaMayusculas := strings.ToUpper(cadena2)
	fmt.Println("La cadena en mayúsculas es:", cadenaMayusculas)
	cadenaMinusculas := strings.ToLower(cadena2)
	fmt.Println("La cadena en minúsculas es:", cadenaMinusculas)
	fmt.Println("La cadena repetida 3 veces es:", cadena2)
}

func palindromos() {
	fmt.Print("Ingrese una cadena de caracteres: ")
	var cadena string
	fmt.Scan(&cadena)
	var cadenaInvertida string
	for i := len(cadena) - 1; i >= 0; i-- {
		cadenaInvertida += string(cadena[i])
	}
	if cadena == cadenaInvertida {
		fmt.Println("La cadena es un palíndromo")
	} else {
		fmt.Println("La cadena no es un palíndromo")
	}
}

func anagramas() {
	fmt.Print("Ingrese una cadena de caracteres: ")
	var cadena string
	fmt.Scan(&cadena)
	fmt.Print("Ingrese otra cadena de caracteres: ")
	var cadena2 string
	fmt.Scan(&cadena2)
	if len(cadena) != len(cadena2) {
		fmt.Println("Las cadenas no son anagramas")
		return
	}
	for i := 0; i < len(cadena); i++ {
		if !strings.Contains(cadena2, string(cadena[i])) {
			fmt.Println("Las cadenas no son anagramas")
			return
		}
	}
	fmt.Println("Las cadenas son anagramas")
}

func isogramas() {
	fmt.Print("Ingrese una cadena de caracteres: ")
	var cadena string
	fmt.Scan(&cadena)
	for i := 0; i < len(cadena); i++ {
		for j := i + 1; j < len(cadena); j++ {
			if cadena[i] == cadena[j] {
				fmt.Println("La cadena no es un isograma")
				return
			}
		}
	}
	fmt.Println("La cadena es un isograma")
}


