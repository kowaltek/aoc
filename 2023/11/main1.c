#include <ctype.h>
#include <math.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
  size_t x;
  size_t y;
} Position;

void trim(char *s) {
  char c;
  while ((c = *s++) != '\0') {
    if (c == '\n' || c == ' ') {
      *(s - 1) = '\0';
      break;
    }
  }
}

void expand_universe(Position *galaxies, size_t galaxies_len, size_t *x_size,
                     size_t *y_size) {

  size_t *empty_rows = malloc(sizeof(size_t) * (*y_size));
  size_t empty_rows_i = 0;
  size_t *empty_cols = malloc(sizeof(size_t) * (*x_size));
  size_t empty_cols_i = 0;

  size_t d = 0;
  for (size_t i = 0; i < *y_size; i++) {
    bool is_empty = true;
    for (size_t j = 0; j < galaxies_len; j++) {
      if (galaxies[j].y == i) {
        is_empty = false;
        break;
      }
    }
    if (is_empty) {
      empty_rows[empty_rows_i++] = i + d++;
    }
  }

  d = 0;
  for (size_t i = 0; i < *x_size; i++) {
    bool is_empty = true;
    for (size_t j = 0; j < galaxies_len; j++) {
      if (galaxies[j].x == i) {
        is_empty = false;
        break;
      }
    }
    if (is_empty) {
      empty_cols[empty_cols_i++] = i + d++;
    }
  }

  for (size_t i = 0; i < empty_rows_i; i++) {
    for (size_t j = 0; j < galaxies_len; j++) {
      if (galaxies[j].y > empty_rows[i]) {
        galaxies[j].y++;
      }
    }
  }

  for (size_t i = 0; i < empty_cols_i; i++) {
    for (size_t j = 0; j < galaxies_len; j++) {
      if (galaxies[j].x > empty_cols[i]) {
        galaxies[j].x++;
      }
    }
  }
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

  size_t galaxies_size = 32;
  size_t galaxies_len = 0;
  Position *galaxies = malloc(sizeof(Position) * galaxies_size);

  while (getline(&l, &n, f) != -1) {
    trim(l);
    size_t x_size = strlen(l);
    m[y_len] = malloc(sizeof(char) * x_size);
    strncpy(m[y_len], l, x_size);
    for (int x = 0; x < x_size; x++) {
      if (m[y_len][x] == '#') {
        galaxies[galaxies_len++] = (Position){.x = x, .y = y_len};
        if (galaxies_len == galaxies_size) {
          galaxies_size *= 2;
          galaxies = realloc(galaxies, sizeof(Position) * galaxies_size);
        }
      }
    }
    y_len++;
    x_len = x_size;
    if (y_len == y_size) {
      y_size *= 2;
      m = realloc(m, sizeof(char *) * y_size);
    }
  }

  for (size_t i = 0; i < y_len; i++) {
    for (size_t j = 0; j < x_len; j++) {
      printf("%c", m[i][j]);
    }
    printf("\n");
  }
  printf("\n");

  for (size_t i = 0; i < galaxies_len; i++) {
    printf("x: %ld, y: %ld; ", galaxies[i].x, galaxies[i].y);
  }
  printf("\n");
  printf("\n");

  expand_universe(galaxies, galaxies_len, &x_len, &y_len);

  for (size_t i = 0; i < galaxies_len; i++) {
    printf("x: %ld, y: %ld; ", galaxies[i].x, galaxies[i].y);
  }
  printf("\n");

  for (int i = 0; i < galaxies_len - 1; i++) {
    for (int j = i + 1; j < galaxies_len; j++) {
      Position g1 = galaxies[i];
      Position g2 = galaxies[j];
      uint d = galaxies[i].x > galaxies[j].x ? galaxies[i].x - galaxies[j].x
                                             : galaxies[j].x - galaxies[i].x;
      d += galaxies[i].y > galaxies[j].y ? galaxies[i].y - galaxies[j].y
                                         : galaxies[j].y - galaxies[i].y;
      res += d;
    }
  }

  printf("\n");
  printf("%lld", res);
  printf("\n");
  fclose(f);
  if (l) {
    free(l);
  }
}
