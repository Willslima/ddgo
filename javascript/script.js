function pegaDadosInput(){
    let id = document.querySelector('#ID')
    let data = document.querySelector('#data')
    let lembrete = document.querySelector('#lembrete')
    let treino = document.querySelector('#treino') 
    let afazer = document.querySelector('#afazer') 
    let leuhoje = document.querySelector('#leuhoje') 
    let ingles = document.querySelector('#ingles') 
    let tiroufoto = document.querySelector('#tiroufoto') 
    let creatina = document.querySelector('#creatina') 
    let lendolivro = document.querySelector('#lendolivro') 
    let relato = document.querySelector('#relato') 

    alert(`Os dados abaixo serão inseridos na tabela\n${data.value}\n${lembrete.value}\n${treino.value}\n${afazer.value}\n${leuhoje.value}\n${ingles.value}\n${tiroufoto.value}\n${creatina.value}\n${lendolivro.value}\n${relato.value}\n`)

    criaDadosNoHtml(dados)
    limparDados()
}

function colocaDadosNoInput(id){
    let dados = document.getElementById(id.toString()).children
    for(let i = 0; i < dados.length; i++){
    console.log(dados[i].innerHTML)}
}

function limparDados(){
    data.value = ""
    lembrete.value = ""
    treino.value = ""
    afazer.value = ""
    leuhoje.value = ""
    ingles.value = ""
    tiroufoto.value = ""
    creatina.value = ""
    lendolivro.value = ""
    relato.value = ""
}

function criaDadosNoHtml(registros){
    const tableBody = document.getElementById('table-body'); // Replace with your table body element ID

    registros.forEach(registro => {
    const row = document.createElement('tr');
    row.setAttribute("id",registro.ID)

    const dataCell = document.createElement('td');
    dataCell.textContent = registro.Data;
    const lembreteCell = document.createElement('td');
    lembreteCell.textContent = registro.Lembrete;
    const treinoCell = document.createElement('td');
    treinoCell.textContent = registro.Treino;
    const aFazerCell = document.createElement('td');
    aFazerCell.textContent = registro.AFazer;
    const leuHojeCell = document.createElement('td');
    leuHojeCell.textContent = registro.LeuHoje;
    const inglesCell = document.createElement('td');
    inglesCell.textContent = registro.Ingles;
    const tirouFotoCell = document.createElement('td');
    tirouFotoCell.textContent = registro.TirouFoto;
    const creatinaCell = document.createElement('td');
    creatinaCell.textContent = registro.Creatina;
    const lendoLivroCell = document.createElement('td');
    lendoLivroCell.textContent = registro.Lendo;
    const relatoCell = document.createElement('td');
    relatoCell.textContent = registro.RelatoDoDia;

    const editLink = document.createElement('a');
    editLink.setAttribute("class","badge badge-secondary mt-4")
    editLink.href = `#`; // Replace with your edit link URL
    const image = document.createElement('img')
    image.setAttribute("src","../images/edit_icon.svg")

    editLink.append(image)

    row.appendChild(dataCell);
    row.appendChild(lembreteCell);
    row.appendChild(relatoCell);
    row.appendChild(treinoCell);
    row.appendChild(aFazerCell);
    row.appendChild(leuHojeCell);
    row.appendChild(inglesCell);
    row.appendChild(tirouFotoCell);
    row.appendChild(creatinaCell);
    row.appendChild(lendoLivroCell);
    row.appendChild(editLink);

    tableBody.appendChild(row);
});
}

function pegaDadosDoBanco(){
    axios.get("http://localhost:8080/registros")
        .then(async function(response){
            let dados = await response
            criaDadosNoHtml(dados.data)
        }).catch(function(err){
            console.error(err)
        })
}

setTimeout(() => pegaDadosDoBanco(), 1000)

document.querySelector("#cadastrar").addEventListener("click", pegaDadosInput)
