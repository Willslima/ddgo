function handleData(){
    const username = document.getElementById("username").value
    const password = document.getElementById("password").value

    axios.post("http://localhost:8080/verifica-login", {"usuario": username, "password": password}).then(response => {
        // console.log(response.data)
        let spanText = document.getElementById("error")
        if(response.data.status){
        spanText.setAttribute("class","text-success")
        spanText.innerHTML = response.data.mensagem
        setTimeout(()=> {
            spanText.innerHTML = ""
            spanText.removeAttribute("class")
        }, 3000)
    } else {
        spanText.setAttribute("class","text-danger")
        spanText.innerHTML = response.data.mensagem
        setTimeout(()=> {
            spanText.innerHTML = ""
            spanText.removeAttribute("class")
        }, 3000)
    }
    }).catch(err => console.error(err))
}

document.getElementById("entrar").addEventListener("click", handleData)