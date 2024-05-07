import datetime
import sys
import time
import os
import random
from threading import Thread
from libs import cmd_data, utils
from libs.broker_service import BrokerService
from libs.server import Server
from  libs.interface import Interface

broker_url = os.environ.get('BROKER_URL', 'localhost:5310')
broker_addr = (broker_url.split(':')[0], int(broker_url.split(':')[1]))
bs = BrokerService(broker_addr) # Broker Service

STATUS = False # O Status indica se o sensor está ligado ou desligado
data = {
  "temperature": 30.5,
}

def init():  
  print("="*30)
  print("Ar Condicionado - Sensor Server")
  print(f'Sensor IP: {utils.get_current_ip()}:3333')
  print("="*30)
  
  server = Server("0.0.0.0", 3333) # Cria o servidor
  interface = Interface(server) # Cria a interface
  
  server.register_not_found(not_found_cmd) # Registra o comando de não encontrado
  
  # Registra os comandos
  server.register_command("set_temp", set_temp_cmd)
  server.register_command("turn_on", turn_on_cmd)
  server.register_command("turn_off", turn_off_cmd)
  server.register_command("set_heat", set_heat_cmd)
  server.register_command("set_cool", set_cool_cmd)
  
  Thread(target=server.start).start() # Inicia o servidor
  Thread(target=send_broker_data, args=(server,)).start() # Inicia o envio de dados para o broker
  interface.run() # Inicia a interface
  
def send_broker_data(server: Server):
  '''
  Função que envia dados para o broker a cada 2 segundos.
  '''
  
  while True:
    time.sleep(0.5) # Espera 500 ms
    if not STATUS: # Se o sensor estiver desligado, não envia dados
      continue
    try:
      # Muda a temperatura de forma aleatória
      data["temperature"] = round(data["temperature"] + random.uniform(-0.5, 0.5), 2)
      cmd = cmd_data.Cmd(
        idFrom=server.get_sensor_id(), idTo="BROKER", command='temperature',content=str(data["temperature"])
      )
      bs.send(cmd_data.encode(cmd)) # Envia os dados para o broker
    except Exception as e:
      print("Error sending data to broker:", e)
  
def not_found_cmd(cmd: cmd_data.Cmd):
  '''
  Função que lida com comandos não encontrados.
  '''
  
  return cmd_data.BasicCmd("not_found", f'Comando {cmd.command} não encontrado')

def set_temp_cmd(cmd: cmd_data.Cmd):
  '''
  Comando que seta a temperatura do sensor.
  '''
  
  if not STATUS: # Se o sensor estiver desligado, retorna um erro
    return cmd_data.BasicCmd("error", "O sensor está desligado")
  
  try:
    cmd.content = float(cmd.content) # Tenta converter o valor da temperatura para float
  except:
    return cmd_data.BasicCmd("error", "O valor da temperatura deve ser um número")
  
  data['temperature'] = cmd.content
  return cmd_data.BasicCmd("set_temp", f'Temperature set to {cmd.content}')

def set_heat_cmd(cmd: cmd_data.Cmd):
  '''
  Comando que seta o modo de aquecimento do sensor.
  '''
  
  if not STATUS: # Se o sensor estiver desligado, retorna um erro
    return cmd_data.BasicCmd("error", "O sensor está desligado")
  
  data["temperature"] = 40
  return cmd_data.BasicCmd("set_heat", "Modo de aquecimento ativado")

def set_cool_cmd(cmd: cmd_data.Cmd):
  '''
  Comando que seta o modo de resfriamento do sensor.
  '''
  
  if not STATUS: # Se o sensor estiver desligado, retorna um erro
    return cmd_data.BasicCmd("error", "O sensor está desligado")
  
  data["temperature"] = 16
  return cmd_data.BasicCmd("set_cool", "Modo de resfriamento ativado")

def turn_on_cmd(cmd: cmd_data.Cmd):
  '''
  Comando que liga o sensor.
  '''
  
  global STATUS
  STATUS = True
  return cmd_data.BasicCmd("turn_on", "Ar Condicionado foi ligado")

def turn_off_cmd(cmd: cmd_data.Cmd):
  '''
  Comando que desliga o sensor.
  '''
  
  global STATUS
  STATUS = False
  return cmd_data.BasicCmd("turn_off", "Ar Condicionado foi desligado")

init()