/*
 * EJERCICIO:
 * - Muestra ejemplos de creación de todas las estructuras soportadas por defecto
 *   en tu lenguaje.
 * - Utiliza operaciones de inserción, borrado, actualización y ordenación.
 *
 * DIFICULTAD EXTRA (opcional):
 * Crea una agenda de contactos por terminal.
 * - Debes implementar funcionalidades de búsqueda, inserción, actualización
 *   y eliminación de contactos.
 * - Cada contacto debe tener un nombre y un número de teléfono.
 * - El programa solicita en primer lugar cuál es la operación que se quiere realizar,
 *   y a continuación los datos necesarios para llevarla a cabo.
 * - El programa no puede dejar introducir números de teléfono no numéricos y con más
 *   de 11 dígitos (o el número de dígitos que quieras).
 * - También se debe proponer una operación de finalización del programa.
 */
 package main

import (
    "fmt"
    "sort"
    "strconv"
    "strings"
)

type Contact struct {
    Name  string
    Phone string
}

func main() {
    contacts := []Contact{
        {Name: "Juan", Phone: "123456789"},
        {Name: "Pedro", Phone: "987654321"},
        {Name: "Maria", Phone: "123456789"},
    }

    for {
		fmt.Println("*******Agenda de contactos******")
        fmt.Println("Seleccione una operación:")
        fmt.Println("1. Insertar contacto")
        fmt.Println("2. Borrar contacto")
        fmt.Println("3. Actualizar contacto")
        fmt.Println("4. Ordenar contactos")
        fmt.Println("5. Buscar contacto")
        fmt.Println("6. Mostrar todos los contactos")
		fmt.Println("7. Marcha final")
        fmt.Println("8. Salir")
		fmt.Println("********************************")

        var option int
        fmt.Scan(&option)

        switch option {
        case 1:
            var name, phone string
            fmt.Println("Ingrese el nombre:")
            fmt.Scan(&name)
            fmt.Println("Ingrese el teléfono:")
            fmt.Scan(&phone)
            if isValidPhone(phone) {
                contacts = append(contacts, Contact{Name: name, Phone: phone})
            } else {
                fmt.Println("Número de teléfono no válido.")
            }
        case 2:
            var name string
            fmt.Println("Ingrese el nombre del contacto a borrar:")
            fmt.Scan(&name)
            contacts = deleteContact(contacts, name)
        case 3:
            var name, phone string
            fmt.Println("Ingrese el nombre del contacto a actualizar:")
            fmt.Scan(&name)
            fmt.Println("Ingrese el nuevo teléfono:")
            fmt.Scan(&phone)
            if isValidPhone(phone) {
                contacts = updateContact(contacts, name, phone)
            } else {
                fmt.Println("Número de teléfono no válido.")
            }
        case 4:
            sort.Slice(contacts, func(i, j int) bool {
                return contacts[i].Name < contacts[j].Name
            })
            fmt.Println("Contactos ordenados.")
        case 5:
            var name string
            fmt.Println("Ingrese el nombre del contacto a buscar:")
            fmt.Scan(&name)
            searchContact(contacts, name)
        case 6:
            fmt.Println("Lista de contactos:")
            for _, contact := range contacts {
                fmt.Printf("Nombre: %s, Teléfono: %s\n", contact.Name, contact.Phone)
            }
		case 7:
			fmt.Println("Marcha final.")
			return
        case 8:
            fmt.Println("Saliendo del programa.")
            return
        default:
            fmt.Println("Opción no válida, elige otra.")
        }
    }
}

func isValidPhone(phone string) bool {
    if len(phone) > 11 {
        return false
    }
    _, err := strconv.Atoi(phone)
    return err == nil
}

func deleteContact(contacts []Contact, name string) []Contact {
    for i, contact := range contacts {
        if strings.EqualFold(contact.Name, name) {
            return append(contacts[:i], contacts[i+1:]...)
        }
    }
    fmt.Println("Contacto no encontrado.")
    return contacts
}

func updateContact(contacts []Contact, name, phone string) []Contact {
    for i, contact := range contacts {
        if strings.EqualFold(contact.Name, name) {
            contacts[i].Phone = phone
            return contacts
        }
    }
    fmt.Println("Contacto no encontrado.")
    return contacts
}

func searchContact(contacts []Contact, name string) {
    for _, contact := range contacts {
        if strings.EqualFold(contact.Name, name) {
            fmt.Printf("Nombre: %s, Teléfono: %s\n", contact.Name, contact.Phone)
            return
        }
    }
    fmt.Println("Contacto no encontrado.")
}