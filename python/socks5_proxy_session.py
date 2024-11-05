import requests
from requests import Session

# Proxy credentials and host configuration for SOCKS5
proxy_username = 'change_with_your_proxy_username'  # TODO Replace with your proxy username
proxy_password = 'change_with_your_proxy_password'  # TODO Replace with your actual proxy password
proxy_host = '2.56.100.66'  # TODO Replace with your SOCKS5 proxy IP
proxy_port = 44444  # TODO Replace with your SOCKS5 proxy port

# Create a session to maintain settings across multiple requests
with Session() as session:
    # Set up the SOCKS5 proxy with authentication
    session.proxies = {
        "http": f"socks5://{proxy_username}:{proxy_password}@{proxy_host}:{proxy_port}",
        "https": f"socks5h://{proxy_username}:{proxy_password}@{proxy_host}:{proxy_port}"
    }

    try:
        # Make a request through the SOCKS5 proxy
        response = session.get("https://tools.proxysocks5.com/api/ip/info")

        # Print the IP address returned by the API (should reflect the proxy IP)
        print("IP address using Session with SOCKS5 proxy:", response.content.decode('utf-8'))
    except requests.RequestException as e:
        # Handle any exceptions that occur during the request
        print("Error occurred during the SOCKS5 proxy request with Session:", e)
