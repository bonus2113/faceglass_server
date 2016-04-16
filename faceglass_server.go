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
}

type Users []User

func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", index)
    router.HandleFunc("/users", userIndex)
    router.HandleFunc("/users/{userId}", userShow).Methods("GET")
    router.HandleFunc("/users/{userId}", upload).Methods("POST")
    
    log.Fatal(http.ListenAndServe(":8080", router))
}

func index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}



// upload logic
func upload(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method)
    r.ParseMultipartForm(32 << 20)
    file, handler, err := r.FormFile("image")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()
    
    
    fmt.Fprintf(w, "%v", handler.Header)
    vars := mux.Vars(r)
    userID := vars["userId"]

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
        User{Name: "dario", Text: "programmer, backend"},
        User{Name: "alexander", Text: "programmer, frontend"},
        User{Name: "yan_wo", Text: "programmer, frontend"},
        User{Name: "jenny_li", Text: "designer"},
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
        User{ID: 0, Name: "dario", Text: "programmer, backend"},
        User{ID: 1, Name: "alexander", Text: "programmer, frontend"},
        User{ID: 2, Name: "yan_wo", Text: "programmer, frontend"},
        User{ID: 3, Name: "jenny_li", Text: "designer"},
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