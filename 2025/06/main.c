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
  int m[16][2048];
  int y_len = 0;
  int x_len = 0;
  char s[2048];
  int s_len = 0;
  unsigned long sum = 0;
  while (getline(&l, &n, f) != -1) {
    int x_i = 0;
    int i = 0;
    int r_i = 0;
    int b_i = 0;
    char buf[64] = {0};
    bool is_sign_line = false;
    while (!isdigit(l[i])) {
      if (l[i] == '+' || l[i] == '*') {
        is_sign_line = true;
        break;
      }
      i++;
    }
    if (!is_sign_line) {
      bool in_num;
      while (l[i] != '\n') {
        in_num = isdigit(l[i]);
        if (in_num) {
          buf[b_i++] = l[i];
        } else {
          if (b_i == 0) {
            i++;
            continue;
          }
          buf[b_i] = 0;
          int n = atoi(buf);
          b_i = 0;
          m[y_len][x_i++] = n;
          if (x_i > x_len) {
            x_len = x_i;
          }
        }
        i++;
      }
      if (isdigit(l[i - 1])) {
        buf[b_i] = 0;
        int n = atoi(buf);
        b_i = 0;
        m[y_len][x_i++] = n;
        if (x_i > x_len) {
          x_len = x_i;
        }
      }
      y_len++;
    } else {
      while (l[i] != '\n') {
        if (l[i] == '+' || l[i] == '*') {
          s[s_len++] = l[i];
        }
        i++;
      }
      if (l[i] == '+' || l[i] == '*') {
        s[s_len++] = l[i];
      }
    }
  }
  for (int x = 0; x < x_len; x++) {
    unsigned long part = 0;
    bool first_arg = true;
    for (int y = 0; y < y_len; y++) {
      if (first_arg) {
        part = m[y][x];
        first_arg = false;
        continue;
      }
      switch (s[x]) {
      case '+':
        part += m[y][x];
        break;
      case '*':
        part *= m[y][x];
        break;
      }
    }
    sum += part;
  }
  printf("%ld\n", sum);
  fclose(f);
  if (l) {
    free(l);
  }
}
