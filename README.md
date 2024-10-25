# Token Scanner

It is basic port scanner.

# How To Install

```bash
git clone https://github.com/EfeYorutken/token_scanner.git
cd token_scanner
go build -o tokenscanner
./tokenscanner
```

# How To Use

You can call the binary file without any parameters to see the help menu

## Arguments

- `./tokenscanner <arg1> <params1> <arg2> <params2> ...`

- `-p port1 port2 port3 ...`
    - stands for port
    - scans the specified ports of the target

- `-r <begin port> <end port>`
    - stands for range
    - scans all the ports of the target between `<begin port>` and `<end port>`

- `-t <port>`
    - stands for type
    - specifies which protocol will be used in the scan
    - you can check [this](https://pkg.go.dev/net#Dial) link for the list of values that can be entered

- `-sG <command that will run the script>`
    - stands for script Good
    - when a scan is successfull, will run the command specified using the values given
    - the command that will run the script  will be ran with the address port and protocol values of the scan in that order
    - `<command that will run the script> <succesfull address> <succesfull port> <succesfull protocol>`

- `-sB <command that will run the script>`
    - stands for script Bad
    - same as `-sG` but for unsuccessfull scans

- `-s <command that will run the script>`
    - stands for script
    - same as `-sG` but for all scans
