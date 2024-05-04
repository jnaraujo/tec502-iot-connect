class Cmd:
  '''
  Classe responsável por definir o protocolo de comunicação entre os dispositivos e o broker.
  
  O protocolo é definido por um cabeçalho e um corpo. O cabeçalho é composto por três campos: IdFrom, IdTo e Cmd. O corpo é o conteúdo da mensagem.
  
  Mais detalhes sobre o protocolo podem ser encontrados em: https://github.com/jnaraujo/tec502-iot-connect/blob/80ff70daff61df098ce99ed18c17fee974021a06/README.md#L191
  '''
  
  def __init__(self, idFrom: str, idTo: str, command: str, content: str = ''):
    '''
    Construtor da classe.
    '''
    
    self.idFrom = idFrom
    self.idTo = idTo
    self.command = command
    self.content = content
  
  def __getitem__(self, key):
    '''
    Retorna o valor de um atributo. Esse método é chamado quando se usa a notação de colchetes.
    
    Exemplo:
      
      ```python 
      cmd = Cmd('1', '2', 'cmd', 'content')
      print(cmd['idFrom']) # 1
      ```
    '''
    return getattr(self, key)
  
def BasicCmd(command: str, content: str = ''):
  '''
  Função que cria um comando básico.
  '''
  
  return Cmd(None, None, command, content)

def decode(data: bytes):
  """
    Função que decodifica os dados recebidos para um objeto Cmd.
  """
  
  data = data.decode('utf-8') # Decodifica os dados
  data = data.split('\n\n') # Separa o cabeçalho do corpo
  
  if len(data) < 1: # Se não houver dados, retorna um erro
    raise ValueError('Invalid data')
  
  header = data[0].split('\n') # Separa o cabeçalho em linhas
  
  if len(header) < 2: # Se não houver linhas no cabeçalho, retorna um erro
    raise ValueError('Invalid header')
  
  if not header[0].startswith('IdFrom: '): # Se a primeira linha do cabeçalho não começar com "IdFrom: ", retorna um erro
    raise ValueError('Header must start with "IdFrom: "')
  
  if not header[1].startswith('IdTo: '): # Se a segunda linha do cabeçalho não começar com "IdTo: ", retorna um erro
    raise ValueError('Header must start with "IdTo: "')
  
  if not header[2].startswith('Cmd: '): # Se a terceira linha do cabeçalho não começar com "Cmd: ", retorna um erro
    raise ValueError('Header must start with "Cmd: "')
  
  idFrom = header[0].split(': ')[1]
  idTo = header[1].split(': ')[1]
  command = header[2].split(': ')[1]
  
  return Cmd(
    idFrom=idFrom, idTo=idTo, command=command,
    content=data[1] if len(data) > 1 else ''
  )
  
def encode(cmd: Cmd):
  '''
  Função que codifica um objeto Cmd para bytes.
  '''
  
  return f'IdFrom: {cmd.idFrom}\nIdTo: {cmd.idTo}\nCmd: {cmd.command}\n\n{cmd.content}'.encode('utf-8')