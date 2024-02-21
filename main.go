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
    ID, Data, Lembrete, RelatoDoDia, Treino, AFazer, LeuHoje, Ingles, TirouFoto, Creatina, Lendo string
}

func main() {
	registros := getRegistros()

    fmt.Println("Serve is running on http://localhost:8080/ 🚀")
    tmpl := template.Must(template.ParseFiles("./html/template/index.html"))

    // Função para manipular a requisição "/"
    http.HandleFunc("/registros", func(w http.ResponseWriter, r *http.Request) {
        // Escreve o cabeçalho da resposta
        w.WriteHeader(http.StatusOK)

        // Define o tipo de conteúdo da resposta como HTML
        w.Header().Set("Content-Type", "text/html")

        // Executa o template e escreve o resultado na resposta
        tmpl.Execute(w, registros)
    })

    http.HandleFunc("/update_registro", editRegistro)


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

func editRegistro(w http.ResponseWriter, r *http.Request){
    // Obter dados do formulário
    err := r.ParseForm()
    if err != nil {
        fmt.Fprintf(w, "Erro ao analisar o formulário: %v", err)
        return
    }
    // Validar dados

    db, err := sql.Open("postgres", "postgres://postgres:689df2c8@localhost:5432/Registros?sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Criar instrução SQL para atualizar o registro
    stmt, err := db.Prepare("UPDATE diarios SET data = $1, lembrete = $2, relatoDoDia = $3, treino = $4, aFazer = $5, leujoje = $6, ingles = $7, tirouFoto = $8, creatina = $9, lendo = $10 WHERE id = $11;")
    if err != nil {
        log.Fatal(err)
    }
    defer stmt.Close()

    // Substituir os valores na instrução SQL
    
    data := "21-02-2024"
    lembrete := "lembrete"
    relatoDoDia := "relatoDoDia"
    treino := "treino"
    aFazer := "aFazer"
    leuHoje := "leujoje"
    ingles := "ingles"
    tirouFoto := "tirouFoto"
    creatina := "creatina"
    lendo := "lendo"
    id := 4
    
    // Executar a instrução SQL e verificar se a atualização foi bem-sucedida
    _, err = stmt.Exec(data, lembrete, relatoDoDia, treino, aFazer, leuHoje, ingles, tirouFoto, creatina, lendo, id)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Registro atualizado com sucesso!")
}