TARGET=main

all:
	make install
	go build -o $(TARGET)

install:
	go install

clean:
	rm -f $(TARGET)