package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

type Registro struct {
    ID, Data, Lembrete, RelatoDoDia, Treino, AFazer, LeuHoje, Ingles, TirouFoto, Creatina, Lendo string
}

type DadosDiarios struct {
    Data       string `json:"data"`
    Lembrete   string `json:"lembrete"`
    RelatoDoDia string `json:"relato_do_dia"`
    Treino      string `json:"treino"`
    AFazer     string `json:"a_fazer"`
    LeuHoje     string `json:"leu_hoje"`
    Ingles      string `json:"ingles"`
    TirouFoto  string `json:"tirou_foto"`
    Creatina   string `json:"creatina"`
    Lendo       string `json:"lendo"`
    ID         int    `json:"id"`
}

type Data struct {
    Mensagem string
  }

func main() {
    fmt.Println("Serve is running on http://localhost:8080/ 🚀")
    mux := http.NewServeMux()
    handler := cors.Default().Handler(mux)

    mux.HandleFunc("/js/main.js", func(w http.ResponseWriter, r *http.Request) {
        // Leia o arquivo JavaScript.
        data, err := ioutil.ReadFile("main.js")
        if err != nil {
            fmt.Fprintf(w, "Erro ao ler o arquivo JavaScript: %v", err)
            return
        }

        // Escreva o conteúdo do arquivo JavaScript na resposta HTTP.
        w.Header().Set("Content-Type", "application/javascript")
        w.Write(data)
    })

    mux.HandleFunc("/js/script.js", func(w http.ResponseWriter, r *http.Request) {
        // Leia o arquivo JavaScript.
        data, err := ioutil.ReadFile("script.js")
        if err != nil {
            fmt.Fprintf(w, "Erro ao ler o arquivo JavaScript: %v", err)
            return
        }
        // Escreva o conteúdo do arquivo JavaScript na resposta HTTP.
        w.Header().Set("Content-Type", "application/javascript")
        w.Write(data)
    })

    mux.HandleFunc("/registros", pegaRegistrosDoBanco)     //GET /registros
    mux.HandleFunc("/inserir_registro", inserirRegistro)   //POST /inserir_registro
    mux.HandleFunc("/update_registro", editRegistro)       //POST /update_registro
    mux.HandleFunc("/delete", deleteRegistro)              //POST /delete
    mux.HandleFunc("/login", loginHandler)
    mux.HandleFunc("/dia-dia", indexHandler)
    mux.HandleFunc("/verifica-login", verificaLogin)
    http.ListenAndServe(":8080", handler)                       // Inicia o servidor na porta 8080
}


func pegaRegistrosDoBanco(w http.ResponseWriter, r *http.Request) {
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

    // Escrever o cabeçalho da resposta
    w.WriteHeader(http.StatusOK)

    // Definir o tipo de conteúdo da resposta como JSON
    w.Header().Set("Content-Type", "application/json")

    // Codificar os registros em JSON e escrever na resposta
    json, err := json.Marshal(registros)
    if err != nil {
        fmt.Fprintf(w, "Erro ao codificar os registros em JSON: %v", err)
        return
    }
    w.Write(json)
}

func inserirRegistro(w http.ResponseWriter, r *http.Request){
    // Decodificar o corpo da requisição
body, err := ioutil.ReadAll(r.Body)
if err != nil {
    fmt.Fprintf(w, "Erro ao ler o corpo da requisição: %v", err)
    return
}
// Criar um novo registro de dados diários
dados := DadosDiarios{}
err = json.Unmarshal(body, &dados)
if err != nil {
    fmt.Fprintf(w, "Erro ao decodificar o corpo da requisição: %v", err)
    return
}
// Retornar o status 201 Created
w.WriteHeader(http.StatusCreated)

db, err := sql.Open("postgres", "postgres://postgres:689df2c8@localhost:5432/Registros?sslmode=disable")
if err != nil {
    log.Fatal(err)
}
defer db.Close()

// Criar instrução SQL para atualizar o registro
stmt, err := db.Prepare("INSERT INTO diarios (data, lembrete, relatoDoDia, treino, aFazer, leujoje, ingles, tirouFoto, creatina, lendo) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);")
if err != nil {
    log.Fatal(err)
}
defer stmt.Close()

// Substituir os valores na instrução SQL
data := dados.Data
lembrete := dados.Lembrete
relatoDoDia := dados.RelatoDoDia
treino := dados.Treino
aFazer := dados.AFazer
leuHoje := dados.LeuHoje
ingles := dados.Ingles
tirouFoto := dados.TirouFoto
creatina := dados.Creatina
lendo := dados.Lendo

// // Executar a instrução SQL e verificar se a atualização foi bem-sucedida
_, err = stmt.Exec(data, lembrete, relatoDoDia, treino, aFazer, leuHoje, ingles, tirouFoto, creatina, lendo)

// fmt.Println(data, lembrete, relatoDoDia, treino, aFazer, leuHoje, ingles, tirouFoto, creatina, lendo)
if err != nil {
    log.Fatal(err)
}
fmt.Println("Registro inserido com sucesso!")
}

