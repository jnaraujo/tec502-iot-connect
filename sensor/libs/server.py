import socket
from threading import Thread
import random
from libs import cmd_data

class Server:
  HANDSHAKE_RECEIVED = b'hello, sensor!'
  HANDSHAKE_SENT = b'hello, server!'
  
  sensor_id=f'sensor-{random.randint(0, 1000)}'
  
  def __init__(self, host: str, port: int):
    self.host = host
    self.port = port
    self.commands = {}

  def start(self):
    self.sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    self.sock.bind((self.host, self.port))
    self.sock.listen(10)
    
    self.handle_connections()
    
  def get_sensor_id(self):
    return self.sensor_id
  
  def stop(self):
    self.sock.close()
    
  def handle_connections(self):
    while True:
      conn, client_address = self.sock.accept()
      print('Connection from:', client_address)
      
      try:
        if not self.validate_connection(conn):
          conn.close()
          continue
        
        Thread(target=self.handle_connection, args=(conn,)).start()
      except Exception as e:
        print('Error handling connection:', e)
        
  def handle_connection(self, conn: socket.socket):    
    data = conn.recv(1024)
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
    
    if command == "set_id":
      self.sensor_id = data['content']
      cmd = cmd_data.Cmd(
        idFrom=self.sensor_id,
        idTo="BROKER",
        command='id',
        content=self.sensor_id
      )
      conn.sendall(cmd_data.encode(cmd))
      return
    
    if self.sensor_id != data['idTo']:
      self.sensor_id = data['idTo'] # Atualiza o ID do sensor caso seja diferente
    
    if command == "get_commands":
      conn.sendall(self.get_commands())
      return
    
    resCmd: cmd_data.Cmd = None
    
    if command not in self.commands:
      if 'not_found' in self.commands:
        resCmd = self.commands['not_found'](data)
    else:
      resCmd = self.commands[command](data)
    
    # Se o comando não tiver um ID, seta o ID do sensor
    if resCmd.idFrom is None:
      resCmd.idFrom = self.sensor_id
    if resCmd.idTo is None: # Se o comando não tiver um destinatário, seta o destinatário como o broker
      resCmd.idTo = 'BROKER'
    
    conn.sendall(cmd_data.encode(resCmd))
      
  def validate_connection(self, conn: socket.socket) -> bool:
    data = conn.recv(len(self.HANDSHAKE_RECEIVED))
    
    if data == self.HANDSHAKE_RECEIVED:
      conn.sendall(self.HANDSHAKE_SENT)
      return True
    else:
      conn.sendall(b'invalid handshake')
      return False
  
  def get_commands(self):
    commands = list(self.commands.keys())    
    cmd = cmd_data.encode(
      cmd_data.Cmd(
        idFrom=self.sensor_id, idTo="BROKER",
        command="commands", content=", ".join(commands)
      )
    )
    return cmd
  
  def register_not_found(self, callback: callable):
    self.commands['not_found'] = callback
  
  def register_command(self, command: str, callback: callable):
    self.commands[command] = callback