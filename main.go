package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	 _ "github.com/lib/pq"
)

type Registro struct {
	ID int
    Data, Lembrete, RelatoDoDia, Treino, AFazer, LeuHoje, Ingles, TirouFoto, Creatina, Lendo string
}

func main() {
	registros := getRegistros()

    fmt.Println("Serve is running on http://localhost:8080/ 🚀")
    tmpl := template.Must(template.ParseFiles("./html/template/index.html"))

    // Função para manipular a requisição "/"
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Escreve o cabeçalho da resposta
        w.WriteHeader(http.StatusOK)

        // Define o tipo de conteúdo da resposta como HTML
        w.Header().Set("Content-Type", "text/html")

        // Executa o template e escreve o resultado na resposta
        tmpl.Execute(w, registros)
    })

    // Inicia o servidor na porta 8080
    http.ListenAndServe(":8080", nil)
}

func getRegistros() []Registro {
	db, err := sql.Open("postgres", "postgres://postgres:689df2c8@localhost:5432/Registros?sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Executar uma consulta SQL para recuperar os dados
    rows, err := db.Query("SELECT * FROM diarios")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    // Ler os dados da consulta e preenchê-los na struct
    var registros []Registro
    for rows.Next() {
        var registro Registro
        err := rows.Scan(&registro.ID, &registro.Data, &registro.Lembrete, &registro.RelatoDoDia, &registro.Treino, &registro.AFazer, &registro.LeuHoje, &registro.Ingles, &registro.TirouFoto, &registro.Creatina, &registro.Lendo)
        if err != nil {
            log.Fatal(err)
        }
        registros = append(registros, registro)
    }

    return registros
}