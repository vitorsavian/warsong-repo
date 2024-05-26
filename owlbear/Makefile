CC = clang
CFLAGS = -Wall -Wextra
SRC_DIR = .

all: owlbear clean

owlbear: main.o utils/utils.o cli/cli.o
	$(CC) $(CFLAGS) -o $@ $^

main.o: main.c
	$(CC) $(CFLAGS) -c -o $@ $<

utils/utils.o: utils/utils.c
	$(CC) $(CFLAGS) -c -o $@ $<

cli/cli.o: cli/cli.c
	$(CC) $(CFLAGS) -c -o $@ $<

clean: 
	rm -f *.o utils/*.o cli/*.o
