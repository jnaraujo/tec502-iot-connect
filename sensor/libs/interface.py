import urllib.request
from libs.server import Server
from libs.cmd_data import Cmd
import json

class Interface:
  '''
  Classe que representa a interface do sensor. A interface é responsável por receber comandos do usuário e executá-los, além de permitir algumas conexões com o broker.
  '''
  
  # Comandos de configuração
  config_cmds = [
    {
      "name": "setup",
      "description": "Criar sensor no broker",
      "usage": "<broker_addr> <sensor_id> <sensor_addr>"
    },
    {
      "name": "delete",
      "description": "Deletar sensor no broker",
      "usage": "<broker_addr> <sensor_id>"
    }
  ]
  
  def __init__(self, server: Server):
    '''
    Construtor da classe Interface.
    '''
    
    self.server = server
  
  def run(self):
    '''
    Inicia a interface.
    '''
    
    print("=+"*15 + "=")
    print("Comandos do sensor:")
    for cmd in self.server.commands: # Mostra os comandos do Sensor
      print(f'Comando: {cmd}')
    print("")
    print("Comandos de configuração:")
    for cmd in self.config_cmds: # Mostra os comandos de configuração
      print(f'Comando: {cmd["name"]} - {cmd["description"]} - Uso: {cmd["usage"]}')
    print("=+"*15 + "=")
    
    while True:
      try:
        try:
          cmd_in = input("Digite o comando: ")
        except KeyboardInterrupt:
          break
        except EOFError:
          print("\nSaindo...")
          break
        values = cmd_in.split(' ')
                  
        if cmd_in.startswith('setup'): # Para criar um sensor no broker
          broker_addr = values[1]
          sensor_id = values[2]
          sensor_addr = values[3]
          
          self.server.sensor_id = sensor_id # Atualiza o ID do sensor
          
          resp = self.register_sensor_on_broker(broker_addr=broker_addr, sensor_id=sensor_id, sensor_addr=sensor_addr)
          print(f'Resposta: {resp}')
          continue
        
        if cmd_in.startswith('delete'): # Para deletar um sensor no broker
          broker_addr = values[1]
          sensor_id = values[2]
          
          resp = self.delete_sensor_on_broker(broker_addr=broker_addr, sensor_id=sensor_id)
          print(f'Resposta: {resp}')
          continue
        
        cmd = Cmd(idFrom="INTERFACE", idTo=self.server.sensor_id, command=values[0], content=' '.join(values[1:])) 
        cmd_out = self.server.handle_command(cmd) # Executa o comando e pega a resposta
        
        print(f'Resposta: {cmd_out.content}')
      except Exception as e:
        print('Error:', e)
        
  def register_sensor_on_broker(self, broker_addr: str, sensor_id:str, sensor_addr: str):
    '''
    Registra o sensor no broker.
    '''
    
    try:
      req = urllib.request.Request(f'http://{broker_addr}/sensor', method='POST') # Cria a requisição
      req.add_header('Content-Type', 'application/json') # Adiciona o cabeçalho de JSON
      req.data = json.dumps({
        'address': sensor_addr,
        'id': sensor_id
      }).encode('utf-8') # Adiciona o corpo da requisição
      
      resp = urllib.request.urlopen(req, timeout=1) # Envia a requisição (com timeout de 1 segundo)
      return json.loads(resp.read())
    except Exception as e:
      return {'error': 'Erro durante a solicitação: ' + str(e)}
  
  def delete_sensor_on_broker(self, broker_addr: str, sensor_id:str):
    '''
    Deleta o sensor no broker.
    '''
    
    resp = urllib.request.urlopen(f'http://{broker_addr}/sensor/{sensor_id}', method='DELETE') # Envia a requisição
    return json.loads(resp.read())