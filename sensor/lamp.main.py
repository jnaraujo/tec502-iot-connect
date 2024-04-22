import datetime
import sys
import time
import os
from threading import Thread
from libs import cmd_data, utils
from libs.broker_service import BrokerService
from libs.server import Server

broker_url = os.environ.get('BROKER_URL', 'localhost:5310')
broker_addr = (broker_url.split(':')[0], int(broker_url.split(':')[1]))
bs = BrokerService(broker_addr) # Broker Service

STATUS = False # O Status indica se o sensor está ligado ou desligado
data = {
  'lux': 0,
}

def init():  
  print("="*30)
  print("Lampada - Sensor Server")
  print(f'Sensor IP: {utils.get_current_ip()}:3333')
  print("="*30)
  
  server = Server("0.0.0.0", 3333)
  
  server.register_not_found(not_found_cmd) # Registra o comando de não encontrado

  server.register_command("turn_on", turn_on_cmd)
  server.register_command("turn_off", turn_off_cmd)
  server.register_command("set_lux", set_lux_cmd)
  
  Thread(target=server.start).start()
  Thread(target=send_broker_data, args=(server,)).start()
  
def send_broker_data(server: Server):
  while True:
    time.sleep(0.5) # Envia dados a cada 0.5 segundos
    if not STATUS: # Se o sensor estiver desligado, não envia dados
      continue
    try:
      cmd = cmd_data.Cmd(
        idFrom=server.get_sensor_id(), idTo="BROKER", command='lux',content=str(data["lux"])
      )
      bs.send(cmd_data.encode(cmd))
    except Exception as e:
      print("Error sending data to broker:", e)
  
def not_found_cmd(cmd: cmd_data.Cmd):
  return cmd_data.BasicCmd("not_found", f'Comando {cmd.command} não encontrado')

def set_lux_cmd(cmd: cmd_data.Cmd):
  if not STATUS:
    return cmd_data.BasicCmd("error", "Sensor is off")
  
  data['lux'] = cmd.content
  return cmd_data.BasicCmd("set_lux", f'A luminosidade foi definida para {cmd.content}')

def turn_on_cmd(cmd: cmd_data.Cmd):
  global STATUS
  STATUS = True
  return cmd_data.BasicCmd("turn_on", "Lampada foi ligada")

def turn_off_cmd(cmd: cmd_data.Cmd):
  global STATUS
  STATUS = False
  return cmd_data.BasicCmd("turn_off", "Lampada foi desligada")

init()