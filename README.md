# API - REST PROYECTO BDD

Para correr sus servidor necesitan:

1. Revisar que tengan setteado su GOPATH. Para revisar esto ejecutar `ECHO $GOPATH` en la terminal, debería arrojar la ruta de donde se encuentra su GOPATH.

2. Descargar el repo clonandolo o descargandolo como zip y exayendolo en la carpeta `go/src/github.com/jcarrionramos`.

3. Entrar a la carpeta *bdd-backend* y ejecutar en la terminal `go run main.go`.

4. ALTERNATIVAMENTE, si les tira un error, deben ejecutar las siguientes dos lineas, en donde se instalan los framework GIN y MATTN:
    1. `go get -u github.com/gin-gonic/gin`
    2. `go get github.com/mattn/go-sqlite3`
