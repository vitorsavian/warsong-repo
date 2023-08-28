CC = clang
CFLAGS = -Wall -Wextra
SRC_DIR = .

all: compile clean

compile: main.o utils/utils.o
	$(CC) $(CFLAGS) -o $@ $^

main.o: main.c utils/utils.h
	$(CC) $(CFLAGS) -c -o $@ $<

utils/utils.o: utils/utils.c utils/utils.h
	$(CC) $(CFLAGS) -c -o $@ $<

clean: 
	rm -f *.o utils/*.o
