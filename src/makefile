C=go
TAGS=-tags static
LDFLAGS=-ldflags "-H=windowsgui"
TARG=bricks
BIN=..//bin//$(TARG).exe

all: build

build:
	$(C) build $(TAGS) $(LDFLAGS) -o $(BIN)

clean:
	rm bin//main.exe