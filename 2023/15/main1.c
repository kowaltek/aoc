#include <ctype.h>
#include <math.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void trim(char *s) {
  char c;
  while ((c = *s++) != '\0') {
    if (c == '\n' || c == ' ') {
      *(s - 1) = '\0';
      break;
    }
  }
}

int hash(char *s, size_t len) {
  int res = 0;
  for (int i = 0; i < len; i++) {
    res += s[i];
    res = res * 17 % 256;
  }
  return res;
}

int main() {
  FILE *f;
  if ((f = fopen("input.txt", "r")) == NULL) {
    perror("can't open input.txt");
    exit(1);
  }

  long long res = 0;
  char *l = NULL;
  size_t n = 0;
  char *line = NULL;

  while (getline(&l, &n, f) != -1) {
    line = l;

    while (*line != 0 && *line != '\n') {
      int len = 0;
      while (line[len] != ',' && line[len] != 0 && line[len] != '\n') {
        len++;
      }
      printf("hashing at: %d\n", len);
      res += hash(line, len);
      line += len + 1;
    }

    res += 0;
  }

  printf("\n");
  printf("%lld", res);
  printf("\n");
  fclose(f);
  if (l) {
    free(l);
  }
}
