<div align="center">
  <h1>IoT Connect - TEC502: MI Concorrência e Conectividade</h1>

  ![Most used language](https://img.shields.io/github/languages/top/jnaraujo/tec502-iot-connect?style=flat-square)
  ![GitHub](https://img.shields.io/github/license/jnaraujo/tec502-iot-connect)
</div>

<p>
  O objetivo do IOT Connect é criar um sistema de comunicação e gerenciamento entre dispositivos IOT e suas diferentes aplicações. O sistema deve ser capaz de criar, remover e verificar o status de dispositivos IOT, além de receber, armazenar e disponibilizar dados dos dispositivos IOT para aplicações web.
</p>

## Arquitetura do projeto

### Tecnologias utilizadas
- Client
  - [React](https://reactjs.org/): Biblioteca JavaScript para a criação de interfaces de usuário.
  - [Vite](https://vitejs.dev/): Ferramenta de build para aplicações web.
  - [TypeScript](https://www.typescriptlang.org/): Superset da linguagem JavaScript que adiciona tipagem estática ao código.
  - [TanStack Query](https://tanstack.com/query/latest): Biblioteca para gerenciamento de estado e requisições HTTP. Responsável por fazer a comunicação com o Broker.
- Broker
  - [Go](https://golang.org/): Linguagem de programação utilizada para o desenvolvimento do Broker.
  - [Gin](https://gin-gonic.com/) Framework web utilizado para a criação das rotas HTTP e API REST.
- Sensores
  - [Python](https://www.python.org/): Linguagem de programação utilizada para o desenvolvimento dos Sensores.

### Pré-requisitos
O sistema foi desenvolvido utilizando Docker e Docker Compose. Assim, para executar o projeto, é necessário ter o Docker e o Docker Compose instalados na máquina.

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

### Como executar
1. Clone o repositório:

```bash
git clone https://github.com/jnaraujo/tec502-iot-connect
```

2. Entre na pasta do projeto:

```bash
cd tec502-iot-connect
```

3. Execute o comando:

```bash
docker-compose up --build
```

4. Acesse a aplicação em [http://localhost:3000](http://localhost:3000)

### Como utilizar
#### Como adicionar um novo sensor
<div align="center">
<img src="./docs/imgs/adicionar-novo-sensor.png" alt="Adicionar Sensor" height="300px" width="auto" />
</div>

Para adicionar um novo sensor, clique no botão "+" na caixa de "Lista de Sensores". Digite um ID único para o sensor e o endereço IP do sensor (ex: 172.19.0.2:3333 ou 172.19.0.4:3334). Clique em "Adicionar Sensor". Uma mensagem irá aparecer informando se o sensor foi adicionado com sucesso ou se houve algum erro.

#### Como enviar um comando para um sensor
<div align="center">
<img src="./docs/imgs/enviar-comando.png" alt="Enviar Comando" height="300px" width="auto" />
</div>

Para enviar um comando para um sensor, na caixa "Enviar comando", selecione o id do sensor em "Sensor ID", selecione o comando que deseja enviar em "Comando", escreva o conteúdo do comando (se necessário) e clique em "Enviar Comando". Uma mensagem irá aparecer informando se o comando foi enviado com sucesso ou se houve algum erro.

#### Como visualizar os dados de um sensor
<div align="center">
<img src="./docs/imgs/resposta-dos-sensores.png" alt="Dados Sensor" height="300px" width="auto" />
</div>

Na caixa "Respostas dos sensores" irá aparecer todos os dados recebidos dos sensores, bem como o ID do sensor que enviou o dado, qual o comando, o conteúdo do dado, um histórico de envio e a data de envio.

#### Como remover um sensor
<div align="center">
<img src="./docs/imgs/remover-sensor.png" alt="Remover Sensor" height="100px" width="auto" />
</div>

Para remover um sensor, na caixa de "Lista de sensores", clique no ícone de lixeira ao lado do sensor que deseja remover. Um modal irá aparecer perguntando se você realmente deseja remover o sensor. Clique em "Remover" para confirmar a remoção.

### Estrutura do projeto
As principais pastas do projeto são:

- `client`: Aplicação web desenvolvida em React.
- `broker`: Serviço de mensageria desenvolvido em Go.
- `sensor`: Simulação de dispositivos IOT desenvolvida em Python.

#### Client
Dentro da pasta `client`, temos o código da aplicação web desenvolvida em React. A aplicação é responsável por criar, remover e visualizar dispositivos IOT, além de visualizar os dados enviados pelos dispositivos.

```bash
client
├── src # Código fonte da aplicação
│   ├── components # Componentes React
│       │── sensor-list-box # Componente que exibe a lista de sensores
│       │── sensor-response-list-box # Componente que exibe a lista de respostas dos sensores
│       │── ui # Componentes genéricos, como botões e inputs
│       │── chart.tsx # Componente que exibe um gráfico de linhas
│       │── send-command-box.tsx # Componente que exibe o formulário para envio de comandos
│   ├── constants # Constantes utilizadas na aplicação
│   ├── hooks # Funções para gerenciamento de estado
│       │── use-command-list.ts # Responsável por buscar os comandos disponíveis para um determinado sensor
│       │── use-create-sensor.ts # Responsável por criar um novo sensor
│       │── use-delete-sensor.ts # Responsável por deletar um sensor
│       │── use-send-command.ts # Responsável por enviar um comando para um sensor
│       │── use-sensor-list.ts # Responsável por buscar a lista de sensores
│       │── use-sensor-responses.ts # Responsável por buscar as respostas (dados) de um sensor
│   ├── lib # Funções utilitárias
│   ├── routes # Rotas da aplicação
│   ├── env.ts # Variáveis de ambiente
│   ├── index.css # Estilos globais
│   ├── main.tsx # Ponto de entrada da aplicação
```

#### Broker
Dentro da pasta `broker`, temos o código do serviço de mensageria desenvolvido em Go. O Broker é responsável por permitir a troca de mensagens entre os dispositivos e as aplicações que precisam desses dados. O Broker é responsável por receber os dados dos Sensores, armazenar esses dados e disponibilizá-los para as aplicações.

```bash
broker
├── cmd # Comandos para execução do Broker
│   │── api/main.go # Ponto de entrada da API REST
├── internal # Código interno do Broker
│   │── cmd # Protocolo para simplificar de comunicação entre o Broker e os Sensores
│   │── constants # Constantes utilizadas no Broker
│   │── http # Código para criação das rotas HTTP
│       │── routes # Rotas da API REST
│       │── http.go # Configuração do servidor HTTP
│   │── queue # Estrutura de dados para armazenar os dados dos Sensores
│   │── sensor_conn # Responsável por gerenciar a conexão com os Sensores via TCP/IP
│   │── storage # Responsável por armazenar os dados dos Sensores
│   │── time # Funções utilitárias para manipulação de tempo
│   │── udp_server # Responsável por receber os dados dos Sensores via UDP
│   │── utils # Funções utilitárias
```

#### Sensores
Dentro da pasta `sensor`, temos o código dos Sensores, que são responsáveis por simular dispositivos IOT que enviam dados para o Broker. Os Sensores são desenvolvidos em Python.

```bash
sensor
├── libs # Bibliotecas utilizadas
│   │── broker_service.py # Classe para comunicação com o Broker via UDP
│   │── cmd_data.py # Protocolo para simplificar de comunicação entre o Broker e os Sensores
│   │── interface.py # Interface que permite gerenciar o sensor, como enviar dados e comandos.
│   │── server.py # Responsável por criar um servidor TCP/IP para receber comandos do Broker.
│   │── utils.py # Funções utilitárias
├── air_cond.main.py # Sensor que simula um ar condicionado
├── lamp.main.py # Sensor que simula uma lâmpada
```

## Comunicação
### Protocolo de comunicação
Para a comunicação entre os dispositivos IOT e o Broker, foi criado um protocolo de comunicação. O protocolo é baseado texto e funciona tanto sobre TCP/IP quanto UDP. O protocolo é composto por 3 partes: `IdFrom`, `IdTo`, `Cmd` e o conteúdo. Todos os campos são obrigatórios e devem ser enviados em ordem. O conteúdo é opcional e depende do comando que está sendo enviado. O protocolo é sempre enviado em texto e é finalizado com uma quebra de linha (`\n`).

O protocolo segue o seguinte formato:

```txt
IdFrom: <id>
IdTo: <id>
Cmd: <comando>

<conteúdo (opcional)>
```

#### Exemplo de protocolo
```txt
IdFrom: temp1
IdTo: BROKER
Cmd: set_temp

25
```

Nesse exemplo, o Sensor `temp1` está enviando um comando para o Broker. O comando é `set_temp` e o conteúdo é `25`.

#### Campos do protocolo
##### idFrom
ID do dispositivo que está enviando a mensagem. No caso do Broker, o `idFrom` é sempre `BROKER` (um exemplo pode ser visto na <a href="https://github.com/jnaraujo/tec502-iot-connect/blob/main/broker/internal/http/routes/create-sensor.go#L62">rota de criação de um novo sensor</a>). No caso dos Sensores, o `idFrom` é o ID do Sensor. Por exemplo, se um sensor for criado com o ID `temp1`, o `idFrom` será sempre `temp1`.

##### IdTo
ID do dispositivo que irá receber a mensagem. Por exemplo, se o Broker quiser enviar um comando para o Sensor `temp1`, o `idTo` será `temp1`. Se um Sensor quiser enviar dados para o Broker, o `idTo` será sempre `BROKER`.

##### Cmd
Comando que está sendo enviado. O comando pode ser `turn_on`, `turn_off`, `set_temp`, entre outros. O comando é sempre uma string.

##### Conteúdo
Qualquer dado que será enviado junto ao comando, como a temperatura que o Sensor está enviando, por exemplo. O conteúdo é enviado como string (embora possa ser convertido para o tipo desejado no Broker) e é opcional.


### Comunicação entre Client e Broker
A comunicação entre o Client e o Broker é feita através de HTTP (REST API). Essa api permite a criação, remoção, visualização de dispositivos IOT e a visualização dos dados enviados pelos dispositivos.

#### Rotas
Todas as rotas descritas abaixo estão disponíveis em `/broker/internal/http`.

##### GET /
Rota que retorna a página inicial da aplicação.
> O arquivo responsável por essa rota pode ser encontrado em `broker/internal/http/routes/get-root.go`.

##### POST /message
Rota que envia um comando para um Sensor. O corpo da requisição deve conter o `sensorId`, o `command` e o `content`.
> O arquivo responsável por essa rota pode ser encontrado em `broker/internal/http/routes/post-message.go`.

##### POST /sensor
Rota que cria um novo sensor. O corpo da requisição deve conter o `sensorId`, o `command` e o `content`.
> O arquivo responsável por essa rota pode ser encontrado em `broker/internal/http/routes/create-sensor.go`.

##### GET /sensor
Rota que retorna a lista de sensores cadastrados.
> O arquivo responsável por essa rota pode ser encontrado em `broker/internal/http/routes/find-all-sensors.go`.

##### GET /sensor/commands/:sensor_id
Rota que retorna a lista de comandos disponíveis para um determinado Sensor.
> O arquivo responsável por essa rota pode ser encontrado em `broker/internal/http/routes/find-sensor-commands.go`.

##### GET /sensor/data
Rota que retorna a lista de dados enviados pelos Sensores.
> O arquivo responsável por essa rota pode ser encontrado em `broker/internal/http/routes/find-all-sensor-data.go`.

##### DELETE /sensor/:sensor_id
Rota que remove um Sensor.
> O arquivo responsável por essa rota pode ser encontrado em `broker/internal/http/routes/delete-sensor.go`.

### Comunicação entre Broker e Sensores
A comunicação entre o Broker e os Sensores é feita através de TCP/IP e UDP. Para envio de comandos do Broker para os Sensores, é utilizando uma abordagem confiável (TCP/IP). Assim, sempre que o Broker precisa enviar algum dado para os Sensores, ele inicia uma conexão TCP/IP com o Sensor e envia o comando. 

Para o envio de dados dos Sensores para o Broker, é utilizado uma abordagem não confiável (UDP). Assim, os Sensores enviam os dados para o Broker através de pacotes UDP. A abordagem não confiável foi escolhida pois o protocolo UDP é mais leve e mais rápido que o TCP. Desse modo, como os sensores estão constantemente enviando dados para o Broker, caso algum pacote seja perdido, o Sensor irá enviar novamente na próxima iteração, não causando maiores problemas ao sistema.

Todas essas comunicações são feitas através do protocolo de comunicação [descrito acima](#protocolo-de-comunicação).

## Tolerância a falhas
Para um sistema de troca de mensagens ser tolerante a falhas, é necessário que ele seja capaz de lidar com falhas de comunicação entre os dispositivos. Para isso, algumas estratégias foram adotadas.

Para a troca de mensagens entre o Client e o Broker, foi utilizado o protocolo HTTP. Esse protocolo é confiável, ou seja, ele garante que a mensagem enviada pelo Client será entregue ao Broker. Assim, caso ocorra algum problema na comunicação, o usuário é sempre avisado.

O envio de comandos entre o Broker e os Sensores é feito através de uma abordagem confiável (TCP/IP). Assim, caso ocorra algum problema na comunicação, o Broker informará ao usuário que houve uma falha na comunicação.

O envio de dados dos Sensores para o Broker é feito através de uma abordagem não confiável (UDP). Assim, caso ocorra algum problema na comunicação, o Sensor irá enviar novamente na próxima iteração, não causando maiores problemas ao sistema. No momento em que a conexão é estabelecida, os dados voltam a chegar ao Broker, que irá processá-los normalmente. Vale destacar que, caso o Broker não receba os dados de um Sensor por um determinado tempo, ele considerará o Sensor como offline.

