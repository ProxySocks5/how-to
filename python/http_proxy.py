import requests

# Proxy credentials and host configuration
proxy_username = 'change_with_your_proxy_username'  # TODO Replace with your proxy username
proxy_password = 'change_with_your_proxy_password'  # TODO Replace with your actual proxy password
proxy_host_port = '45.43.134.162:43222'  # TODO Replace with your proxy's IP:Port

# Format HTTP and HTTPS proxy URLs, using HTTP Basic Authentication
proxies = {
    "http": "http://{0}:{1}@{2}".format(proxy_username, proxy_password, proxy_host_port),
    "https": "http://{0}:{1}@{2}".format(proxy_username, proxy_password, proxy_host_port)
}

# Attempt to make a request using the configured HTTP/HTTPS proxy settings
try:
    response = requests.get("https://tools.proxysocks5.com/api/ip/info", proxies=proxies)
    # Decode the response content to get the IP address (should reflect proxy IP)
    print("IP address using HTTP proxy:", response.content.decode('utf-8'))
except requests.RequestException as e:
    # Handle any exceptions that occur during the request
    print("Error occurred during the HTTP proxy request:", e)
