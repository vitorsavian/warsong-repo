#include "utils.h"
#include <string.h>

int isStrEqual(char *str1, char *str2) {
  if (strcmp(str1, str2) == 0) {
    return 1;
  };

  return 0;
}
