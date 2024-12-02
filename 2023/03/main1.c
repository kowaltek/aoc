#include <ctype.h>
#include <limits.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

/* typedef struct { */
/*   int x, y; */
/*   int ratio1, ratio2; */
/* } gear; */
/*  */
/* gear *get_valid_gears(gear *gears) {} */

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
  /* size_t gear_list_len = 100; */
  /* gear *gear_list = malloc(sizeof(gear) * gear_list_len); */
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
        if (start - 1 >= 0 && m[y][start - 1] != '.') {
          to_add = true;
        }
        if (!to_add && end + 1 < width && m[y][end + 1] != '.') {
          to_add = true;
        }
        if (!to_add && y > 0) {
          for (int c = start - 1 >= 0 ? start - 1 : 0;
               c <= (end + 1 < width ? end + 1 : end); c++) {
            if (m[y - 1][c] != '.') {
              to_add = true;
              break;
            }
          }
        }
        if (!to_add && y < i - 1) {
          for (int c = start - 1 >= 0 ? start - 1 : 0;
               c <= (end + 1 < width ? end + 1 : end); c++) {
            if (m[y + 1][c] != '.') {
              to_add = true;
              break;
            }
          }
        }
        if (to_add) {
          sum += atoi(buf);
        }
      }
    }
  }
  printf("%d\n", sum);
  fclose(f);
  if (l) {
    free(l);
  }
}
