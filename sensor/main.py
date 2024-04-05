import datetime
import socket
import sys
import time
import os
from threading import Thread
from server import Server
from cmd_data import Cmd
from broker_service import BrokerService

broker_url = os.environ.get('BROKER_URL', 'localhost:5173')
broker_addr = (broker_url.split(':')[0], int(broker_url.split(':')[1]))

bs = BrokerService(broker_addr)

data = {}

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
  
  server.register_command("get_time", get_time_cmd)
  server.register_command("get_ip", get_ip_cmd)
  server.register_command("test", test_cmd)
  server.register_command("delay", delay)
  server.register_command("set_temp", set_temp_cmd)
  server.register_command("get_temp", get_temp_cmd)
  
  Thread(target=server.start).start()
  
def not_found_cmd(cmd: Cmd):
  res = Cmd(cmd.id, cmd.content, "Command not found")
  bs.send(res)
  
def set_temp_cmd(cmd: Cmd):
  data['temp'] = cmd.content
  res = Cmd(cmd.id, cmd.content, "Temperature set")
  bs.send(res)
  
def get_temp_cmd(cmd: Cmd):
  temp = data.get('temp', 'N/A')
  res = Cmd(cmd.id, cmd.content, temp)
  bs.send(res)
  
def delay(cmd: Cmd):
  time.sleep(5)
  res = Cmd(cmd.id, cmd.content, "Delayed response")
  bs.send(res)
  
def test_cmd(cmd: Cmd):
  res = Cmd(cmd.id, "test", "Hello from sensor!")
  bs.send(res)
  
def get_ip_cmd(req: Cmd):
  res = Cmd(req.id, "get_ip", get_current_ip())
  bs.send(res)
  
def get_time_cmd(req: Cmd):
  res = Cmd(req.id, "get_time", datetime.datetime.now().strftime('%Y-%m-%d %H:%M:%S'))
  bs.send(res)
  
def get_current_ip():
  IPAddr = socket.gethostbyname(socket.gethostname())
  return IPAddr

init()