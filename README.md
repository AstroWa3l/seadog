# seadog
CLI based AI Chat Bot for SPOs

## Requirements
- Go 1.20 or higher

### Installing Go on Raspberry Pi4
If you need to install Go and using a raspberry pi 3 or 4 computer, you can use snap.

1. Install snapd

```bash
sudo apt update
sudo apt install snapd
```

2. Install Go

```bash
sudo snap install go --classic
```

3. Check Go version

```bash
go version
```

## Installation
1. Clone the repository

```bash
git clone https://github.com/AstroWa3l/seadog.git
```

2. Create a .env file and add the following variable in it
```bash
mkdir .env
```
`MENDABLE_API_KEY=YOUR_MENDABLE_API_KEY`

3. Build the executable

```bash
go build seadog.go
```

4. Run the executable and ask for help to find the commands

```bash
./seadog -h
```