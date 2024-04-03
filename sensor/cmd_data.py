def decode(data: bytes):
  """
    Data format:
    
    ```
    Cmd: <comando>

    <conteúdo (opcional)>
    ```
  """
  
  data = data.decode('utf-8')
  data = data.split('\n\n')
  
  if len(data) < 1:
    raise ValueError('Invalid data')

  if not data[0].startswith('Cmd: '):
    raise ValueError('Data must start with "Cmd: "')
  
  return {
    'command': data[0].split(': ')[1],
    'content': '\n'.join(data[1:])
  }
  
def encode(command: str, content: str = ''):
  return f'Cmd: {command}\n\n{content}'.encode('utf-8')