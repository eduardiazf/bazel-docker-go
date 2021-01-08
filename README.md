# bazel-docker-go
Monorepo Golang with Bazel and Docker

Vas a aprender a utilizar Bazel y Docker como una opción para la creación de monorepo en Go 

La administración de dependencias y construcción de binarios es una de las partes mas frustrantes del desarrollar código al igual que tener que manejar múltiples proyectos con diferentes herramientas de construcción.

## Documentation
### Que es Bazel?
Bazel es una herramienta de software gratuita que permite la automatización de la construcción y prueba de software. 

### Por que Bazel?
<strong>Acelera tus construcciones y pruebas</strong>
bazel unicamente reconstruye lo necesario, Con almacenamiento en caché local y distribuido avanzado, análisis de dependencia optimizado y ejecución paralela, obtiene compilaciones incrementales y rápidas.

<strong>Scalable</strong>
Bazel helps you scale your organization, codebase and Continuous Integration system. It handles codebases of any size, in multiple repositories or a huge monorepo.

<strong>Una herramienta para todos</strong>
Puedes construir y testear múltiples lenguajes con una sola herramienta en un monorepo

---
### Primeros pasos 
Creamos un proyecto en Go con estos comandos:
```
mkdir -p bazel-docker-go
go mod init bairesapp
```




