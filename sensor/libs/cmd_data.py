class Cmd:
  def __init__(self, idFrom: str, idTo: str, command: str, content: str = ''):
    self.idFrom = idFrom
    self.idTo = idTo
    self.command = command
    self.content = content
  
  def __getitem__(self, key):
    return getattr(self, key)
  
def BasicCmd(command: str, content: str = ''):
  return Cmd(None, None, command, content)

def decode(data: bytes):
  """
    Data format:
    
    ```
    IdFrom: <id>
    IdTo: <id>
    Cmd: <comando>

    <conteúdo (opcional)>
    ```
  """
  
  data = data.decode('utf-8')
  data = data.split('\n\n')
  
  if len(data) < 1:
    raise ValueError('Invalid data')
  
  header = data[0].split('\n')
  
  if len(header) < 2:
    raise ValueError('Invalid header')
  
  if not header[0].startswith('IdFrom: '):
    raise ValueError('Header must start with "IdFrom: "')
  
  if not header[1].startswith('IdTo: '):
    raise ValueError('Header must start with "IdTo: "')
  
  if not header[2].startswith('Cmd: '):
    raise ValueError('Header must start with "Cmd: "')
  
  idFrom = header[0].split(': ')[1]
  idTo = header[1].split(': ')[1]
  command = header[2].split(': ')[1]
  
  print(idFrom, idTo, command)
  
  return Cmd(
    idFrom=idFrom, idTo=idTo, command=command,
    content=data[1] if len(data) > 1 else ''
  )
  
def encode(cmd: Cmd):
  return f'IdFrom: {cmd.idFrom}\nIdTo: {cmd.idTo}\nCmd: {cmd.command}\n\n{cmd.content}'.encode('utf-8')