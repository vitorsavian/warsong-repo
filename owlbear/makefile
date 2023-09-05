CC = clang
CFLAGS = -Wall -Wextra
SRC_DIR = .

all: compile clean

compile: main.o utils
	$(CC) $(CFLAGS) -o $@ $^

main.o: main.c utils
	$(CC) $(CFLAGS) -c -o $@ $<

utils/utils.o: utils utils
	$(CC) $(CFLAGS) -c -o $@ $<

clean: 
	rm -f *.o utils/*.o
