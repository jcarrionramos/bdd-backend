# API - REST PROYECTO BDD

Para correr sus servidor necesitan:

1. Revisar que tengan setteado su GOPATH. Para setearlo deben modificar el archivo `.bashrc` ubicado en la carpeta __Home__ agregando las siguientes lineas:

         export GOPATH=$HOME/Documents/go
         PATH=$PATH:$GOPATH/bin

  En este caso el GOPATH está seteado en la ruta  `/Documents/go`.

2. Descargar el repo clonandolo o descargandolo como zip y extrayendolo en la carpeta `go/src/github.com/jcarrionramos`.

3. Entrar a la carpeta *bdd-backend* y ejecutar en la terminal `go run main.go`.

4. __ALTERNATIVAMENTE__, si les tira un error, deben ejecutar las siguientes dos lineas, en donde se instalan los framework GIN y MATTN:
         go get -u github.com/gin-gonic/gin
         go get github.com/mattn/go-sqlite3
