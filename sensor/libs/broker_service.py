from libs import cmd_data
import socket

class BrokerService:
  '''
  Classe que representa um serviço de broker. O serviço de broker é responsável por enviar mensagens para um broker via UDP.
  '''
  
  def __init__(self, address: tuple):
    '''
    Construtor da classe BrokerService.
    '''
    
    self.address = address
    
  def set_address(self, address: tuple):
    '''
    Define o endereço do broker.
    '''
    
    self.address = address
  
  def send(self, content: str):
    '''
    Envia uma mensagem para o broker.
    '''
    
    sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    sock.sendto(content, self.address)