func editRegistro(w http.ResponseWriter, r *http.Request){
        // Decodificar o corpo da requisição
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Fprintf(w, "Erro ao ler o corpo da requisição: %v", err)
        return
    }

    // Criar um novo registro de dados diários
    dados := DadosDiarios{}
    err = json.Unmarshal(body, &dados)
    if err != nil {
        fmt.Fprintf(w, "Erro ao decodificar o corpo da requisição: %v", err)
        return
    }

    // Retornar o status 201 Created
    w.WriteHeader(http.StatusCreated)
    
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
    
    data := dados.Data
    lembrete := dados.Lembrete
    relatoDoDia := dados.RelatoDoDia
    treino := dados.Treino
    aFazer := dados.AFazer
    leuHoje := dados.LeuHoje
    ingles := dados.Ingles
    tirouFoto := dados.TirouFoto
    creatina := dados.Creatina
    lendo := dados.Lendo
    id := dados.ID
    
    // // Executar a instrução SQL e verificar se a atualização foi bem-sucedida
    _, err = stmt.Exec(data, lembrete, relatoDoDia, treino, aFazer, leuHoje, ingles, tirouFoto, creatina, lendo, id)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Registro atualizado com sucesso!")
}

func deleteRegistro(w http.ResponseWriter, r *http.Request){
    // Decodificar o corpo da requisição
body, err := ioutil.ReadAll(r.Body)
if err != nil {
    fmt.Fprintf(w, "Erro ao ler o corpo da requisição: %v", err)
    return
}

dados := DadosDiarios{}
err = json.Unmarshal(body, &dados)
if err != nil {
    fmt.Fprintf(w, "Erro ao decodificar o corpo da requisição: %v", err)
    return
}

w.WriteHeader(http.StatusNoContent)

db, err := sql.Open("postgres", "postgres://postgres:689df2c8@localhost:5432/Registros?sslmode=disable")
if err != nil {
    log.Fatal(err)
}
defer db.Close()

// Criar instrução SQL para atualizar o registro
stmt, err := db.Prepare("DELETE FROM diarios WHERE id = $1;")
if err != nil {
    log.Fatal(err)
}
defer stmt.Close()

// Substituir os valores na instrução SQL
id := dados.ID

// // Executar a instrução SQL e verificar se a atualização foi bem-sucedida
_, err = stmt.Exec(id)
if err != nil {
    log.Fatal(err)
    }
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    // Carregar o template
    t := template.Must(template.ParseFiles("html/template/login.html"))
  
    // Criar dados para o template
    data := Data{
      Mensagem: "Olá, por favor, faça login.",
    }
  
    // Renderizar o template
    t.Execute(w, data)
  }

func indexHandler(w http.ResponseWriter, r *http.Request) {
    // Carregar o template
    t := template.Must(template.ParseFiles("html/template/index.html"))
  
    // Renderizar o template
    t.Execute(w, nil)
  }

func verificaLogin(w http.ResponseWriter, r *http.Request) {
    // Decodificar o corpo da requisição
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Fprintf(w, "Erro ao ler o corpo da requisição: %v", err)
        return
    }
    
    type DadosUsuario struct {
        Usuario string`json:"usuario"`
        Password string`json:"password"`
}
    // Criar um novo registro de dados diários
    dados := DadosUsuario{}
    err = json.Unmarshal(body, &dados)
    if err != nil {
        fmt.Fprintf(w, "Erro ao decodificar o corpo da requisição: %v", err)
        return
    }
    // Abrir conexão com o banco de dados
    db, err := sql.Open("postgres", "postgres://postgres:689df2c8@localhost:5432/Registros?sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    // Criar instrução SQL para verificar se o usuário existe
    stmt, err := db.Prepare("SELECT EXISTS (SELECT 1 FROM usuarios WHERE usuario = $1 AND senha = $2);")
    if err != nil {
        log.Fatal(err)
    }
    defer stmt.Close()

    // Substituir os valores na instrução SQL
    usuario := dados.Usuario
    password := dados.Password

    // Executar a instrução SQL
    row := stmt.QueryRow(usuario, password)

    // Verificar o resultado da query
    var existe bool
    err = row.Scan(&existe)
    if err != nil {
        log.Fatal(err)
    }
    // Retornar o resultado para o usuário
    if existe {
        // Usuário existe
        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, `{"status": true, "mensagem": "Login realizado com sucesso!"}`)
    } else {
        // Usuário não existe
        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, `{"status": false, "mensagem": "Usuário e/ou senha incorretos."}`)
    }
}

