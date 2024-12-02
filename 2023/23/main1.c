#include <ctype.h>
#include <limits.h>
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

int solve(char **m, size_t y_len, size_t x_len, int y, int x, int curr_dist) {
  if (y == y_len - 1 && x == x_len - 2) {
    return curr_dist;
  }

  int tmp;
  int res = -1;
  char c = m[y][x];
  m[y][x] = '#';
  if (x - 1 >= 0 && m[y][x - 1] != '#' &&
      (m[y][x - 1] == '.' || m[y][x - 1] == '<')) {
    res = solve(m, y_len, x_len, y, x - 1, curr_dist + 1);
  }

  if (x + 1 < x_len && m[y][x + 1] != '#' &&
      (m[y][x + 1] == '.' || m[y][x + 1] == '>')) {
    tmp = solve(m, y_len, x_len, y, x + 1, curr_dist + 1);
    res = res == -1 ? tmp : res < tmp ? tmp : res;
  }

  if (y - 1 >= 0 && m[y - 1][x] != '#' &&
      (m[y - 1][x] == '.' || m[y - 1][x] == '^')) {
    tmp = solve(m, y_len, x_len, y - 1, x, curr_dist + 1);
    res = res == -1 ? tmp : res < tmp ? tmp : res;
  }

  if (y + 1 < y_len && m[y + 1][x] != '#' &&
      (m[y + 1][x] == '.' || m[y + 1][x] == 'v')) {
    tmp = solve(m, y_len, x_len, y + 1, x, curr_dist + 1);
    res = res == -1 ? tmp : res < tmp ? tmp : res;
  }
  m[y][x] = c;

  return res != -1 ? res : curr_dist;
}

int main() {
  FILE *f;
  if ((f = fopen("input.txt", "r")) == NULL) {
    perror("can't open input.txt");
    exit(1);
  }

  unsigned long long res = 0;
  char *l = NULL;
  char *line = NULL;
  size_t n = 0;

  size_t y_size = 8;
  size_t y_len = 0;
  size_t x_len = 0;
  char **m = malloc(sizeof(char *) * y_size);

  while (getline(&l, &n, f) != -1) {
    line = l;
    trim(line);

    x_len = strlen(line);
    m[y_len] = malloc(sizeof(char) * x_len);
    strcpy(m[y_len], line);
    y_len++;

    if (y_len == y_size) {
      y_size *= 2;
      m = realloc(m, sizeof(char *) * y_size);
    }
  }

  res = solve(m, y_len, x_len, 0, 1, 0);

  for (int i = 0; i < y_len; i++) {
    for (int j = 0; j < x_len; j++) {
      printf("%c", m[i][j]);
    }
    printf("\n");
  }

  printf("\n");
  printf("\n");
  printf("%lld\n", res);
  printf("\n");
  fclose(f);
  if (l) {
    free(l);
  }
}
