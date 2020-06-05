package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//
type Filme struct {
	Titulo   string `json:"Titulo"`
	Ano      string `json:"Ano"`
	Conteudo string `json:"Conteudo"`
}

//
var Filmes = []Filme{
	{Titulo: "Monstros S.A", Ano: "2005", Conteudo: "Filme"},
	{Titulo: "Toy Story", Ano: "2004", Conteudo: "Filme"},
	{Titulo: "Vida de Inseto", Ano: "2003", Conteudo: "Filme"},
}

func main() {
	handleRequests()
}

func testPostretorneTodosFilmes(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Teste POST endpoint worked")

}

func paginaPrincipal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem vindo a minha lista de filmes!")
}

func handleRequests() {

	meuRouter := mux.NewRouter().StrictSlash(true)

	meuRouter.HandleFunc("/", paginaPrincipal)
	meuRouter.HandleFunc("/Filmes", retorneTodosFilmes).Methods("GET")
	meuRouter.HandleFunc("/Filmes", testPostretorneTodosFilmes).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", meuRouter))
}

func retorneTodosFilmes(w http.ResponseWriter, r *http.Request) {
	fmt.Println("retorneTodosFilmes")
	json.NewEncoder(w).Encode(Filmes)
}
