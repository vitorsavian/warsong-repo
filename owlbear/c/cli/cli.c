#include "cli.h"
#include <stdio.h>

#define ROW "---------------------------------\n"

void cliHelp() 
{
	CliItem *commands[] = {
		&(CliItem){
			"help", 
			"this command help you do the stuff"
		},
		&(CliItem){
			"character",
			"This command begins a new character and you need to pass a little bit of infos to it"
		},
	};

}

int cliVerify() 
{
	return 1;
}
