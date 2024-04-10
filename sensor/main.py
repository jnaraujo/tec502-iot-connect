import datetime
import socket
import sys
import time
import os
from threading import Thread
from server import Server
import cmd_data
from broker_service import BrokerService

broker_url = os.environ.get('BROKER_URL', 'localhost:5310')
broker_addr = (broker_url.split(':')[0], int(broker_url.split(':')[1]))

bs = BrokerService(broker_addr)

STATUS = False # O Status indica se o sensor está ligado ou desligado
data = {
  "temperature": 30.5,
}

def init():
  IP_ADDR = "0.0.0.0"
  IP_PORT = 3333
  
  if len(sys.argv) > 1:
    IP_PORT = int(sys.argv[1])
  
  print("="*30)
  print("Sensor Server")
  print(f'Sensor IP: {get_current_ip()}:{IP_PORT}')
  print("="*30)
  
  server = Server(IP_ADDR, IP_PORT)
  
  server.register_not_found(not_found_cmd)
  
  server.register_command("set_temp", set_temp_cmd)
  server.register_command("turn_on", turn_on_cmd)
  server.register_command("turn_off", turn_off_cmd)
  
  Thread(target=server.start).start()
  Thread(target=send_broker_data, args=(server,)).start()
  
def send_broker_data(server: Server):
  while True:
    time.sleep(5) # Envia dados a cada 5 segundos
    if not STATUS: # Se o sensor estiver desligado, não envia dados
      print("Sensor is off")
      continue
    try:
      cmd = cmd_data.Cmd(
        idFrom=server.get_sensor_id(), idTo="BROKER", command='temperature',content=str(data["temperature"])
        )
      bs.send(cmd_data.encode(cmd))
    except Exception as e:
      print("Error sending data to broker:", e)
  
def not_found_cmd(cmd: cmd_data.Cmd):
  return cmd_data.BasicCmd("not_found", f'Command {cmd.command} not found')

def set_temp_cmd(cmd: cmd_data.Cmd):
  data['temperature'] = cmd.content
  return cmd_data.BasicCmd("set_temp", f'Temperature set to {cmd.content}')

def turn_on_cmd(cmd: cmd_data.Cmd):
  global STATUS
  STATUS = True
  return cmd_data.BasicCmd("turn_on", "Sensor turned on")

def turn_off_cmd(cmd: cmd_data.Cmd):
  global STATUS
  STATUS = False
  return cmd_data.BasicCmd("turn_off", "Sensor turned off")
  
def get_current_ip():
  IPAddr = socket.gethostbyname(socket.gethostname())
  return IPAddr

init()