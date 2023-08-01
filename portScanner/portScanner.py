import socket
import sys


if len(sys.argv) == 1:
    print("please specify host to scan")
    sys.exit(1)
host = sys.argv[1]

print(f"scanning {host}")
try:
    for port in range(1, 65535):
        sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        sock.settimeout(1)
        connection = sock.connect_ex(("1.0.0.1", port))
        if connection == 0 :
            print(f"{port} is open")
        sock.close()
except Exception as err:
    print(err)
