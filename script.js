function pegaDadosInput() {
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
  let verificacao = confirm(
    `Os dados abaixo serão inseridos na tabela\n
    ${data.value}\n
    ${lembrete.value}\n
    ${treino.value}\n
    ${afazer.value}\n
    ${leuhoje.value}\n
    ${ingles.value}\n
    ${tiroufoto.value}\n
    ${creatina.value}\n
    ${lendolivro.value}\n
    ${relato.value}\n`
  )

  if (verificacao) {
    let dados = [
      {
        ID: id.value,
        Data: data.value,
        Lembrete: lembrete.value,
        Relato: relato.value,
        Treino: treino.value,
        AFazer: afazer.value,
        LeuHoje: leuhoje.value,
        Ingles: ingles.value,
        TirouFoto: tiroufoto.value,
        Creatina: creatina.value,
        LendoLivro: lendolivro.value
      }
    ]

    try {
      insereDadosNoBanco({
        data: dados[0]['Data'],
        lembrete: dados[0]['Lembrete'],
        relato_do_dia: dados[0]['Relato'],
        treino: dados[0]['Treino'],
        a_fazer: dados[0]['AFazer'],
        leu_hoje: dados[0]['LeuHoje'],
        ingles: dados[0]['Ingles'],
        tirou_foto: dados[0]['TirouFoto'],
        creatina: dados[0]['Creatina'],
        lendo: dados[0]['LendoLivro']
      })
    } catch {
      console.error('Algo deu errado')
    } finally {
      setTimeout(() => window.location.reload(true), 500)
    }
  }
  limparDados()
}

function insereDadosNoBanco(dados) {
  axios
    .post('http://localhost:8080/inserir_registro', dados)
    .then(response => console.log(response))
    .catch(err => console.error(err))
}

function pegaDadosDaTr(id) {
  let inputs = []
  inputs.push(id)
  let dados = document.getElementById(id.toString()).children
  for (let i = 0; i < dados.length - 1; i++) {
    inputs.push(dados[i].innerHTML)
  }
  let valores = [
    {
      ID: inputs[0],
      Data: inputs[1],
      Lembrete: inputs[2],
      Relato: inputs[3],
      Treino: inputs[4],
      AFazer: inputs[5],
      LeuHoje: inputs[6],
      Ingles: inputs[7],
      TirouFoto: inputs[8],
      Creatina: inputs[9],
      LendoLivro: inputs[10]
    }
  ]
  console.log(valores)
  colocaDadosNoInput(valores)
}

function limparDados() {
  data.value = ''
  lembrete.value = ''
  treino.value = ''
  afazer.value = ''
  leuhoje.value = ''
  ingles.value = ''
  tiroufoto.value = ''
  creatina.value = ''
  lendolivro.value = ''
  relato.value = ''
}

function colocaDadosNoInput(dados) {
  let id = (document.querySelector('#ID').value = dados[0].ID)
  let data = (document.querySelector('#data').value = dados[0].Data)
  let lembrete = (document.querySelector('#lembrete').value = dados[0].Lembrete)
  let treino = (document.querySelector('#treino').value = dados[0].Treino)
  let afazer = (document.querySelector('#afazer').value = dados[0].AFazer)
  let leuhoje = (document.querySelector('#leuhoje').value = dados[0].LeuHoje)
  let ingles = (document.querySelector('#ingles').value = dados[0].Ingles)
  let tiroufoto = (document.querySelector('#tiroufoto').value =
    dados[0].TirouFoto)
  let creatina = (document.querySelector('#creatina').value = dados[0].Creatina)
  let lendolivro = (document.querySelector('#lendolivro').value =
    dados[0].LendoLivro)
  let relato = (document.querySelector('#relato').value = dados[0].Relato)

  document.getElementById('edicao').removeAttribute('hidden')
  document.getElementById('cadastrar').setAttribute('hidden', '')
}

