#Saludos en GO

## Instalación
Ejecuta el siguiente comoando para instalar el paquete:

```bash
go get -u github.com/Cachosel/greetings

```

## Ejemplo de uso
Copia y pega este código para tener un ejemplo de como usar este módulo

```
package main

import (
	"fmt"

	"github.com/Cachosel/greetings"
)
func main() {
	message, err := greetings.Hello("River Plate")

	if err != nil {
		fmt.Println("Ocurrión un error", err)
		return
	}
	fmt.Println(message)
}
```