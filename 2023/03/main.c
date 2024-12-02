#include <ctype.h>
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
  int m_size = 20;
  char **m = malloc(sizeof(char *) * m_size);
  int i = 0;
  while (getline(&l, &n, f) != -1) {
    m[i] = malloc(sizeof(char) * (strlen(l) + 1));
    strcpy(m[i++], l);
    if (i == m_size) {
      m_size *= 2;
      m = realloc(m, sizeof(char *) * m_size);
    }
  }
  // -1 for \n
  int width = strlen(m[0]) - 1;
  // num of neighboring numbers
  int *gear_list = calloc(width * i, sizeof(int));
  // stores the multiplication of neighboring numbers
  int *gear_ratio_list = malloc(sizeof(int) * width * i);
  for (int n = 0; n < width * i; n++) {
    gear_ratio_list[n] = 1;
  }
  for (int y = 0; y < i; y++) {
    for (int x = 0; x < width; x++) {
      if (isdigit(m[y][x])) {
        int start = x;
        int end = start;
        char buf[10];
        int c = 0;
        while (isdigit(m[y][end])) {
          buf[c++] = m[y][end++];
        }
        buf[c] = '\0';
        x = end;
        end--;
        bool to_add = false;
        int num = atoi(buf);
        if (start - 1 >= 0 && m[y][start - 1] != '.') {
          if (m[y][start - 1] == '*') {
            gear_list[y * width + (start - 1)]++;
            gear_ratio_list[y * width + (start - 1)] *= num;
          }
          to_add = true;
        }
        if (!to_add && end + 1 < width && m[y][end + 1] != '.') {
          if (m[y][end + 1] == '*') {
            gear_list[y * width + (end + 1)]++;
            gear_ratio_list[y * width + (end + 1)] *= num;
          }
          to_add = true;
        }
        if (!to_add && y > 0) {
          for (int c = start - 1 >= 0 ? start - 1 : 0;
               c <= (end + 1 < width ? end + 1 : end); c++) {
            if (m[y - 1][c] != '.') {
              if (m[y - 1][c] == '*') {
                gear_list[(y - 1) * width + (c)]++;
                gear_ratio_list[(y - 1) * width + (c)] *= num;
              }
              to_add = true;
              break;
            }
          }
        }
        if (!to_add && y < i - 1) {
          for (int c = start - 1 >= 0 ? start - 1 : 0;
               c <= (end + 1 < width ? end + 1 : end); c++) {
            if (m[y + 1][c] != '.') {
              if (m[y + 1][c] == '*') {
                gear_list[(y + 1) * width + (c)]++;
                gear_ratio_list[(y + 1) * width + (c)] *= num;
              }
              to_add = true;
              break;
            }
          }
        }
      }
    }
  }
  for (int n = 0; n < width * i; n++) {
    if (gear_list[n] == 2) {
      sum += gear_ratio_list[n];
    }
  }
  printf("%d\n", sum);
  fclose(f);
  if (l) {
    free(l);
  }
}
