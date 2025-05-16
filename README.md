# Calculadora de Matrices

## Reto Tecnico
Este prueba tecnica basicamente es una calculadora web que trabaja con matrices y tiene tres partes:

1. **Frontend**: Donde puedes ingresar y ver matrices
2. **Calculadora Principal (Go - Puerto 3000)**: Divide una matriz en dos partes (Q y R)
3. **Calculadora de Estadísticas (Node.js - Puerto 3001)**: Analiza las matrices y muestra información importante como:
   - Traza (suma de los números en la diagonal)
   - Rango (cuántas filas/columnas son independientes)
   - Determinante (un número que indica ciertas propiedades de la matriz)
   - Norma (mide el "tamaño" de la matriz)

## Paso a Paso del Proceso

## Cómo iniciar el proyecto

### Opción 1 - Usando Docker (mucho mas sencillo):
```powershell
docker-compose up --build
```

### Opción 2 - Manual (necesitas tener Go y Node.js instalados):
## Utilice manual, ya que mi docker desktop fallaba 

1. Inicia la calculadora principal:
```powershell
cd api-go
go run main.go
```

2. En otra ventana, inicia la calculadora de estadísticas:
```powershell
cd api-nodejs
npm install
node index.js
```

3. En una tercera ventana, inicia la página web:
```powershell
cd frontend
npm install
npm run serve
```

## Cómo usar la calculadora

1. Abre tu navegador y ve a: http://localhost:8080

2. En la página veras:
   - Selector para el tamaño de la matriz
   - Casillas para ingresar números
   - Botón "Calcular"
   - Botón "Limpiar Todo"

3. Ejemplo:
   - Selecciona una matriz 3x3
   - Ingresa números como:
     ```
     1  2  3
     4  5  6
     7  8  9
     ```
   - Haz clic en "Calcular"

## Funcion de la calculadora

1. Cuando presionas "Calcular":
   - Divide tu matriz en dos matrices especiales (Q y R)
   - Calcula estadísticas útiles de estas matrices
   - Muestra los resultados de forma ordenada

2. Las estadísticas que calcula incluyen:
   - **Traza**: Suma de los números en la diagonal
   - **Rango**: Qué tan "independientes" son las filas/columnas
   - **Determinante**: Un número que dice si la matriz es "invertible"
   - **Norma**: Mide el "tamaño total" de la matriz

## Ejemplo de posible resultado

#### Respuesta de la API Go:
```json
{
    "q": [
        [-0.123, 0.456, 0.789],
        [-0.234, 0.567, 0.890],
        [-0.345, 0.678, 0.901]
    ],
    "r": [
        [-14.282, -18.569, -22.855],
        [0, -1.198, -2.396],
        [0, 0, 0]
    ]
}
```

#### Estadísticas Calculadas (API Node.js):
```json
{
    "qStats": {
        "determinant": 1.0,
        "trace": 3.0,
        "norm": 4.242,
        "rank": 3
    },
    "rStats": {
        "determinant": 1.0,
        "trace": 3.0,
        "norm": 4.242,
        "rank": 3
    },
    "productStats": {
        "determinant": 1.0,
        "trace": 3.0
    }
}
```

## Solución de problemas

Si algo no funciona, verifica:

1. Que los servicios estén funcionando:
   - Página web: http://localhost:8080
   - Calculadora principal: http://localhost:3000/health
   - Calculadora de estadísticas: http://localhost:3001/health

2. Que no haya otros programas usando los puertos:
   - 8080 (página web)
   - 3000 (calculadora principal)
   - 3001 (calculadora de estadísticas)

3. Para detener todo:
   - Si usas Docker: Presiona Ctrl+C o escribe `docker-compose down`
   - Si lo iniciaste manualmente: Presiona Ctrl+C en cada ventana

## Nota importante agregado
Este proyecto usa:
- Go: Para cálculos matemáticos rápidos y precisos
- Node.js Express: Para análisis estadístico avanzado
- Vue.js: Para una interfaz web moderna y fácil de usar
- JWT: Para seguridad entre servicios
