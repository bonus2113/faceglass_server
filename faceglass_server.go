package main

import (
    "fmt"
    "log"
    "io"
    "net/http"
    "strconv"
    "encoding/json"
    "github.com/gorilla/mux"
    "os"
)

type User struct {
    ID        int
    Name      string
    Text      string
    Status    string
}

type Users []User

func main() {
    
    router := mux.NewRouter().StrictSlash(true)
    
    router.PathPrefix("/asset/").Handler( http.StripPrefix("/asset/", http.FileServer(http.Dir("./asset/"))) )
  
    router.HandleFunc("/label", getLabelHandler).Methods("GET")
    router.HandleFunc("/users", userIndex)
    router.HandleFunc("/users/{userId}", userShow).Methods("GET")
    router.HandleFunc("/users/{userId}", addUser).Methods("POST")
 
    os.MkdirAll("./asset/users", 0777)
    
    fmt.Println("Serving content at :8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}

func getLabelHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, getLabel(10));
}

func index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

// upload logic
func addUser(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method)
    r.ParseMultipartForm(32 << 20)
    fmt.Println(r.Form.Encode());
    file, handler, err := r.FormFile("image")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()
    
    
    fmt.Fprintf(w, "%v", handler.Header)
    vars := mux.Vars(r)
    userID := vars["userId"]

    os.MkdirAll("./asset/users/" + userID, 0777);

    f, err := os.OpenFile("./asset/users/" + userID + "/" + handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer f.Close()
    io.Copy(f, file)
}

func userIndex(w http.ResponseWriter, r *http.Request) {
    users := Users{
        User{ID: 0, Name: "dario", Text: "programmer, backend", Status: "DnD"},
        User{ID: 1, Name: "alexander", Text: "programmer, frontend", Status: "Come talk to me"},
        User{ID: 2, Name: "yan_wo", Text: "programmer, frontend", Status: "DnD"},
        User{ID: 3, Name: "jenny_li", Text: "designer", Status: "BrB"},
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(users); err != nil {
        panic(err)
    }
}

func userShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userID, err := strconv.Atoi(vars["userId"])
    if err != nil {
        w.WriteHeader(http.StatusNotFound)
        return
    }
    
    users := Users{
        User{ID: 0, Name: "dario", Text: "programmer, backend", Status: "DnD"},
        User{ID: 1, Name: "alexander", Text: "programmer, frontend", Status: "Come talk to me"},
        User{ID: 2, Name: "yan_wo", Text: "programmer, frontend", Status: "DnD"},
        User{ID: 3, Name: "jenny_li", Text: "designer", Status: "BrB"},
    }
    
    var foundUser User
    
    foundUser.ID = -1
    
    for _,user := range users {
        if user.ID == userID {
            foundUser = user
            break
        }
    } 
    
    if foundUser.ID == -1 {
        w.WriteHeader(http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(foundUser); err != nil {
        panic(err)
    }
}
