package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "os"
    "strconv"
)

type Usuario struct {
    ID     int    `json:"id"`
    Nombre string `json:"name"`
    Email  string `json:"email"`
}

var usuarios []Usuario

// Usuarios
func Users(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        w.Header().Set("Content-Type", "application/json")
        w.Header().Set("test-header", "header")
        json.NewEncoder(w).Encode(usuarios)
        
    case http.MethodPost:
        body, err := io.ReadAll(r.Body)
        if err != nil {
            http.Error(w, "Error al leer el body", http.StatusBadRequest)
            return
        }
        var user Usuario
        err = json.Unmarshal(body, &user)
        if err != nil {
            http.Error(w, "Error parseando el JSON", http.StatusBadRequest)
            return
        }
        user.ID = len(usuarios) + 1
        usuarios = append(usuarios, user)

        w.Header().Set("Content-Type", "application/json")
        w.Header().Set("test-header", "header")
        json.NewEncoder(w).Encode(user)
        
    case http.MethodPut:
        // Leer el ID desde la URL (asumimos que viene como query parameter ?id=)
        idStr := r.URL.Query().Get("id")
        if idStr == "" {
            http.Error(w, "ID es requerido", http.StatusBadRequest)
            return
        }
        
        id, err := strconv.Atoi(idStr)
        if err != nil {
            http.Error(w, "ID inválido", http.StatusBadRequest)
            return
        }

        // Leer el body
        body, err := io.ReadAll(r.Body)
        if err != nil {
            http.Error(w, "Error al leer el body", http.StatusBadRequest)
            return
        }
        
        var updatedUser Usuario
        err = json.Unmarshal(body, &updatedUser)
        if err != nil {
            http.Error(w, "Error parseando el JSON", http.StatusBadRequest)
            return
        }

        // Buscar y actualizar el usuario
        for i, user := range usuarios {
            if user.ID == id {
                updatedUser.ID = id // Mantener el mismo ID
                usuarios[i] = updatedUser
                w.Header().Set("Content-Type", "application/json")
                json.NewEncoder(w).Encode(updatedUser)
                return
            }
        }
        http.Error(w, "Usuario no encontrado", http.StatusNotFound)
        
    case http.MethodDelete:
        // Leer el ID desde la URL (asumimos que viene como query parameter ?id=)
        idStr := r.URL.Query().Get("id")
        if idStr == "" {
            http.Error(w, "ID es requerido", http.StatusBadRequest)
            return
        }
        
        id, err := strconv.Atoi(idStr)
        if err != nil {
            http.Error(w, "ID inválido", http.StatusBadRequest)
            return
        }

        // Buscar y eliminar el usuario
        for i, user := range usuarios {
            if user.ID == id {
                usuarios = append(usuarios[:i], usuarios[i+1:]...)
                w.WriteHeader(http.StatusNoContent) // 204 indica éxito sin contenido
                return
            }
        }
        http.Error(w, "Usuario no encontrado", http.StatusNotFound)
        
    default:
        http.Error(w, "Método no permitido", 405)
    }
}

func Ping(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        fmt.Fprintln(w, "pong")
    default:
        http.Error(w, "Método no permitido", 405)
    }
}

func Index(w http.ResponseWriter, r *http.Request) {
    content, err := os.ReadFile("./public/index.html")
    if err != nil {
        fmt.Fprintln(w, "error leyendo el html")
        return
    }
    fmt.Fprintln(w, string(content))
}

func main() {
    usuarios = append(usuarios, Usuario{
        ID:     1,
        Nombre: "Alfredo",
        Email:  "Alfredo@mail.com",
    })
    http.HandleFunc("/ping", Ping)
    http.HandleFunc("/v1/users", Users)
    http.HandleFunc("/", Index)

    fmt.Println("Servidor escuchando en el puerto 3000")
    http.ListenAndServe(":3000", nil)
}