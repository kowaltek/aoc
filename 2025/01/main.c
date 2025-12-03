#include <limits.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main() {
  FILE *f;
  if ((f = fopen("input.txt", "r")) == NULL) {
    perror("can't open input.txt");
    exit(1);
  }
  char *l = NULL;
  size_t n = 0;
  uint sum = 0;
  int pos = 50;
  while (getline(&l, &n, f) != -1) {
    char dir = l[0];
    int d = atoi(&l[1]);
    printf("%d\n", d);
    switch (dir) {
    case 'R':
      for (int i = 0; i < d; i++) {
          pos++;
          pos%=100;
        if (pos == 0) {
          sum++;
        }
      }
      break;
    case 'L':
      for (int i = 0; i < d; i++) {
          pos--;
          pos = pos == -1 ? 99 : pos;
        if (pos == 0) {
          sum++;
        }
      }
      break;
    }
    printf("pos: %d\n", pos);
  }
  printf("%d\n", sum);
  fclose(f);
  if (l) {
    free(l);
  }
}
