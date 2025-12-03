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
  size_t n = 0;
  unsigned long sum = 0;
  while (getline(&l, &n, f) != -1) {
    int len = strlen(l);
    int digit_place = 12;
    int d_i = 0;
    while (digit_place) {
      int d = 0;
      for (int i = d_i; i < len - digit_place; i++) {
        if (l[i] - 48 > d) {
          d_i = i;
          d = l[i] - 48;
        }
      }
      sum += d * (unsigned long)pow(10, digit_place - 1);
      digit_place--;
      d_i++;
    }
  }
  printf("%ld\n", sum);
  fclose(f);
  if (l) {
    free(l);
  }
}
