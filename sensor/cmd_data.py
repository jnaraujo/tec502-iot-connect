class Cmd:
  def __init__(self, id: str, command: str, content: str = ''):
    self.id = id
    self.command = command
    self.content = content
  
  def __getitem__(self, key):
    return getattr(self, key)

def decode(data: bytes):
  """
    Data format:
    
    ```
    Id: <id>
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
  
  if not header[0].startswith('Id: '):
    raise ValueError('Header must start with "Id: "')
  
  if not header[1].startswith('Cmd: '):
    raise ValueError('Header must start with "Cmd: "')
  
  return Cmd(
    id=header[0][4:],
    command=header[1][5:],
    content=data[1] if len(data) > 1 else ''
  )
  
def encode(cmd: Cmd):
  return f'Id: {cmd.id}\nCmd: {cmd.command}\n\n{cmd.content}'.encode('utf-8')