#include "./utils/utils.h"
#include <stdio.h>
#include <string.h>

#define EXIT_SUCCESS 0

int main(int argc, char *argv[])
{
	int i;

	// this begins the CLI and make it work
        if (argc == 1 || isStrEqual("1", "1")) {
                printf( "please you need to pass some info\n" );
        } else {
                printf( "Arguments:\n" );

                for ( i = 1; i < argc; ++i ) {
                        printf( "  %d. %s\n", i, argv[i] );
                }
        }

	return EXIT_SUCCESS;
}
