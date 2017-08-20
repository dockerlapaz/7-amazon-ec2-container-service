// Docker La Paz
// Votacion entre Amazon ECS vs. Kubernetes

// Empezamos
package main

// Importar todos los paquetes
import (
  "fmt"
  "html/template"
  "log"
  "net/http"
  "strconv"
  "github.com/garyburd/redigo/redis"
)

// Funcion principal
func main() {
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
}

func Votar(w http.ResponseWriter, r *http.Request) {
  // Parsear el formulario
  r.ParseForm()
  // Obtenemos el voto
  voto := r.Form.Get("scheduler")

  // Logear el voto
  fmt.Println("El usuario ha votado por ", voto)

  // Nos conectamos a Redis
  c, err := redis.Dial("tcp", ":6379")
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

  // Otra conexion a redis
  // TODO: declarar la conexion globalmente
  c, err := redis.Dial("tcp", ":6379")
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