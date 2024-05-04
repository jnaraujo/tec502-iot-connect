import socket
from threading import Thread
import random
from libs import cmd_data

class Server:
  '''
  Classe que representa um servidor de sensores. O servidor é responsável por receber comandos de sensores e mapear esses comandos para funções específicas.
  '''
  
  HANDSHAKE_RECEIVED = b'hello, sensor!'
  HANDSHAKE_SENT = b'hello, server!'
  
  sensor_id=f'sensor-{random.randint(0, 1000)}'
  
  def __init__(self, host: str, port: int):
    '''
    Construtor da classe Server.
    '''
    
    self.host = host # IP do servidor
    self.port = port # Porta do servidor
    self.commands = {} # Dicionário de comandos

  def start(self):
    '''
    Inicia o servidor e começa a escutar por conexões.
    '''
    
    self.sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    self.sock.bind((self.host, self.port))
    self.sock.listen(10)
    
    self.handle_connections()
    
  def get_sensor_id(self):
    '''
    Retorna o ID do sensor.
    '''
    
    return self.sensor_id
  
  def stop(self):
    '''
    Para o servidor.
    '''
    
    self.sock.close()
    
  def handle_connections(self):
    '''
    Lida com as conexões recebidas.
    '''
    
    while True:
      conn, client_address = self.sock.accept()
      print('Connection from:', client_address)
      
      try:
        if not self.validate_connection(conn): # Faz o handshake e valida a conexão
          conn.close() # Fecha a conexão caso o handshake falhe
          continue
        
        Thread(target=self.handle_connection, args=(conn,)).start() # Inicia uma thread para lidar com a conexão
      except Exception as e:
        print('Error handling connection:', e)
        
  def handle_connection(self, conn: socket.socket):
    '''
    Lida com uma conexão específica.
    '''
    
    data = conn.recv(1024) # Recebe os dados da conexão (máximo de 1024 bytes)
    if not data: # Se não houver dados, fecha a conexão
      conn.close()
      return
    
    try:
      cmd = cmd_data.decode(data) # Decodifica os dados recebidos
      out_cmd = self.handle_command(cmd) # Lida com o comando
      conn.sendall(cmd_data.encode(out_cmd)) # Retorna a resposta
    except Exception as e:
      conn.sendall(b'error decoding data')
      print('Error decoding data:', e)
    finally:
      conn.close()
        
  def handle_command(self, data: cmd_data.Cmd) -> cmd_data.Cmd:
    '''
    Lida com um comando específico.
    '''
    
    command = data['command'] # Pega o comando
    
    if command == "set_id": # Se o comando for set_id, seta o ID do sensor
      self.sensor_id = data['content']
      cmd = cmd_data.Cmd(
        idFrom=self.sensor_id,
        idTo="BROKER",
        command='id',
        content=self.sensor_id
      )
      return cmd
    
    if self.sensor_id != data['idTo']: # Se o destinatário do comando for diferente do ID do sensor, atualiza o ID do sensor
      self.sensor_id = data['idTo']
    
    if command == "get_commands": # Se o comando for get_commands, retorna os comandos disponíveis
      return self.get_commands()
    
    resCmd: cmd_data.Cmd = None # Inicializa a resposta do comando
    
    if command not in self.commands: # Se o comando não existir, retorna um comando de erro
      if 'not_found' in self.commands:
        resCmd = self.commands['not_found'](data)
    else:
      resCmd = self.commands[command](data)
    
    # Se o comando não tiver um ID, seta o ID do sensor
    if resCmd.idFrom is None:
      resCmd.idFrom = self.sensor_id
    if resCmd.idTo is None: # Se o comando não tiver um destinatário, seta o destinatário como o broker
      resCmd.idTo = 'BROKER'
    
    return resCmd
      
  def validate_connection(self, conn: socket.socket) -> bool:
    '''
    Valida a conexão.
    '''
    
    data = conn.recv(len(self.HANDSHAKE_RECEIVED)) # Espera receber o primeiro dado da conexão
    
    if data == self.HANDSHAKE_RECEIVED: # Se o dado recebido for o handshake, envia o handshake de volta
      conn.sendall(self.HANDSHAKE_SENT)
      return True
    else: # Se o dado não for o handshake, envia uma mensagem de erro
      conn.sendall(b'invalid handshake')
      return False
  
  def get_commands(self):
    '''
    Retorna os comandos disponíveis.
    '''
    
    commands = list(self.commands.keys()) # Pega os comandos disponíveis
    cmd = cmd_data.Cmd(
        idFrom=self.sensor_id, idTo="BROKER",
        command="commands", content=", ".join(commands)
      )
    return cmd
  
  def register_not_found(self, callback: callable):
    '''
    Registra um callback para comandos não encontrados.
    '''
    
    self.commands['not_found'] = callback
  
  def register_command(self, command: str, callback: callable):
    '''
    Registra um callback para um comando específico.
    '''
    
    self.commands[command] = callback