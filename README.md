[![CI](https://github.com/algo7/TF2-RCON-MISC/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/algo7/TF2-RCON-MISC/actions/workflows/ci.yml)
# TF2-RCON-MISC
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
![Launch Options](https://github.com/algo7/TF2-RCON-MISC/blob/main/launch_options.png?raw=true)

The password really doesn't matter as nobody will be accessing it except for you. At the moment the password use to connect to RCON is hardcoded as `123` so please don't change it; otherwise, the program will not work.

### (optional) ChatGPT support:
If you want your client to respond to !gpt commands you need to get an openai api key.
Go to openai page and create an API-Key: https://platform.openai.com/account/api-keys

Edit run.bat and set the key into environment variable **OPENAI_APIKEY**.

## Test

```
go test -v utils/utils_test.go 
```

## Misc

### Counterstrike, Dystopia
To run this app using other source relevant games like dystopia, use the runDystopia.bat, or runCounterstrike.bat.
There it is just about the console.log path, the app is capable of detecting dystopia consolelog. Counterstrike is only barely working.