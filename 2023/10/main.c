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

uint longest_distance(char **m, size_t y_len, size_t x_len, size_t s,
                      char curr) {
  size_t p = -1;
  char c;
  if (s % x_len != 0 &&
      ((c = m[s / x_len][s % x_len - 1]) == '-' || c == 'F' || c == 'L') &&
      (curr == '-' || curr == 'J' || curr == '7' || curr == 'S')) {
    p = s - 1;
  }
  if (s % x_len != x_len - 1 &&
      ((c = m[s / x_len][s % x_len + 1]) == '-' || c == '7' || c == 'J') &&
      (curr == '-' || curr == 'L' || curr == 'F' || curr == 'S')) {
    p = s + 1;
  }
  if (s >= x_len &&
      ((c = m[s / x_len - 1][s % x_len]) == '|' || c == '7' || c == 'F') &&
      (curr == '|' || curr == 'J' || curr == 'L' || curr == 'S')) {
    p = s - x_len;
  }
  if (s / x_len < y_len - 1 &&
      ((c = m[s / x_len + 1][s % x_len]) == '|' || c == 'L' || c == 'J') &&
      (curr == '|' || curr == 'F' || curr == '7' || curr == 'S')) {
    p = s + x_len;
  }

  printf("p: %ld\n", p);

  if (p != -1) {
    curr = m[p / x_len][p % x_len];
    m[p / x_len][p % x_len] = '.';
    return 1 + longest_distance(m, y_len, x_len, p, curr);
  }
  return 1;
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

  size_t y_size = 64;
  size_t y_len = 0;
  size_t x_size;
  size_t x_len = 0;
  char **m = malloc(sizeof(char *) * y_size);

  size_t start;
  char *start_p = NULL;

  while (getline(&l, &n, f) != -1) {
    trim(l);
    size_t x_size = strlen(l);
    if ((start_p = strchr(l, 'S')) != NULL) {
      start = x_size * y_len + (start_p - l);
    }
    m[y_len] = malloc(sizeof(char) * x_size);
    strncpy(m[y_len], l, x_size);
    y_len++;
    x_len = x_size;
    if (y_len == y_size) {
      y_size *= 2;
      m = realloc(m, sizeof(char *) * y_size);
    }
  }

  for (size_t i = 0; i < y_len; i++) {
    for (size_t j = 0; j < x_len; j++) {
      printf("%c, ", m[i][j]);
    }
    printf("\n");
  }

  printf("start: %ld\n", start);

  res = longest_distance(m, y_len, x_len, start, 'S');

  printf("\n");
  printf("%lld", res / 2);
  printf("\n");
  fclose(f);
  if (l) {
    free(l);
  }
}
