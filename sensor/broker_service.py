import cmd_data
import socket

class BrokerService:
  def __init__(self, address: str):
    self.address = address
  
  def send(self, cmd: cmd_data.Cmd):
    sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    sock.connect(self.address)
    
    data = cmd_data.encode(cmd)
    sock.sendall(data)
    
    sock.close()