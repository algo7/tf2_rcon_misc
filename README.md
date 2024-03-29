[![CI](https://github.com/algo7/tf2_rcon_misc/actions/workflows/ci.yml/badge.svg)](https://github.com/algo7/tf2_rcon_misc/actions/workflows/ci.yml)
# Prerequisite
1. MongoDB installed locally: https://www.mongodb.com/try/download/community
2. OpenAI API Key (if you want to use the ChatGPT integration): https://platform.openai.com/account/api-keys

# github.com/algo7/tf2_rcon_misc
Go program that performs various commands via RCON base on local TF2 console output.
Get all the players' name and steamID on the server and store them into the local MongoDB instance.

# Build and run
## Linux
```bash
$ ./auto_build.sh
$ ./run.sh
```

## Windows
```bash
$ ./auto_build.sh
$ ./run.bat
```

## Configuration

### (required) Launch Options:
---
You should launch your TF2 with the following launch options:
```bash
-flushlog -rpt -novid -usercon -ip 0.0.0.0 +rcon_password 123 +net_start
```
![Launch Options](https://raw.githubusercontent.com/algo7/tf2_rcon_misc/main/launch_options.png?raw=true)

The password really doesn't matter as nobody will be accessing it except for you. At the moment the password use to connect to RCON is hardcoded as `123` so please don't change it; otherwise, the program will not work.