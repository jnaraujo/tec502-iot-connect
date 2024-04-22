from libs import cmd_data
import socket

class BrokerService:
  def __init__(self, address: tuple):
    self.address = address
    
  def set_address(self, address: tuple):
    self.address = address
  
  def send(self, content: str):
    sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    sock.sendto(content, self.address)