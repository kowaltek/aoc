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

size_t process_cols(const uint *col, size_t col_len, size_t prev) {
  for (int i = 0; i < col_len - 1; i++) {
    int j = 0;
    int tmp = 0;
    while (i - j >= 0 && i + 1 + j < col_len && col[i - j] == col[i + 1 + j]) {
      tmp++;
      j++;
    }
    if (tmp != 0 && (i - j == -1 || i + 1 + j == col_len) && i + 1 != prev) {
      return i + 1;
    }
  }
  return 0;
}

size_t process_rows(const uint *row, size_t row_len, size_t prev) {
  for (int i = 0; i < row_len - 1; i++) {
    int j = 0;
    int tmp = 0;
    while (i - j >= 0 && i + 1 + j < row_len && row[i - j] == row[i + 1 + j]) {
      tmp++;
      j++;
    }
    if (tmp != 0 && (i - j == -1 || i + 1 + j == row_len) && i + 1 != prev) {
      return i + 1;
    }
  }
  return 0;
}

size_t process_map(char **m, uint *row, size_t row_size, uint *col,
                   size_t col_size, size_t prev) {
  for (int i = 0; i < row_size; i++) {
    row[i] = 0;
  }
  for (int i = 0; i < col_size; i++) {
    col[i] = 0;
  }
  size_t row_len = 0;
  size_t col_len = 0;

  for (int i = 0; i < row_size; i++) {
    col_len = 0;
    for (int j = 0; j < col_size; j++) {
      char c = m[i][j];
      col[col_len] += c == '.' ? 0 : 1 << row_len;
      row[row_len] += c == '.' ? 0 : 1 << col_len;
      col_len++;
    }
    row_len++;
  }

  size_t col_res = process_cols(col, col_len, prev);
  if (col_res != 0 && col_res != prev) {
    return col_res;
  }
  size_t row_res = process_rows(row, row_len, prev / 100);
  return row_res * 100;
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

  size_t m_len = 0;
  size_t m_size = 32;
  char **m = malloc(sizeof(char *) * m_size);
  size_t s_size = 0;

  while (getline(&l, &n, f) != -1) {
    line = l;

    if (strchr(line, '.') != NULL || strchr(line, '#') != NULL) {
      // add to stored map
      s_size = strlen(line);
      char *s = malloc(sizeof(char) * s_size);
      strcpy(s, line);
      s[--s_size] = 0;
      m[m_len++] = s;
    } else {
      // process the map
      uint *row = malloc(sizeof(uint) * m_len);
      uint *col = malloc(sizeof(uint) * s_size);
      size_t res1 = process_map(m, row, m_len, col, s_size, 0);
      printf("res1: %ld\n", res1);
      bool done = false;
      for (int i = 0; i < m_len; i++) {
        if (done) {
          break;
        }
        for (int j = 0; j < s_size; j++) {
          m[i][j] = m[i][j] == '.' ? '#' : '.';
          /* for (int k = 0; k < m_len; k++) { */
          /*   for (int l = 0; l < s_size; l++) { */
          /*     printf("%c", m[k][l]); */
          /*   } */
          /*   printf("\n"); */
          /* } */
          size_t res2 = process_map(m, row, m_len, col, s_size, res1);
          m[i][j] = m[i][j] == '.' ? '#' : '.';
          if (res1 != res2 && res2 > 0) {
            printf("res2: %ld\n", res2);
            res += res2;
            done = true;
            break;
          }
        }
      }
      for (int i = 0; i < m_len; i++) {
        free(m[i]);
      }
      m_len = 0;
    }

    res += 0;
  }

  // process the map
  uint *row = malloc(sizeof(uint) * m_len);
  uint *col = malloc(sizeof(uint) * s_size);
  size_t res1 = process_map(m, row, m_len, col, s_size, 0);
  printf("res1: %ld\n", res1);
  bool done = false;
  for (int i = 0; i < m_len; i++) {
    if (done) {
      break;
    }
    for (int j = 0; j < s_size; j++) {
      m[i][j] = m[i][j] == '.' ? '#' : '.';
      /* for (int k = 0; k < m_len; k++) { */
      /*   for (int l = 0; l < s_size; l++) { */
      /*     printf("%c", m[k][l]); */
      /*   } */
      /*   printf("\n"); */
      /* } */
      size_t res2 = process_map(m, row, m_len, col, s_size, res1);
      m[i][j] = m[i][j] == '.' ? '#' : '.';
      if (res1 != res2 && res2 > 0) {
        printf("res2: %ld\n", res2);
        res += res2;
        done = true;
        break;
      }
    }
  }
  for (int i = 0; i < m_len; i++) {
    free(m[i]);
  }
  m_len = 0;

  printf("\n");
  printf("%lld", res);
  printf("\n");
  fclose(f);
  if (l) {
    free(l);
  }
}
