from server import Server
import datetime
import socket
import sys



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
  
  server.register_command("get_time", get_time_cmd)
  server.register_command("get_ip", get_ip_cmd)
  server.register_command("test", test_cmd)
  
  server.start()
  
def test_cmd(req: dict):
  return {
    'command': 'test',
    'content': req
  }
  
def get_ip_cmd(req: dict):
  return {
    'command': 'get_ip',
    'content': get_current_ip()
  }
  
def get_time_cmd(req: dict):
  return {
    'command': 'get_time',
    'content': datetime.datetime.now().strftime('%Y-%m-%d %H:%M:%S')
  }
  
def get_current_ip():
  IPAddr = socket.gethostbyname(socket.gethostname())
  return IPAddr

if __name__ == '__main__':
  init()