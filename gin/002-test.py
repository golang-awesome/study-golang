import requests

resp = requests.get("http://localhost:8080/home")
print(resp.json())
login = requests.get("http://localhost:8080/auth/signin")
print(login.cookies)
resp = requests.get("http://localhost:8080/home", cookies=login.cookies)
print(resp.json())

