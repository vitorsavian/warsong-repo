#ifndef CLI_H_
#define CLI_H_

typedef struct {
  char *command;
  char *description;
} CliItem;

void cliHelp();
int cliVerify();

#endif // !CLI_H_
