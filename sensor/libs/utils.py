import socket

def get_current_ip():
  IPAddr = socket.gethostbyname(socket.gethostname())
  return IPAddr