import socket

server_address = ('127.0.0.1', 3333)

def init():
  sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
  sock.bind(server_address)
  sock.listen(1)

  while True:
    print('Waiting for a connection...')
    connection, client_address = sock.accept()
    try:
      handleConnection(connection, client_address)
    except Exception as e:
      print('Error:', e)
    finally:
      connection.close()

def handleConnection(conn: socket.socket, addr: tuple):
  print('Connection from', addr)

  while True:
    data = conn.recv(1024)
    print('Received', data)
    if data:
      print('Sending data back to the client')
      conn.sendall(data)
    else:
      print('No more data from', addr)
      break

if __name__ == '__main__':
  init()