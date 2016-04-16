package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"
    "encoding/json"
    "github.com/gorilla/mux"
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
    router.HandleFunc("/users/{userId}", userShow)

    log.Fatal(http.ListenAndServe(":8080", router))
}

func index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
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