# Powershell Proxy

[![Build Powershell-Proxy](https://github.com/thomassampson/powershell-proxy/actions/workflows/go.yml/badge.svg)](https://github.com/thomassampson/powershell-proxy/actions/workflows/go.yml)

Lightweight rest api that allows users to run Powershell commands over HTTP. Requests require a valid JWT and responses are returned in JSON format.

## Build

Build the application by running the build.sh script on Linux or WSL:

```bash
$ ./build.sh

[BUILD START] 🔥 Building Powershell Proxy - Version: 0.0.1639180082
[BUILD] 🔵 Cleaning Build Directory ./build
[BUILD] 🟢 Build Directory Cleaned
[BUILD] 🔵 Compiling Windows Binary
[BUILD] 🟢 Windows Binary Compiled to ./build/win/powershell-proxy_0.0.1639180082
[BUILD] 🔵 Compiling Linux Binary
[BUILD] 🟢 Linux Binary Compiled to ./build/linux/powershell-proxy_0.0.1639180082
[BUILD] ⬇️  Binaries Successfully Created

./build:
linux  win

./build/linux:
powershell-proxy_0.0.1639180082

./build/win:
powershell-proxy_0.0.1639180082.exe

[BUILD SUCCESS] ✅ Built Powershell Proxy | Version: '0.0.1' | Build Time: '1 sec'
```

## Environment Variables

To run this application, you will need to add the following environment variables:

`PWSHPRXY_LISTEN_ADDR` - optional - default: 0.0.0.0

`PWSHPRXY_LISTEN_PORT` - optional - default: 8000

`PWSHPRXY_TYPE` - required - valid values: "core" or "windows"

`PWSHPRXY_OKTA_CLIENT_ID` - required

`PWSHPRXY_OKTA_ISSUER` - required

`PWSHPRXY_OKTA_AUDIENCE` - required

## Run

To run the application ensure all the required environment variables are set then execute the binary:

#### Windows

```cmd
C:\> powershell-proxy.exe
```

#### Linux

```bash
$ ./powershell-proxy
```

#### Output

```
2021/12/10 19:02:40 🔵 Starting
2021/12/10 19:02:40 INFO: Env Variable 'PWSHPRXY_LISTEN_ADDR' not set, defaulting to 0.0.0.0
2021/12/10 19:02:40 INFO: IP Address: 0.0.0.0 is Valid
2021/12/10 19:02:40 INFO: Env Variable 'PWSHPRXY_APP_NAME' not set, defaulting to 8000
2021/12/10 19:02:40 INFO: Using Powershell Type: pwsh
2021/12/10 19:02:40 INFO: Using AppName: Powershell Proxy API
2021/12/10 19:02:40 INFO: Using ListenPort: 8888
2021/12/10 19:02:40 INFO: Using ListenAddress: 0.0.0.0
2021/12/10 19:02:40 INFO: Using OktaClientId: ***********
2021/12/10 19:02:40 INFO: Using OktaAudience: api://default
2021/12/10 19:02:40 INFO: Using OktaIssuer: https://tenant.okta.com/oauth2/default
2021/12/10 19:02:40 🟢 Started Powershell Proxy API
```

## Usage/Examples

### Get API Info

```http
  GET /api
```

#### Example Requests

curl

```bash
curl -X GET \
  'http://localhost:8000/api'
```

python

```python
import requests

reqUrl = "http://localhost:8000/api/"

response = requests.request("GET", reqUrl)

print(response.text)
```

javascript

```js
fetch("http://localhost:8000/api", {
  method: "GET",
})
  .then(function (response) {
    return response.text();
  })
  .then(function (data) {
    console.log(data);
  });
```

#### Example Response

```
✋ Powershell Proxy API
```

### Run Command

```http
  POST /api/command
```

#### Query Parameters

| Parameter | Type  | Description                                                            |
| :-------- | :---- | :--------------------------------------------------------------------- |
| `depth`   | `int` | **Optional**. Set the depth of json responses. Default: 4 (range: 1-6) |

#### Headers Parameters

| Header          | Description                                              |
| :-------------- | :------------------------------------------------------- |
| `Authorization` | **Required**. Valid JWT Access Token generated from Okta |

#### Example Requests

curl

```bash
curl -X POST \
  'http://127.0.0.1:8000/api/command?depth=4' \
  -H 'Authorization: Bearer <JWT>' \
  -H 'Content-Type: application/json' \
  -d '{"commands": ["Get-ChildItem | Select-Object Name"]}'
```

python

```python
import requests

reqUrl = "http://127.0.0.1:8000/api/command?depth=4"

headersList = {
 "Authorization": "Bearer <JWT>",
 "Content-Type": "application/json"
}

payload = "{\n\"commands\":[\"Get-ChildItem | Select-Object Name\"]\n}"

response = requests.request("POST", reqUrl, data=payload,  headers=headersList)

print(response.text)
```

javascript

```javascript
let headersList = {
  Authorization: "Bearer <JWT>",
  "Content-Type": "application/json",
};

fetch("http://127.0.0.1:8000/api/command?depth=4", {
  method: "POST",
  body: '{\n"commands":["Get-ChildItem | Select-Object Name"]\n}',
  headers: headersList,
})
  .then(function (response) {
    return response.text();
  })
  .then(function (data) {
    console.log(data);
  });
```

#### Example Response

```json
[
  {
    "Name": "build"
  },
  {
    "Name": "build.sh"
  },
  {
    "Name": "go.mod"
  },
  {
    "Name": "go.sum"
  },
  {
    "Name": "main.go"
  },
  {
    "Name": "README.md"
  }
]
```

## Author

- [@thomassampson](https://www.github.com/thomassampson)
