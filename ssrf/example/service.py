## Usage : python2 service.py
import socket
import BaseHTTPServer
from SimpleHTTPServer import SimpleHTTPRequestHandler

# Announce the IP address and port we will serve on
port = 8000
print("Serving on %s:%s") % (socket.gethostbyname(socket.getfqdn()), port)

# Start a server to accept traffic
addr = ("127.0.0.1", port)
server = BaseHTTPServer.HTTPServer(addr, SimpleHTTPRequestHandler)
server.serve_forever()
