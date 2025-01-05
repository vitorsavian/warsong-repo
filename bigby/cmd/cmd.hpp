#ifndef CMD_H_
#define CMD_H_

typedef struct {
  char *command;
  char *description;
} CliItem;

void cliHelp();
void cliTitle();
int cliVerify();

#endif // !CMD_H_
