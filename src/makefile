C=go
TAGS=-tags static
CFLAGS=
LDFLAGS=-ldflags
TARG=bricks
BIN=..//bin//$(TARG).exe

ifeq ($(OS),Windows_NT)
# check CGO_LDFLAGS environment variable has been set
	ifeq ($(CGO_LDFLAGS),)
# set CGO_LDFLAGS environment variable
		$(shell $env:CGO_LDFLAGS = "-static-libgcc -static -lpthread")
		CFLAGS += -a -x
	endif
	LDFLAGS += "-H=windowsgui"
else
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Linux)
		LDFLAGS += ""
	endif
	ifeq ($(UNAME_S),Darwin)
		LDFLAGS += ""
	endif
endif

all: build

build:
	$(C) build $(TAGS) $(CFLAGS) $(LDFLAGS) -o $(BIN)

clean:
	rm bin//main.exe