# server-info
Simple application made to collect IP addresses and name server from a host.

## Help

```
NAME:
   server-info - Search IP Address and Names on internet

USAGE:
   server-info [global options] command [command options] [arguments...]

COMMANDS:
   ips      Search IP Address in internet
   servers  Search Server Names in internet
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

## Exemple usage

```
# ./server-info ips --host google.com

Output:
The following ips were found for the host 'google.com' 
- 2800:3f0:4004:80a::200e 
- 142.250.79.14
```
