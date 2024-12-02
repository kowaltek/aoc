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

  size_t row_len = 0;
  size_t row_size = 32;
  uint *row = malloc(sizeof(uint) * row_size);

  size_t col_len = 0;
  size_t col_size = 32;
  uint *col = malloc(sizeof(uint) * col_size);

  for (int i = 0; i < col_size; i++) {
    col[i] = 0;
  }
  for (int i = 0; i < row_size; i++) {
    row[i] = 0;
  }

  while (getline(&l, &n, f) != -1) {
    line = l;

    if (strchr(line, '.') != NULL) {
      // add to stored map
      col_len = 0;
      char c;
      while ((c = *line++) == '.' || c == '#') {
        col[col_len] += c == '.' ? 0 : 1 << row_len;
        row[row_len] += c == '.' ? 0 : 1 << col_len;
        col_len++;
      }
      row_len++;
    } else {
      // process the map
      printf("cols\n");
      for (int i = 0; i < col_len - 1; i++) {
        int j = 0;
        int tmp = 0;
        while (i - j >= 0 && i + 1 + j < col_len &&
               col[i - j] == col[i + 1 + j]) {
          tmp++;
          j++;
        }
        if (tmp != 0 && (i - j == -1 || i + 1 + j == col_len)) {
          res += i + 1;
        }
        printf("%d, ", col[i]);
      }
      printf("\n");
      printf("rows\n");
      for (int i = 0; i < row_len; i++) {
        int j = 0;
        int tmp = 0;
        while (i - j >= 0 && i + 1 + j < row_len &&
               row[i - j] == row[i + 1 + j]) {
          tmp++;
          j++;
        }
        if (tmp != 0 && (i - j == -1 || i + 1 + j == row_len)) {
          res += (i + 1) * 100;
        }
        printf("%d, ", row[i]);
      }
      printf("\n");

      for (int i = 0; i < col_size; i++) {
        col[i] = 0;
      }
      for (int i = 0; i < row_size; i++) {
        row[i] = 0;
      }
      col_len = 0;
      row_len = 0;
    }

    res += 0;
  }

  printf("cols\n");
  for (int i = 0; i < col_len - 1; i++) {
    int j = 0;
    int tmp = 0;
    while (i - j >= 0 && i + 1 + j < col_len && col[i - j] == col[i + 1 + j]) {
      tmp++;
      j++;
    }
    if (tmp != 0 && (i - j == -1 || i + 1 + j == col_len)) {
      res += i + 1;
    }
    printf("%d, ", col[i]);
  }
  printf("\n");
  printf("rows\n");
  for (int i = 0; i < row_len; i++) {
    int j = 0;
    int tmp = 0;
    while (i - j >= 0 && i + 1 + j < row_len && row[i - j] == row[i + 1 + j]) {
      tmp++;
      j++;
    }
    if (tmp != 0 && (i - j == -1 || i + 1 + j == row_len)) {
      res += (i + 1) * 100;
    }
    printf("%d, ", row[i]);
  }
  printf("\n");

  printf("\n");
  printf("%lld", res);
  printf("\n");
  fclose(f);
  if (l) {
    free(l);
  }
}
