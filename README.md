# Hello Server

Ejemplo de Hola Munfo con un servidor Go y GoFr.

```go
package main

import "gofr.dev/pkg/gofr"

func main() {
	app := gofr.New()

	app.GET("/greet", func(ctx *gofr.Context) (any, error) {
		return "Hello, World!", nil
	})

	app.Run()
}
```

Detallando el ejemplo realizado:
1. Creando el servidor con GoFr
La instrucción gofr.New() crea una instancia de la aplicación. 
2. Definiendo la ruta greet
La siguiente instrucción define una ruta HTTP Get con una función anónima como parámetro. La cual toma un contexto y devuelve la resupesta "Hello, World!" y un segundo parámetro de error a nulo.
```go
app.GET('/greet', func(ctx *gofr.Context) (any, error) {
		return "Hello, World!", nil
	})
```
3. Iniciando el servidor
La siguiente instrucción inicia el servidor, momento en el que el mismo comienza a escuchar las solicitudes entrantes.