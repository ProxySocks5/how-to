import requests
from requests import Session

# Proxy credentials and host configuration
proxy_username = 'change_with_your_proxy_username'  # TODO Replace with your proxy username
proxy_password = 'change_with_your_proxy_password'  # TODO Replace with your actual proxy password
proxy_host_port = '45.43.134.162:43222'  # TODO Replace with your HTTP proxy IP and Port in the format IP:Port

# Set up the proxies with embedded credentials
proxies = {
    "http": f"http://{proxy_username}:{proxy_password}@{proxy_host_port}",
    "https": f"http://{proxy_username}:{proxy_password}@{proxy_host_port}"
}

# Use a session to maintain the settings across requests
with Session() as session:
    # Configure session to use the proxy settings with credentials embedded
    session.proxies.update(proxies)

    try:
        # Make a request through the HTTP proxy
        response = session.get("https://tools.proxysocks5.com/api/ip/info")

        # Print the IP address returned by the API (should reflect the proxy IP)
        print("IP address using HTTP proxy:", response.content.decode('utf-8'))
    except requests.RequestException as e:
        # Handle any exceptions that occur during the request
        print("Error occurred during the HTTP proxy request with Session:", e)
