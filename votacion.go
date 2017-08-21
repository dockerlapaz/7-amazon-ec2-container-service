// Docker La Paz
// Votacion entre Amazon ECS vs. Kubernetes

// Empezamos
package main

// Importar todos los paquetes
import (
  "os"
  "fmt"
  "html/template"
  "log"
  "net/http"
  "strconv"
  "github.com/garyburd/redigo/redis"
)

// Funcion principal
func main() {
  // Dar feedback al sysadmin es mucho muy importante
  log.Println("Iniciando la aplicación...")
  // Funcion principal
  http.HandleFunc("/", Inicio)
  // Aqui manejamos el voto
  http.HandleFunc("/votar", Votar)
  // Servimos archivos estaticos
  http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
  // Iniciamos el server
  log.Fatal(http.ListenAndServe(":8080", nil))
}

// Definimos estructuras
type ContarVotos struct {
  Ecs         int
  Kubernetes  int
  Host        string
}

func Votar(w http.ResponseWriter, r *http.Request) {
  // Parsear el formulario
  r.ParseForm()
  // Obtenemos el voto
  voto := r.Form.Get("scheduler")

  // Logear el voto
  fmt.Println("El usuario ha votado por ", voto)

  redis_db := ":6379"

  if os.Getenv("REDIS_DB") != "" {
    redis_db = os.Getenv("REDIS_DB")
  }

  // Nos conectamos a Redis
  log.Println("Conectando a ", redis_db)
  c, err := redis.Dial("tcp", redis_db)
    if err != nil {
      panic(err)
    }
  defer c.Close()

  // Incrementamos el voto en redis
  votar, err := redis.String(c.Do("INCR", voto))

  fmt.Println(votar)

  // Redirigimos a /
  http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Inicio(w http.ResponseWriter, r *http.Request){

  redis_db := ":6379"
  if os.Getenv("REDIS_DB") != "" {
    redis_db = os.Getenv("REDIS_DB")
  }

  // Otra conexion a redis
  // TODO: declarar la conexion globalmente
  log.Println("Conectando a ", redis_db)
  c, err := redis.Dial("tcp", redis_db)
    if err != nil {
      panic(err)
    }
  defer c.Close()

  // Obtener votos para ecs y kubernetes
  ecsval, err := redis.String(c.Do("GET","ecs"))
  ecs, err := strconv.Atoi(ecsval)

  // TODO: reemplazar redis.String en redis.Int
  kubeval, err := redis.String(c.Do("GET","kubernetes"))
  kubernetes, err := strconv.Atoi(kubeval)

  // Definimos la estructura para renderizar el html
  ObtenerVotos := ContarVotos{
    Ecs: ecs,
    Kubernetes: kubernetes,
    Host: os.Getenv("HOSTNAME"),
  }

  // Cargamos el archivo html
  t, err := template.ParseFiles("votacion.html")
  if err != nil { // if there is an error
    log.Print("template parsing error: ", err)
  }

  // Renderizar el html
  err = t.Execute(w, ObtenerVotos)
  if err != nil {
    log.Print("template executing error: ", err)
  }
}