import socket
from threading import Thread
import cmd_data

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
    
    self.handle_connections()
  
  def stop(self):
    self.sock.close()
    
  def handle_connections(self):
    while True:
      print('Waiting for connection...')
      conn, client_address = self.sock.accept()
      
      print('Connection from:', client_address)
      
      try:
        if not self.validate_connection(conn):
          conn.close()
          continue
        
        Thread(target=self.handle_connection, args=(conn,), daemon=True).start()
      except Exception as e:
        print('Error handling connection:', e)
        
  def handle_connection(self, conn: socket.socket):    
    data = conn.recv(1024)
    print('Received:', data)
    if not data:
      conn.close()
      return
    
    try:
      cmd = cmd_data.decode(data)
      self.handle_command(cmd, conn)
    except Exception as e:
      conn.sendall(b'error decoding data')
      print('Error decoding data:', e)
      
    conn.close()
        
  def handle_command(self, data: cmd_data.Cmd, conn: socket.socket):
    command = data['command']
    
    if command not in self.commands:
      if 'not_found' in self.commands:
        self.commands['not_found'](data)
      conn.sendall(b'command not found')
      return
    
    conn.sendall(b'command received')
    self.commands[command](data)
      
  def validate_connection(self, conn: socket.socket) -> bool:
    data = conn.recv(len(self.HANDSHAKE_RECEIVED))
    
    if data == self.HANDSHAKE_RECEIVED:
      conn.sendall(self.HANDSHAKE_SENT)
      return True
    else:
      conn.sendall(b'invalid handshake')
      return False
  
  def register_not_found(self, callback: callable):
    self.commands['not_found'] = callback
  
  def register_command(self, command: str, callback: callable):
    self.commands[command] = callback