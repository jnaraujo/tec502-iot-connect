import socket

def get_current_ip():
  '''
  Retorna o IP atual da máquina (no caso de estar rodando no Docker, retorna p IP do container)
  '''
  
  IPAddr = socket.gethostbyname(socket.gethostname())
  return IPAddr