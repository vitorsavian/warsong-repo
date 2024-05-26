#include "cli.h"
#include <stdio.h>

#define ROW "---------------------------------\n"
#define COMMAND_LENGTH 3

void cliHelp() {
  CliItem *commands[] = {
      &(CliItem){"help", "this command help you do the stuff"},
      &(CliItem){"character", "This command let you work with characters, you "
                              "can create and so on"},
      &(CliItem){
          "monster",
          "This command let you work with monsters, you can create and so on"}};

  printf("%s\n", ROW);

  for (int i = 0; i < COMMAND_LENGTH; i++) {
    printf("\tCommand: %s\n\tDescription: %s\n", commands[i]->command,
           commands[i]->description);
  };
}

int cliVerify() { return 1; }
