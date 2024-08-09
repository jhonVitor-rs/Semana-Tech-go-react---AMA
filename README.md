# Rocketseat Tech Week: Go & React Project

## Descrição do Projeto

Este projeto foi desenvolvido durante a "Semana Tech Go e React" da [Rocketseat](https://app.rocketseat.com.br/). Ele consiste em uma aplicação web que permite a criação de salas de discussão, onde os usuários podem enviar perguntas em tempo real. A aplicação é construída com um backend em Go, utilizando WebSocket para comunicação em tempo real, um frontend em Vite com React, e um banco de dados PostgreSQL para armazenar as informações das salas e perguntas.

## Funcionalidades

- **Criação de Salas:** Os usuários podem criar suas próprias salas de discussão.
- **Envio de Perguntas:** Dentro de uma sala, os participantes podem enviar perguntas em tempo real.
- **Compartilhamento de Salas:** As salas podem ser compartilhadas via link para que outras pessoas possam participar.
- **Comunicação em Tempo Real:** Utiliza WebSocket para garantir que as perguntas apareçam instantaneamente para todos os participantes da sala.

## Tecnologias Utilizadas

- **Go:** Linguagem de programação utilizada para construir o backend e gerenciar a comunicação WebSocket.
- **Vite:** Ferramenta de desenvolvimento rápido utilizada para criar o frontend com React.
- **React:** Biblioteca JavaScript para construção de interfaces de usuário.
- **PostgreSQL:** Banco de dados relacional utilizado para armazenar informações das salas e perguntas.

## Rodando o Projeto

- **Certifique-se de ter o Docker instalado em sua máquina.** Para instruções detalhadas sobre como instalar o Docker, consulte a [documentação oficial do Docker](https://docs.docker.com/get-docker/).

- **Clone este repositório:**

  ```bash
  git clone <url-do-repositorio>
  ```

- **Navegue até o diretório do projeto:**

  ```bash
  cd <diretorio-do-projeto>
  ```

- **Construa e inicie os containers:**

  ```bash
  docker-compose up --build
  ```

- **Acesse a aplicação em:**
  - **_Frontend_**: http://localhost:5173
