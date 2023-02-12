# TF2-RCON-MISC
Go program that performs various commands via RCON base on local TF2 console output.

### Launch Options:
---
You should launch your TF2 with the following launch options:
```bash
-conclearlog -console -novid -condebug -usercon -ip 0.0.0.0 +rcon_password 123 +net_start
```
![Launch Options](https://github.com/algo7/TF2-RCON-MISC/blob/main/launch_options.png?raw=true)

The password really doesn't matter as nobody will be accessing it except for you. At the moment the password use to connect to RCON is hardcoded as `123` so please don't change it; otherwise, the program will not work.
