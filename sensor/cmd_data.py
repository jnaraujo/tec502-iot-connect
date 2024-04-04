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
  
  return {
    'id': header[0].split(': ')[1],
    'command': header[1].split(': ')[1],
    'content': '\n'.join(data[1:])
  }
  
def encode(id: str, command: str, content: str = ''):
  return f'Id: {id}\nCmd: {command}\n\n{content}'.encode('utf-8')