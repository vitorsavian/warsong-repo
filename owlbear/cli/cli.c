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
          "npc",
          "This command let you work with npcs, you can create and so on"}};

  printf("%s", ROW);
  cliTitle();

  for (int i = 0; i < COMMAND_LENGTH; i++) {
    printf("\tCommand: %s\n\tDescription: %s\n\n", commands[i]->command,
           commands[i]->description);
  };
}

void cliTitle() {
  printf(" OWLBEAR\n");
  printf(" A CLI app for handling and managing Dungeons and "
         "Dragons things\n\n");
  return;
}

int cliVerify() { return 1; }
