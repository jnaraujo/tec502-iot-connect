import socket

class Server:
  HANDSHAKE_RECEIVED = b'hello, sensor!'
  HANDSHAKE_SENT = b'hello, server!'

  def __init__(self, host: str, port: int):
    self.host = host
    self.port = port
    self.commands = {}

  def start(self):
    self.sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    self.sock.bind((self.host, self.port))
    self.sock.listen(1)
    
    try:
      self.handle_connections()
    except KeyboardInterrupt:
      self.stop()
      print('Server stopped')
  
  def stop(self):
    self.sock.close()
    
  def handle_connections(self):
    while True:
      print('Waiting for a connections...')
      conn, client_address = self.sock.accept()
      
      print('Connection from:', client_address)
      
      if not self.validate_connection(conn):
        conn.close()
        continue
      
      while True:
        data = conn.recv(1024)
        
        if not data:
          break
        
        try:
          data = self.decode_data(data)
          self.handle_command(data, conn)
        except Exception as e:
          conn.sendall(
            bytes('Cmd: error\n\n' \
            'Error: Invalid data', encoding='utf8')
          )
          print('Error:', e)
          raise e
        
  def handle_command(self, data: dict, conn: socket.socket):
    command = data['command']
    content = data['content']
    
    if command not in self.commands:
      conn.sendall(
        bytes('Cmd: error\n\n' \
        'Error: Command not found', encoding='utf8')
      )
      return
    
    res = self.commands[command](data)
    
    conn.sendall(
      bytes(f'Cmd: {res["command"]}\n\n' \
      f'{res["content"]}', encoding='utf8')
    )
      
  def decode_data(self, data: bytes):
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
      
  def validate_connection(self, conn: socket.socket) -> bool:
    data = conn.recv(len(self.HANDSHAKE_RECEIVED))
    
    if data == self.HANDSHAKE_RECEIVED:
      conn.sendall(self.HANDSHAKE_SENT)
      return True
    else:
      conn.sendall(b'invalid handshake')
      return False
    
  
  def register_command(self, command: str, callback: callable):
    self.commands[command] = callback