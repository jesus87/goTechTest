# IPCOM Problem Parte I

# Problemática

Un cliente de IPCOM realiza ventas en sus call center y desea un API para analizarlas. Las ventas
pueden ser consultadas mediante un API REST.

# Ejecuçión del API

## Build

Ejecute en su línea de comandos el comando

```bash
go build
```

## Run

Ejecute en su línea de comandos el comando

```bash
./ipcom
```

## Probar usando un HTTP Cliente como Postman o Insomnia

Agregue un nuevo request del tipo GET y ejecute el siguiente endpoint de su local

```bash
 http://localhost:8080/resumen/2020-02-20?dias=5
```

## Run Tests

Ejecute en su línea de comandos el comando

```bash
go test
```

