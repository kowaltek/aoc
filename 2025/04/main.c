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
  const size_t s = 135;
  int m[s + 2][s + 2] = {};
  int r[s + 2][s + 2] = {};
  int y = 1;
  while (getline(&l, &n, f) != -1) {
    for (int x = 1; x <= s; x++) {
      if (l[x - 1] == '@') {
        m[y - 1][x - 1]++;
        m[y - 1][x]++;
        m[y - 1][x + 1]++;
        m[y][x - 1]++;
        m[y][x + 1]++;
        m[y + 1][x - 1]++;
        m[y + 1][x]++;
        m[y + 1][x + 1]++;
        r[y][x] = 1;
      }
    }
    y++;
  }
  unsigned long sum = 0;
  unsigned long prev_sum = -1;
  while (sum != prev_sum) {
    prev_sum = sum;
    for (int y = 1; y <= s; y++) {
      for (int x = 1; x <= s; x++) {
        if (m[y][x] < 4 && r[y][x] == 1) {
          sum++;
          r[y][x] = 0;
        }
      }
    }
    for (int y = 0; y < s + 2; y++) {
        for (int x = 0; x < s + 2; x++) {
            m[y][x] = 0;
        }
    }
    for (int y = 1; y <= s; y++) {
      for (int x = 1; x <= s; x++) {
        if (r[y][x]) {
          m[y - 1][x - 1]++;
          m[y - 1][x]++;
          m[y - 1][x + 1]++;
          m[y][x - 1]++;
          m[y][x + 1]++;
          m[y + 1][x - 1]++;
          m[y + 1][x]++;
          m[y + 1][x + 1]++;
          r[y][x] = 1;
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
