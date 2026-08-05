# Learning Go

Repositorio dedicado al aprendizaje práctico del lenguaje de programación **Go (Golang)**. Cada carpeta contiene ejemplos independientes y explicaciones sobre los conceptos fundamentales del lenguaje.

## 📁 Estructura del Proyecto

El proyecto está organizado por temas para facilitar el estudio progresivo:

```text
learning-go/
│
├── ⚙️ Inicio y Configuración
│   ├── Introduccion/        # Primeros pasos, estructura básica y "Hello World"
│   └── Ambiente/            # Configuración del entorno de desarrollo Go
│
├── 🧱 Fundamentos del Lenguaje
│   ├── Variables/           # Declaración de variables, constantes e inferencia de tipos
│   ├── Valores/             # Tipos de datos primitivos (int, float, string, bool)
│   └── Runas/               # Tipo rune y manejo de caracteres Unicode en Go
│
├──  Control de Flujo
│   ├── For/                 # Bucles, iteraciones y uso de range
│   ├── Condicionales/       # Sentencias if, else if y else
│   └── Switch/              # Estructuras switch-case y fallthrough
│
├── ⚡ Funciones
│   ├── Funciones/           # Declaración, parámetros y retorno de funciones
│   ├── FuncionesMultiples/  # Retorno de múltiples valores en una función
│   ├── Variadicas/          # Funciones variádicas (número variable de argumentos)
│   └── Recursividad/        # Funciones recursivas y casos base
│
├── 🗄️ Estructuras de Datos
│   ├── Arrays/              # Arreglos de tamaño fijo en Go
│   ├── slices/              # Slices: arreglos dinámicos y flexibles
│   ├── Maps/                # Mapas (diccionarios) clave-valor
│   ├── Structs/             # Estructuras y composición de tipos personalizados
│   ├── Enuns/               # Simulación de enumeraciones con iota y constantes
│   └── HASHES/              # Algoritmos de hashing y funciones hash en Go
│
├── 🧠 Punteros y Memoria
│   └── Punteros/            # Punteros, paso por referencia y manejo de memoria
│
├──  Concurrencia
│   ├── Goroutines/          # Goroutines: hilos ligeros de Go
│   ├── Channels/            # Canales de comunicación entre goroutines
│   ├── Select/              # Multiplexación de canales con select
│   ├── Workers/             # Patrón worker pool y concurrencia avanzada
│   └── Timeout/             # Manejo de timeouts y contextos en concurrencia
│
├──  Manejo de Errores
│   └── Errores/             # Tratamiento de errores, errors.Is, errors.As y panic/recover
│
├── 📦 Paquetes y Módulos
│   └── packages/            # Importación y creación de paquetes propios
│
├──  Formatos de Datos
│   └── Json/                # Serialización y deserialización de JSON en Go
│
├── go.mod                   # Configuración del módulo Go
└── README.md                # Documentación del proyecto
```
