import requests

url = "http://localhost:8081/readings/count"

response = requests.get(url)

print(response.json())