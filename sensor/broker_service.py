import cmd_data
import socket

class BrokerService:
  def __init__(self, address: str):
    self.address = address
  
  def send(self, cmd: cmd_data.Cmd):
    try:
      sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
      print("test")
      sock.connect(self.address)
      print("after connect")
      
      data = cmd_data.encode(cmd)
      sock.sendall(data)
      print("after sendall")
      
      sock.close()
      
      return True
    except Exception as e:
      print('Error sending data:', e)
      return False