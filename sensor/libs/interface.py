import requests
from libs.server import Server
from libs.cmd_data import Cmd

class Interface:
  config_cmds = [
    {
      "name": "setup",
      "description": "Configurar sensor no broker",
      "usage": "<broker_addr> <sensor_id> <sensor_addr>"
    },
    {
      "name": "delete",
      "description": "Deletar sensor no broker",
      "usage": "<broker_addr> <sensor_id>"
    }
  ]
  
  def __init__(self, server: Server):
    self.server = server
  
  def run(self):
    print("=+"*15 + "=")
    print("Comandos do sensor:")
    for cmd in self.server.commands:
      print(f'Comando: {cmd}')
    
    print("")
    print("Comandos de configuração:")
    for cmd in self.config_cmds:
      print(f'Comando: {cmd["name"]} - {cmd["description"]} - Uso: {cmd["usage"]}')
    print("=+"*15 + "=")
    while True:
      try:
        cmd_in = input("Digite o comando: ")
        values = cmd_in.split(' ')
                  
        if cmd_in.startswith('setup'):
          broker_addr = values[1]
          sensor_id = values[2]
          sensor_addr = values[3]
          
          self.server.sensor_id = sensor_id
          
          resp = self.register_sensor_on_broker(broker_addr=broker_addr, sensor_id=sensor_id, sensor_addr=sensor_addr)
          print(f'Resposta: {resp}')
          continue
        
        if cmd_in.startswith('delete'):
          broker_addr = values[1]
          sensor_id = values[2]
          
          resp = self.delete_sensor_on_broker(broker_addr=broker_addr, sensor_id=sensor_id)
          print(f'Resposta: {resp}')
          continue
        
        cmd = Cmd(idFrom="INTERFACE", idTo=self.server.sensor_id, command=values[0], content=' '.join(values[1:]))
        cmd_out = self.server.handle_command(cmd)
        
        print(f'Resposta: {cmd_out.content}')
      except Exception as e:
        print('Error:', e)
        
  def register_sensor_on_broker(self, broker_addr: str, sensor_id:str, sensor_addr: str):
    resp = requests.post(f'http://{broker_addr}/sensor', json={
      'address': sensor_addr,
      'id': sensor_id
    })
    
    return resp.json()
  
  def delete_sensor_on_broker(self, broker_addr: str, sensor_id:str):
    resp = requests.delete(f'http://{broker_addr}/sensor/{sensor_id}')
    return resp.json()