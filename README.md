# JWT CLI
![JWT CLI](./terminal.svg)

Decode/Verify and generate JWT from your terminal

## Installation

```bash
brew tap felixvo/jwtcli
brew install jwtcli
```

## Usage

### Decode

If you just want to check payload.

```bash
jwtcli your-token
```

Or with `secret`

``` bash
jwtcli --secret your-secret token
```

Ex:

```bash
jwtcli -s your-256-bit-secret your-token
HEADER
alg:  HS256
typ:  JWT
PAYLOAD
{
 "iat": 1516239022,
 "name": "John Doe",
 "sub": "1234567890"
}
INFO
iat:  1516239022 formated:2018-01-18 08:30:22
SIG
SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c <nil>
```

### Encode

```bash
jwtcli --secret you-secret sign --payload '{"user_id":6,"exp":1587971056}'
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODc5NzEwNTYsImlhdCI6MTU2NzY4MjM0MywidXNlcl9pZCI6Nn0.LtMr-_nZQCukKi6y4XTGGHdmzo8rDW20BnAdDyfLxTc
```
