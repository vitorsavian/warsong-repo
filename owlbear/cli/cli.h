#ifndef CLI_H_
#define CLI_H_

typedef struct {
  char *command;
  char *description;
} CliItem;

void cliHelp();
void cliTitle();
int cliVerify();

#endif // !CLI_H_
