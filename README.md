## JWT CLI 
Decode/Verify and generate JWT from your terminal

## Installation
### Mac
```
```
### Linux
```
```

## Usage
### Decode
If you just want to check payload.  
```
$ jwtcli {token}
```
  
Or with `secret`  
```
$ jwtcli --secret your-secret token
```

Ex:  
<img src="https://i.imgur.com/E2UXAix.png" alt="JWT CLI" width="600" height="300"/>

### Generate
```
$jwtcli --secret you-secret sign --payload '{"user_id":6,"exp":1587971056}'
```
Ex:  
<img src="https://i.imgur.com/4PlBRBf.png" alt="Generate JWT" width="600" height="300"/>

