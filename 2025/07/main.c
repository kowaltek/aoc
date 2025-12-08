#include <ctype.h>
#include <limits.h>
#include <math.h>
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
  size_t n;
  unsigned long sum = 0;
  int beams[256] = {0};
  while (getline(&l, &n, f) != -1) {
    for (int i = 0; i < n; i++) {
      switch (l[i]) {
      case 'S':
        beams[i] = 1;
        break;
      case '^':
        if (beams[i] == 1) {
          sum++;
          if (i > 0) {
            beams[i - 1] = 1;
          }
          beams[i + 1] = 1;
          beams[i] = 0;
        }
      }
    }
  }
  printf("%ld\n", sum);
  fclose(f);
  if (l) {
    free(l);
  }
}
