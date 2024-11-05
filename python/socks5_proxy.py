import requests
import socks
import socket

# Proxy configuration
proxy_username = 'change_with_your_proxy_username'  # TODO Replace with your proxy username
proxy_password = 'change_with_your_proxy_password'  # TODO Replace with your proxy password
proxy_host = '2.56.100.66'  # TODO Replace with your proxy IP
proxy_port = 44444  # TODO Replace with your proxy port

# Format the proxy URLs for HTTP and HTTPS connections using SOCKS5
proxies = {
    "http": "socks5://{0}:{1}@{2}:{3}".format(proxy_username, proxy_password, proxy_host, proxy_port),
    "https": "socks5h://{0}:{1}@{2}:{3}".format(proxy_username, proxy_password, proxy_host, proxy_port)
}

# Make a request using the proxy settings via Requests library
try:
    response = requests.get("https://tools.proxysocks5.com/api/ip/info", proxies=proxies)
    # Decode and print the IP address returned by the API
    print("IP address using requests with proxy:", response.content.decode('utf-8'))
except requests.RequestException as e:
    print("Error with proxy request using Requests library:", e)

# -- Using `socks.set_default_proxy` to route all socket requests through the proxy -- #

# Set the default SOCKS5 proxy with authentication for all socket operations
socks.set_default_proxy(socks.SOCKS5, proxy_host, proxy_port, username=proxy_username, password=proxy_password)

# Override the default socket with `socksocket` to ensure all socket-based requests use the proxy
socket.socket = socks.socksocket

# Make a request to verify IP using the default proxy setting through socket-based requests
try:
    response = requests.get("https://tools.proxysocks5.com/api/ip/info")
    # Decode and print the IP address returned by the API
    print("IP address using socket with proxy:", response.content.decode('utf-8'))
except requests.RequestException as e:
    print("Error with proxy request using socket-based approach:", e)
