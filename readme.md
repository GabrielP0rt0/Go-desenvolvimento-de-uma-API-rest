<h1>Go: crie uma aplicação web V1.0</h1> 

<p align="center">
  
<img height="45" align="center" src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" />
<img height="45" align="center" src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/docker/docker-original.svg" />
          
</p>

> Status do Projeto: :warning: Em desenvolvimento

### Tópicos 

:small_blue_diamond: [Descrição do projeto](#descrição-do-projeto)

:small_blue_diamond: [Ultimas atualizações](#ultimas-atualizações)

:small_blue_diamond: [Funcionalidades](#funcionalidades) 

:small_blue_diamond: [Pré Requisitos](#pré-requisitos)

:small_blue_diamond: [Como rodar a aplicação](#como-rodar-a-aplicação-arrow_forward)

:small_blue_diamond: [Para testes](#para-testes)


... 

## Descrição do projeto 

<p align="justify">
  O Projeto foi criado como projeto do curso de mesmo nome oferecido pela <a href="https://cursos.alura.com.br/course/go-desenvolvendo-api-rest" > alura </a>. Este projeto tem como finalidade a criação de uma API rest utilizando a linguagem Go.
</p>

## Ultimas atualizações :new:
<p align="justify">
  O projeto atualmente se comunica com um banco de dados criado dentro de um container docker onde pode editar, criar, listar ou deletar itens de uma tabela.
</p>

## Funcionalidades

:heavy_check_mark: Busca todas as personalidades
- Através de uma requisição do tipo Get para /api/personalidades é possivel obter todas as personalidades cadastradas no banco de dados

:heavy_check_mark: Obter personalidade específica
- Através de uma requisição do tipo Get para /api/personalidades/{id}, onde "id" se refere a identificação única de cada personalidadem, podemos obter os dados de uma única personalidade

:heavy_check_mark: Criar uma nova personalidade
- Através de uma requisição do tipo Post para /api/personalidades pode-se criar um novo registro de personalidades no banco de dados. O body deve ser o seguinte:

```
{
    "nome":"nova personalidade",
    "historia":"nova historia"
}
```

:heavy_check_mark: Deletar uma personalidade
- Através de uma requisição do tipo Delete para /api/personalidades/{id} onde "id" se refere a identificação única de cada personalidadem, podemos deletar uma única personalidade.

:heavy_check_mark: Editar uma personalidade
- Através de uma requisição do tipo Put para /api/personalidades/{id} onde "id" se refere a identificação única de cada personalidadem, podemos editar uma personalidade. O body deve ser o seguinte:

```
{
    "nome":"nova personalidade editada",
    "historia":"nova historia editada"
}
```

## Pré-requisitos

:warning: [Go](https://medium.com/xp-inc/primeiros-passos-com-golang-1abdc60bba50)

:warning: [Docker](https://docs.docker.com/desktop/install/windows-install/)


## Como rodar a aplicação :arrow_forward:

No terminal, clone o projeto dentro da pasta de sua preferência utilizando o comando:

```
git clone https://github.com/GabrielP0rt0/Go-desenvolvimento-de-uma-API-rest
```

>obs: o fato da aplicação possuir go.mod impede a necessidade do código estar inserido na pasta src principal

Abra o terminal dentro deste diretório e utilize o seguinte comando:

```
docker-compose up
```

Após iniciado o servidor docker acesse localhost:54321/login e insira os dados de login presentes no arquivo "docker-compose.yml", dentro do servidor postgres crie um banco de dados com o nome Personalidades e insira informações para coneção conforme presentes no arquivo "docker-compose.yml"

Agora abra um novo terminal e digite:

```
go run main.go
```

Agora a aplicação encontra-se iniciada e pode ser testada.

## Para testes

Uma dica é utilizar o [Postman](https://www.postman.com/downloads/) para executar seus testes.

Por padrão a API inicia na porta 4200.

---


Copyright :copyright: 2023 - Go-desenvolvimento-de-uma-API-rest
