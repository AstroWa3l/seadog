# seadog
CLI AI Chat Bot for SPOs

## Requirements
- Go 1.20 or higher
- [Mendable API Key](https://mendable.ai/) (You can sign up for free)

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
nano .env
```
`MENDABLE_API_KEY=YOUR_MENDABLE_API_KEY`

3. Example of how to run the program to start chatting with the bot

(option 1) Run the program using go run

```bash
go run seadog.go -cmd ask
```

(option 2) Build the executable and run it

```bash
go build seadog.go
./seadog -cmd ask
```

(option 3) Build the executable and run it from anywhere (Caution when doing this with any executable XD)

```bash
go build seadog.go
sudo cp seadog /usr/local/bin
sudo cp .env /usr/local/bin
seadog -cmd ask
```