function confirmarEdicao() {
  document.getElementById('cadastrar').removeAttribute('hidden')
  document.getElementById('edicao').setAttribute('hidden', '')

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

  //update where
  let dadosEditados = {
    data: data.value,
    lembrete: lembrete.value,
    relato_do_dia: relato.value,
    treino: treino.value,
    a_fazer: afazer.value,
    leu_hoje: leuhoje.value,
    ingles: ingles.value,
    tirou_foto: tiroufoto.value,
    creatina: creatina.value,
    lendo: lendolivro.value,
    id: parseInt(id.value)
  }
  console.log(dadosEditados)
  axios
    .post('http://localhost:8080/update_registro', dadosEditados)
    .then(async response => {
      console.log(await response)
      alert('Edição realizada com sucesso')
      setTimeout(() => window.location.reload(true), 1000)
    })
    .catch(err => console.error(err))
  //refresh da página
}

function criaDadosNoHtml(registros) {
  const tableBody = document.getElementById('table-body') // Replace with your table body element ID

  registros.forEach(registro => {
    const row = document.createElement('tr')
    row.setAttribute('id', registro.ID)

    const dataCell = document.createElement('td')
    dataCell.textContent = registro.Data
    const lembreteCell = document.createElement('td')
    lembreteCell.textContent = registro.Lembrete
    const treinoCell = document.createElement('td')
    treinoCell.textContent = registro.Treino
    const aFazerCell = document.createElement('td')
    aFazerCell.textContent = registro.AFazer
    const leuHojeCell = document.createElement('td')
    leuHojeCell.textContent = registro.LeuHoje
    const inglesCell = document.createElement('td')
    inglesCell.textContent = registro.Ingles
    const tirouFotoCell = document.createElement('td')
    tirouFotoCell.textContent = registro.TirouFoto
    const creatinaCell = document.createElement('td')
    creatinaCell.textContent = registro.Creatina
    const lendoLivroCell = document.createElement('td')
    lendoLivroCell.textContent = registro.Lendo
    const relatoCell = document.createElement('td')
    relatoCell.setAttribute("colspan","2")
    relatoCell.textContent = registro.RelatoDoDia

    const editButton = document.createElement('button')
    editButton.setAttribute('class', 'badge badge-secondary mt-4')
    editButton.setAttribute('onclick', `pegaDadosDaTr(${registro.ID})`)
    editButton.textContent = 'Editar'

    const deleteButton = document.createElement('button')
    deleteButton.setAttribute('class', 'badge badge-secondary mt-1')
    deleteButton.setAttribute('onclick', `deleteDadosDoBanco(${registro.ID})`)
    deleteButton.textContent = 'Deletar'

    row.appendChild(dataCell)
    row.appendChild(lembreteCell)
    row.appendChild(relatoCell)
    row.appendChild(treinoCell)
    row.appendChild(aFazerCell)
    row.appendChild(leuHojeCell)
    row.appendChild(inglesCell)
    row.appendChild(tirouFotoCell)
    row.appendChild(creatinaCell)
    row.appendChild(lendoLivroCell)
    row.appendChild(editButton)
    row.appendChild(deleteButton)

    tableBody.appendChild(row)
  })
}

function pegaDadosDoBanco() {
  axios
    .get('http://localhost:8080/registros')
    .then(async function (response) {
      let dados = await response
      criaDadosNoHtml(dados.data)
    })
    .catch(function (err) {
      console.error(err)
    })
}

function deleteDadosDoBanco(id) {
  console.log(`This is ID: ${id}`)
  let confirmarExclusao = confirm(
    'Tem certeza que deseja excluir esse registro?'
  )
  if (confirmarExclusao) {
    axios
      .post('http://localhost:8080/delete', {
        id: id
      })
      .then(response => console.log(response))
      .catch(err => console.error(err))
    setTimeout(() => window.location.reload(true), 500)
  } else {
    alert('Solicitação cancelada pelo usuário!')
  }
}

setTimeout(() => pegaDadosDoBanco(), 1000)
document.querySelector('#cadastrar').addEventListener('click', pegaDadosInput)
document.querySelector('#edicao').addEventListener('click', confirmarEdicao)
