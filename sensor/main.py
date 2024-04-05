import datetime
import socket
import sys
import time
from server import Server
from cmd_data import Cmd
from broker_service import BrokerService

bs = BrokerService(('localhost', 5310))

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
  
  server.start()
  
def not_found_cmd(cmd: Cmd):
  res = Cmd(cmd.id, cmd.content, "Command not found")
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

if __name__ == '__main__':
  init()