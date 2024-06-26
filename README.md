<div align="center">
  <h1>IoT Connect - Gerenciamento de dispositivos IoT</h1>
  <p>
    <strong>Projeto desenvolvido para a disciplina TEC502 - MI Concorrência e Conectividade</strong>
  </p>

  ![Most used language](https://img.shields.io/github/languages/top/jnaraujo/tec502-iot-connect?style=flat-square)
  ![GitHub](https://img.shields.io/github/license/jnaraujo/tec502-iot-connect)
</div>

<div align="center">
  <img src="./docs/gif/client.gif" alt="Client web" height="400px" width="auto" />
</div>

<br />

No cenário atual, a crescente necessidade de integração de dispositivos IoT (Internet das Coisas) demanda sistemas robustos de gerenciamento e comunicação para garantir sua eficácia e confiabilidade. O projeto Iot Connect foi criado com o objetivo de permitir a interação entre dispositivos IoT e suas aplicações, facilitando a criação, remoção e monitoramento de dispositivos, bem como o recebimento, armazenamento e disponibilização de dados gerados por esses dispositivos os usuários.

O IoT Connect faz uso de protocolos de comunicação confiáveis e não confiáveis, como TCP/IP e UDP, para garantir a comunicação entre os dispositivos IoT e o Broker, responsável por gerenciar os dispositivos e os dados gerados por eles. Essa variedade de protocolos permite que o IoT Connect adapte-se de forma flexível às diferentes necessidades de comunicação, oferecendo uma abordagem versátil para lidar com diversos tipos de dispositivos e ambientes de rede. O Broker, como peça central do sistema, desempenha um papel crucial no gerenciamento dos dispositivos e dos dados que fluem entre eles, garantindo não apenas a eficiência operacional, mas também a segurança e integridade das informações transmitidas.

O projeto foi desenvolvido utilizando tecnologias modernas e práticas, como Docker, React, Go e Python, que permitem a criação de aplicações escaláveis e de alto desempenho. A arquitetura do sistema foi projetada para ser modular e extensível, facilitando a adição de novos dispositivos e funcionalidades, bem como a integração com outros sistemas e serviços. Este texto explorará em detalhes as tecnologias utilizadas, a arquitetura do projeto, os protocolos de comunicação adotados, a confiabilidade da solução e sua tolerância a falhas, oferecendo uma visão abrangente de como o IoT Connect aborda os desafios do gerenciamento de dispositivos IoT.

## Sumário
- [Sobre o projeto](#sobre-o-projeto)
  - [Tecnologias utilizadas](#tecnologias-utilizadas)
  - [Como executar](#como-executar)
  - [Como utilizar](#como-utilizar)
    - [Como adicionar um novo sensor](#como-adicionar-um-novo-sensor)
    - [Como enviar um comando para um sensor](#como-enviar-um-comando-para-um-sensor)
    - [Como visualizar os dados de um sensor](#como-visualizar-os-dados-de-um-sensor)
    - [Como remover um sensor](#como-remover-um-sensor)
    - [Envio de comandos pelo terminal do Sensor](#envio-de-comandos-pelo-terminal-do-sensor)
- [Arquitetura do projeto](#arquitetura-do-projeto)
  - [Client](#client)
  - [Broker](#broker)
  - [Sensores](#sensores)
- [Comunicação](#comunicação)
  - [Protocolo de comunicação](#protocolo-de-comunicação)
  - [Comunicação entre Client e Broker](#comunicação-entre-client-e-broker)
    - [Rotas da API REST](#rotas-da-api-rest)
  - [Comunicação entre Broker e Sensores](#comunicação-entre-broker-e-sensores)
    - [Envio de comandos do Broker para o Sensor](#envio-de-comandos-do-broker-para-o-sensor)
    - [Comandos disponíveis](#comandos-disponíveis)
    - [Envio de dados dos Sensores para o Broker](#envio-de-dados-dos-sensores-para-o-broker)
      - [Lidando com concorrência](#lidando-com-concorrência)
- [Confiabilidade da solução e tolerância a falhas](#confiabilidade-da-solução-e-tolerância-a-falhas)
- [Testes](#testes)
- [Conclusão](#conclusão)
  

## Sobre o projeto
### Tecnologias utilizadas
- Geral
  - [Docker](https://www.docker.com/): Plataforma de código aberto para criação, execução e gerenciamento de aplicações em containers.
  - [Docker Compose](https://docs.docker.com/compose/): Ferramenta para definir e executar aplicações Docker em múltiplos containers.
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

> Nota: Por uma limitação do Docker Compose, não é possível usar a interface do sensor através do Docker Compose. Para saber como acessar a interface do sensor, veja a seção [Envio de comandos pelo terminal do Sensor](#envio-de-comandos-pelo-terminal-do-sensor).

### Como utilizar
#### Como adicionar um novo sensor
<div align="center">
<img src="./docs/imgs/adicionar-novo-sensor.png" alt="Adicionar Sensor" height="300px" width="auto" /> <br/>
<em>Figura 1. Modal para adicionar um novo sensor</em> <br/>
</div>

Para adicionar um novo sensor, clique no botão "+" na caixa de "Lista de Sensores". Digite um ID único para o sensor e o endereço IP do sensor (ex: 172.19.0.2:3333 ou 172.19.0.4:3334). Clique em "Adicionar Sensor". Uma mensagem irá aparecer informando se o sensor foi adicionado com sucesso ou se houve algum erro.

#### Como enviar um comando para um sensor
<div align="center">
<img src="./docs/imgs/enviar-comando.png" alt="Enviar Comando" height="300px" width="auto" /> <br/>
<em>Figura 2. Caixa para enviar um comando para um sensor</em> <br/>
</div>

Para enviar um comando para um sensor, na caixa "Enviar comando", selecione o id do sensor em "Sensor ID", selecione o comando que deseja enviar em "Comando", escreva o conteúdo do comando (se necessário) e clique em "Enviar Comando". Uma mensagem irá aparecer informando se o comando foi enviado com sucesso ou se houve algum erro.

#### Como visualizar os dados de um sensor
<div align="center">
<img src="./docs/imgs/resposta-dos-sensores.png" alt="Dados Sensor" height="300px" width="auto" /> <br/>
<em>Figura 3. Caixa para visualizar os dados de um sensor</em> <br/>
</div>

Na caixa "Respostas dos sensores" serão exibidos todos os dados recebidos dos sensores, bem como o ID do sensor que enviou o dado, qual o comando, o conteúdo do dado, um histórico de envio e a data de envio.

#### Como remover um sensor
<div align="center">
<img src="./docs/imgs/remover-sensor.png" alt="Remover Sensor" height="100px" width="auto" /> <br/>
<em>Figura 4. Remover um sensor</em> <br/>
</div>

Para remover um sensor, na caixa de "Lista de sensores", clique no ícone de lixeira ao lado do sensor que deseja remover. Um modal irá aparecer perguntando se você realmente deseja remover o sensor. Clique em "Remover" para confirmar a remoção.

#### Envio de comandos pelo terminal do Sensor
<div align="center">
<img src="./docs/imgs/terminal-sensor.png" alt="Interface do Sensor" height="300px" width="auto" /> <br/>
<em>Figura 5. Interface do Sensor</em> <br/>
</div>

O sensor possui um terminal que permite configurar o sensor e enviar comandos para o Broker. Pelo terminal, é possível cadastrar o Sensor, atualizar os dados do Sensor e executar comandos no Sensor. Para isso, basta escrever o comando no terminal e pressionar Enter. O Sensor executa o comando e retornar a resposta. O terminal do Sensor é útil para testar a comunicação entre o Sensor e o Broker.

Devido a uma limitação do Docker Compose, a interface do Sensor não pode ser acessada por meio dele. Para acessá-la, é necessário executar o Dockerfile de forma manual. Para isso, basta entrar na pasta do Sensor e executar:
  
```bash
docker build -t sensor -f <dockerfile> . && docker run  -p <porta>:3333 -e BROKER_URL=<broker url> -it sensor
```

O `<dockerfile>` é o Dockerfile do Sensor que deseja executar (AirCond.Dockerfile ou Lamp.Dockerfile). A `<porta>` é a porta que deseja mapear para acessar a interface do Sensor. O `<broker url>` é a URL do Broker.

Um exemplo de execução do Sensor pode ser visto abaixo:

```bash
docker build -t sensor -f AirCond.Dockerfile . && docker run  -p 3344:3333 -e BROKER_URL=localhost:5310 -it sensor
```

## Arquitetura do projeto
As principais pastas do projeto são:

- `client`: Aplicação web desenvolvida em React.
- `broker`: Serviço de mensageria desenvolvido em Go.
- `sensor`: Simulação de dispositivos IoT desenvolvida em Python.

### Client
Dentro da pasta `client`, temos o código da aplicação web desenvolvida em React. A aplicação é responsável por criar, remover e visualizar dispositivos IoT, além de visualizar os dados enviados pelos dispositivos.

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

### Broker
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

### Sensores
Dentro da pasta `sensor`, temos o código dos Sensores, que são responsáveis por simular dispositivos IoT que enviam dados para o Broker. Os Sensores são desenvolvidos em Python.

Foram criados dois Sensores: um que simula um ar condicionado e outro que simula uma lâmpada. Os Sensores são responsáveis por enviar dados para o Broker e receber comandos do Broker.

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
Para a comunicação entre os dispositivos IoT e o Broker, foi criado um protocolo de comunicação. O protocolo é baseado texto e funciona tanto sobre TCP/IP quanto UDP. O uso do protocolo facilita a comunicação entre os dispositivos e permite a criação de novos comandos de forma simples.

A implementação do protocolo foi feita tanto no [Broker](https://github.com/jnaraujo/tec502-iot-connect/broker/internal/cmd/cmd.go) quanto nos [Sensores](https://github.com/jnaraujo/tec502-iot-connect/sensor/libs/cmd_data.py). Assim, no envio de comandos do Broker para os Sensores, o Broker envia o comando no formato do protocolo sobre TCP/IP. Já no envio de dados dos Sensores para o Broker, os Sensores enviam os dados no formato do protocolo sobre UDP.

#### Formato do protocolo
O protocolo é composto por 4 partes: `IdFrom`, `IdTo`, `Cmd` e o conteúdo. Os três primeiros campos são obrigatórios e devem ser enviados em ordem. O conteúdo é opcional e depende do comando que está sendo enviado. O protocolo é sempre enviado em texto e cada campo é separado por uma quebra de linha (`\n`), com exceção do conteúdo, que é separado por duas quebras de linha (`\n\n`).

O protocolo segue o seguinte formato:

```txt
IdFrom: <id>
IdTo: <id>
Cmd: <comando>

<conteúdo (opcional)>
```

Exemplo de protocolo:

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
<div align="center">
<img src="./docs/imgs/diagrama-cliente-dados.png" alt="Pedido de dados do Client para o Broker" height="300px" width="auto" /> <br/>
<em>Figura 6. Diagrama de pedido de dados do Client para o Broker</em> <br/>
</div>

A comunicação entre o Client e o Broker é feita através de **HTTP**, utilizando o padrão de API REST. A api do Broker permite a criação, remoção, visualização de dispositivos IoT e a visualização dos dados enviados pelos dispositivos. Ao se iniciar o Broker, o acesso a API poderá ser feito através do endereço `http://<endereço>:8080`.

No Broker, foi utilizado a biblioteca [Gin](https://gin-gonic.com/) para a criação das rotas HTTP. Já no Client, foi utilizado a biblioteca [TanStack Query](https://tanstack.com/query/latest) para fazer as requisições HTTP e gerenciar o estado da aplicação. O TanStack Query permite definir um tempo de refetch, ou seja, a cada X segundos, a aplicação irá buscar os dados novamente, garantindo que a aplicação esteja sempre atualizada. O código do servidor pode ser encontrado em [broker/cmd/api/main.go](https://github.com/jnaraujo/tec502-iot-connect/tree/main/broker/internal/http), enquanto o código do cliente pode ser encontrado em [client/src/hooks](https://github.com/jnaraujo/tec502-iot-connect/tree/main/client/src/hooks).

#### Rotas da API REST
As rotas definem os pontos de entrada da API REST, cada uma correspondendo a uma operação específica que pode ser realizada no sistema. Seguindo o padrão de API REST, as rotas são projetadas para serem intuitivas e autoexplicativas, facilitando a compreensão e o uso por parte dos desenvolvedores.

##### GET /
Esta rota retorna a página inicial da aplicação. Ela é utilizada somente para fins de teste e não é utilizada na aplicação em si.
> O arquivo responsável por essa rota pode ser encontrado em `broker/internal/http/routes/get-root.go`.

##### POST /message
Esta [rota](broker/internal/http/routes/post-message.go) é usada para enviar um comando para um sensor específico. O corpo da requisição deve conter o `sensorId`, o `command` e o `content`. O `content` é opcional e depende do tipo de comando que está sendo enviado, como a temperatura que o Sensor deve ser configurado, por exemplo.

Exemplo de requisição:
```http
POST /message
{
  "sensor_id": "ar_condicionado",
  "command": "set_temp",
  "content": "20"
}
```

Exemplo de resposta:
```http
200 OK
{
  "message": "Temperature set to 20.0"
}
```

##### POST /sensor
Esta [rota](broker/internal/http/routes/create-sensor.go) cria um novo sensor no sistema. O corpo da requisição deve conter informações sobre o sensor, como sensorId, name, type, entre outros, dependendo dos requisitos do sistema.

Exemplo de requisição:
```http
POST /sensor
{
  "address": "localhost:3334",
  "id": "temp1"
}
```

Exemplo de resposta:
```http
201 Created
{
  "message": "Sensor criado"
}
```

##### GET /sensor
Esta [rota](broker/internal/http/routes/find-all-sensors.go) retorna uma lista de todos os sensores cadastrados no sistema. Isso permite que os clientes obtenham uma visão geral dos sensores disponíveis.

Exemplo de resposta:
```http
200 OK
{
  "sensors": [
    {
      "id": "temp1",
      "address": "localhost:3334",
      "is_online": true
    },
    {
      "id": "ar_condicionado",
      "address": "localhost:3335",
      "is_online": false
    }
  ]
}
```

##### GET /sensor/commands/:sensor_id
Esta [rota](broker/internal/http/routes/find-sensor-commands.go) retorna uma lista de comandos disponíveis para um determinado sensor. Os comandos podem variar dependendo do tipo de sensor e das funcionalidades suportadas.

Exemplo de resposta:
```http
200 OK
{
  "commands": [
    "not_found",
    "set_temp",
    "turn_on",
    "turn_off"
  ]
}
```

##### GET /sensor/data
Esta [rota](broker/internal/http/routes/find-all-sensor-data.go) retorna uma lista de dados enviados pelos sensores. Isso permite que os clientes visualizem os dados coletados ao longo do tempo. O `content` é uma lista dos últimos 20 dados enviados pelo Sensor.

Exemplo de resposta:
```http
200 OK
[
  {
    "sensor_id": "temp1",
    "name": "temperature",
    "content": [
      21.15,
      20.95,
      21.13,
      20.86
    ],
    "created_at": "2024-05-03T22:25:27.905472828Z",
    "updated_at": "2024-05-03T22:30:40.033540497Z"
  }
]
```

##### DELETE /sensor/:sensor_id
Esta [rota](broker/internal/http/routes/delete-sensor.go) remove um sensor específico do sistema. Isso pode ser útil para gerenciar sensores defeituosos ou desativados.

Exemplo de resposta:
```http
200 OK
{
  "message": "Sensor deletado."
}
```

### Comunicação entre Broker e Sensores
<div align="center">
<img src="./docs/imgs/diagrama-envio-comandos.png" alt="Envio de comandos do Client para o Broker" height="300px" width="auto" /> <br/>
<em>Figura 7. Diagrama de envio de comandos do Client para o Broker</em> <br/>
</div>

A comunicação entre o Broker e os Sensores é feita tanto com TCP/IP quanto com UDP. O Broker é responsável por enviar comandos para os Sensores e receber os dados enviados pelos Sensores. Todas essas comunicações são feitas utilizando o [protocolo de comunicação descrito acima](#protocolo-de-comunicação).

#### Envio de comandos do Broker para o Sensor
Para envio de comandos do Broker para os Sensores, é utilizando uma abordagem confiável (TCP/IP). Assim, sempre que o Broker precisa enviar algum dado para os Sensores, ele inicia uma conexão TCP/IP com o Sensor e envia o comando. O fato da comunicação ser confiável garante que o comando será entregue ao Sensor, ou o Broker será informado caso ocorra algum problema na comunicação.

Ao se iniciar um novo Sensor, um socket é aberto na porta 3333 (desse modo, cabe ao desenvolvedor mapear esse porta através do Docker) e o Sensor fica aguardando por novas conexões. Sempre que o Broker precisa enviar um comando para um Sensor, ele [estabelece uma conexão TCP/IP](https://github.com/jnaraujo/tec502-iot-connect/blob/main/broker/internal/sensor_conn/sensor.go#L21-L22) com o Sensor. Antes de enviar qualquer dado, é realizado um [handshake](https://github.com/jnaraujo/tec502-iot-connect/broker/internal/sensor_conn/sensor.go#L59) para garantir que a conexão foi estabelecida corretamente. Após o handshake, o Broker envia o comando para o Sensor, que executa a ação correspondente.

Para saber qual comando deve ser executado, o Sensor verifica se o comando recebido existe na lista de comandos disponíveis. Caso o comando exista, o Sensor executa a ação correspondente. Caso contrário, o Sensor retorna um comando de erro. Por exemplo, o código abaixo mostra como o Sensor lida com a chegada de um novo comando:
```python
# Código retirado de: /sensor/libs/server.py
if command not in self.commands: # Se o comando não existir, retorna um comando de erro
  # Código para retornar um comando de erro
else:
  resCmd = self.commands[command](data) # Executa o comando
```

Para lidar com o recebimento dos comandos, o Sensor permite ao desenvolvedor [criar](https://github.com/jnaraujo/tec502-iot-connect/sensor/air_cond.main.py#L57) os próprios comandos e [cadastrar](https://github.com/jnaraujo/tec502-iot-connect/sensor/libs/server.py#L60C7-L60C21) no `Server`. Para isso, ele define o nome do comando e a função que será executada quando o comando for recebido. Essa abordagem torna mais fácil criar novos comandos e adicionar novas funcionalidades ao Sensor.

Por exemplo, o código abaixo mostra como o Sensor de ar condicionado lida com o comando `set_temp`:
```python
# Código retirado de: /sensor/air_cond.main.py
def set_temp_cmd(cmd: cmd_data.Cmd):
  if not STATUS: # Se o sensor estiver desligado, retorna um erro
    return cmd_data.BasicCmd("error", "O sensor está desligado")

  try:
    cmd.content = float(cmd.content) # Tenta converter o valor da temperatura para float
  except:
    return cmd_data.BasicCmd("error", "O valor da temperatura deve ser um número")
  
  data['temperature'] = cmd.content
  return cmd_data.BasicCmd("set_temp", f'Temperature set to {cmd.content}')
```

Para registrar o comando `set_temp`, o desenvolvedor deve adicionar o seguinte código:
```python
  # Código retirado de: /sensor/air_cond.main.py
  # Código para registrar o comando set_temp
  server = Server("0.0.0.0", 3333) # Cria um novo servidor
  server.register_command("set_temp", set_temp_cmd) # Registra o comando set_temp
```

Além disso, o Sensor é capaz de lidar com múltiplas conexões simultâneas, garantindo que ele esteja sempre disponível para receber comandos. Para isso, o Sensor [cria uma nova thread](https://github.com/jnaraujo/tec502-iot-connect/sensor/libs/server.py#L40) para cada conexão TCP/IP que é estabelecida. Assim, ele é capaz de receber comandos de múltiplos Brokers simultaneamente. Os problemas relacionados a concorrência são improváveis, visto que seria necessário que dois Brokers enviassem comandos para o mesmo Sensor ao mesmo tempo, algo que não é esperado no sistema.

No Broker, o código para envio de comandos pode ser encontrado em [broker/internal/sensor_conn/sensor.go](https://github.com/jnaraujo/tec502-iot-connect/blob/main/broker/internal/sensor_conn/sensor.go), enquanto no Sensor, o código para receber comandos pode ser encontrado em [sensor/libs/server.py](https://github.com/jnaraujo/tec502-iot-connect/blob/main/sensor/libs/server.py).

#### Comandos disponíveis
Os comandos disponíveis para os Sensores são:

| Comando | Descrição | Observação |
|---------|-----------|------------|
| `turn_on` | Liga o dispositivo | *Disponível para todos os Sensores* |
| `turn_off` | Desliga o dispositivo | *Disponível para todos os Sensores* |
| `fas` | Lista todos os sensores registrados no Broker | *Disponível para todos os Sensores* |
| `fad` | Lista todos os dados dos sensores registrados no Broker | *Disponível para todos os Sensores* |
| `setup` | Cria um novo Sensor no Broker | *Disponível para todos os Sensores* |
| `delete` | Deleta um Sensor do Broker | *Disponível para todos os Sensores* |
| `set_temp` | Configura a temperatura do dispositivo | *Somente para o Sensor de ar condicionado* |
| `set_heat` | Configura a temperatura do ar condicionado para 40 graus | *Somente para o Sensor de ar condicionado* |
| `set_cool` | Configura a temperatura do ar condicionado para 16 graus | *Somente para o Sensor de ar condicionado* |
| `set_lux` | Configura a luminosidade do dispositivo | *Somente para o Sensor de lâmpada* |
| `lux_low` | Configura a luminosidade do dispositivo para baixa | *Somente para o Sensor de lâmpada* |
| `not_found` | Comando não encontrado | *Comando interno do Sensor* |
| `set_id` | Configura o ID do Sensor | *Comando interno do Sensor* |
| `get_commands` | Retorna a lista de comandos disponíveis para o Sensor | *Comando interno do Sensor* |

#### Envio de dados dos Sensores para o Broker
<div align="center">
<img src="./docs/imgs/diagrama-envio-dados.png" alt="Envio de dados do Sensor para o Broker" height="300px" width="auto" /> <br/>
<em>Figura 8. Diagrama de envio de dados do Sensor para o Broker</em> <br/>
</div>

Para o envio de dados dos Sensores para o Broker, é utilizada uma abordagem não confiável (UDP). Assim, os Sensores enviam os dados para o Broker através de pacotes UDP. A abordagem não confiável foi escolhida pois o protocolo UDP é mais leve e mais rápido que o TCP. Desse modo, como os sensores estão constantemente enviando dados para o Broker, caso algum pacote seja perdido, o Sensor irá enviar um novo na próxima iteração, não causando maiores problemas ao sistema.

Ao se iniciar o Broker, um socket UDP é aberto na porta 5310 e o Broker fica aguardando por novos pacotes. Sempre que um Sensor envia um dado, ele envia um pacote UDP para o Broker. O Broker, por sua vez, recebe o pacote e processa o dado.

Sempre que um [novo dado chega ao Broker](https://github.com/jnaraujo/tec502-iot-connect/blob/431feac1735b679ace3b3878374cc705a543573b/broker/internal/udp_server/server.go#L22), é verificado se o Sensor que enviou o dado está cadastrado. Caso não esteja, o dado é descartado. Caso esteja, o dado é [armazenado na lista de conteúdo do Sensor](https://github.com/jnaraujo/tec502-iot-connect/blob/431feac1735b679ace3b3878374cc705a543573b/broker/internal/udp_server/server.go#L46).

No Broker, o código para receber os dados dos Sensores pode ser encontrado em [broker/internal/udp_server/server.go](/broker/internal/udp_server/server.go), enquanto no Sensor, o código para enviar os dados pode ser encontrado em [sensor/libs/broker_service.py](/sensor/libs/broker_service.py).

##### Lidando com concorrência
Como diversos dados diferentes chegam ao Broker ao mesmo tempo, é necessário que o Broker seja capaz de lidar com múltiplos pacotes UDP simultaneamente. Para isso, a cada novo dado que chega, uma nova [goroutine](https://github.com/jnaraujo/tec502-iot-connect/blob/431feac1735b679ace3b3878374cc705a543573b/broker/internal/udp_server/udp.go#L40) é criada para lidar com esse dado. Assim, o Broker é capaz de lidar com múltiplos pacotes UDP simultaneamente, garantindo que ele esteja sempre disponível para receber os dados dos Sensores.

Por exemplo, o código mostra como o Broker lida com a chegada de um novo pacote UDP:
```go	
// Código retirado de: broker/internal/udp_server/udp.go
// código omitido
for {
  buffer := make([]byte, 1024) // Cria um buffer para armazenar o pacote UDP
  n, addr, err := conn.ReadFrom(buffer) // Lê o pacote UDP
  // código omitido
  go u.handler(addr.String(), string(buffer[:n])) // Cria uma nova goroutine para lidar com o pacote UDP
}
// código omitido
```

> Vale destacar que goroutines são semelhantes a threads, mas são mais leves e mais eficientes. Assim, o uso de goroutines permite que o Broker seja capaz de lidar com múltiplas conexões simultâneas sem consumir muitos recursos do sistema.

Para lidar com problemas de concorrência, [foi implementado um mecanismo de trava (mutex)](https://github.com/jnaraujo/tec502-iot-connect/blob/431feac1735b679ace3b3878374cc705a543573b/broker/internal/storage/responses/responses.go#L19) para garantir que apenas uma goroutine acesse à estrutura que armazena os dados dos Sensores por vez. Assim, problemas de concorrência relacionados ao armazenamento dos dados são improváveis.

## Confiabilidade da solução e tolerância a falhas
Em um sistema de troca de mensagens confiável, é necessário adotar de estratégias para lidar com falhas de comunicação entre os dispositivos. Para isso, algumas estratégias foram adotadas, como o uso de protocolos confiáveis, tratamento de erros e envio constante de dados em caso de uso de protocolos não confiáveis.

### Uso de protocolos confiáveis
Para garantir a confiabilidade do sistema, foram utilizados protocolos confiáveis para a troca de mensagens entre o Client e o Broker e entre o Broker e os Sensores. O uso de protocolos confiáveis garante que as mensagens sejam entregues corretamente e que o sistema seja capaz de lidar com problemas na comunicação.

Entre o Client e o Broker, foi utilizado o protocolo HTTP, que é baseado em TCP/IP e garante a entrega das mensagens. Assim, caso ocorra algum problema na comunicação, o usuário é sempre avisado. Além disso, o Broker é capaz de lidar com múltiplas conexões simultâneas, garantindo que a aplicação esteja sempre disponível.

Para a troca de mensagens do Broker para os Sensores, foi adotada uma abordagem confiável utilizando o protocolo TCP/IP. Isso significa que, em caso de problemas na comunicação, como perda de pacotes ou interrupções na conexão, o Broker é capaz de detectar essas falhas e informar ao usuário sobre a ocorrência de qualquer problema na transmissão de dados. Essa confiabilidade é importante para garantir que os comandos enviados pelo Broker sejam sempre entregues aos Sensores.

### Tratamento de erros
Para garantir a confiabilidade do sistema, foram implementados mecanismos de tratamento de erros em todas as partes do sistema. Por exemplo, na comunicação entre o Client e o Broker, caso ocorra algum problema na comunicação, o Broker é capaz de detectar o erro e informar ao usuário sobre a ocorrência de qualquer problema na transmissão de dados. Assim, o usuário é informado na própria interface sobre possíveis problemas na comunicação, além de ser possível visualizar o status dos Sensores e do Broker.

### Envio constante de dados em caso de uso de protocolos não confiáveis
O envio de dados dos Sensores para o Broker é feito através de uma abordagem não confiável (UDP). Assim, caso ocorra algum problema na comunicação, a mensagem será perdida. Porém, como os Sensores estão constantemente enviando dados para o Broker, caso algum pacote seja perdido, o Sensor irá enviar um novo na próxima iteração, não causando maiores problemas ao sistema.

Vale destacar que, caso o Broker não receba os dados de um Sensor por um determinado tempo (por exemplo, 5 segundos), ele considerará o Sensor como offline. Assim que o Sensor voltar a enviar dados, ele voltará a ser considerado online. Essa abordagem garante que o Broker seja capaz de lidar com problemas na comunicação e que o usuário seja sempre informado sobre o status dos Sensores.

## Testes
Para garantir o funcionamento correto do sistema, alguns módulos apresentam testes unitários ([Cmd](/broker/internal/cmd/cmd_test.go) e [Queue](/broker/internal/queue/queue_test.go)). Para rodar os testes unitários, basta executar o comando `go test ./internal/...` na pasta do Broker.

Além disso, todas as rotas da API REST possuem testes de integração. Os testes de integração garantem que a API está funcionando corretamente e que os dados estão sendo armazenados e recuperados corretamente. Os testes foram feitos utilizando a biblioteca padrão do Go. Para rodar os testes, basta executar o comando `go test ./test/...` na pasta do Broker.

> Nota: Para rodar os testes de integração, é necessário ter o Sensor do Ar Condicionado rodando com o IP *localhost:3399*. Para isso, basta executar o Dockerfile do Sensor do Ar Condicionado com `docker build -t sensor -f AirCond.Dockerfile . && docker run  -p 3399:3333 -e BROKER_URL=<broker url> -it sensor`.

Além disso, todas as rotas podem ser testadas utilizando o [Postman](https://www.postman.com/) ou o [Bruno](https://www.usebruno.com/) para garantir que a API está funcionando corretamente. Os arquivos do Postman podem ser encontrados na em [/iot-connect-api-postman.json](/iot-connect-api-postman.json) e no Bruno em [/iot-connect-api-bruno](/iot-connect-api-bruno/).

## Conclusão
O projeto desenvolvido cumpriu com os objetivos propostos, criando um sistema de comunicação e gerenciamento entre dispositivos IoT e suas diferentes aplicações. O sistema é capaz de criar, remover e verificar o status de dispositivos IoT, além de receber, armazenar e disponibilizar dados dos dispositivos IoT para aplicações web.

O sistema foi desenvolvido utilizando tecnologias modernas e atuais, como React, Go e Python. Além disso, o sistema foi desenvolvido utilizando Docker e Docker Compose, o que facilita a execução do sistema em diferentes ambientes.

O sistema é capaz de lidar com falhas de comunicação entre os dispositivos, garantindo que o usuário seja sempre informado caso ocorra algum problema na comunicação. Além disso, o sistema é capaz de lidar com múltiplas conexões simultâneas, garantindo que ele esteja sempre disponível para receber os dados dos Sensores.
