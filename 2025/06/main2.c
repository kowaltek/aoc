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
  char m[16][4096];
  int y_len = 0;
  int x_len = 0;
  unsigned long sum = 0;
  while (getline(&l, &n, f) != -1) {
    for (int i = 0; i < n && l[i] != '\n'; i++) {
      m[y_len][i] = l[i];
    }
    x_len = n - 1;
    y_len++;
  }
  unsigned long nums[32];
  int nums_i = 0;
  for (int x = x_len - 1; x >= 0; x--) {
    char num_buf[32];
    int buf_i = 0;
    for (int y = 0; y < y_len - 1; y++) {
      if (m[y][x] != ' ') {
        num_buf[buf_i++] = m[y][x];
      }
    }
    num_buf[buf_i] = 0;
    if (num_buf[0] != 0) {
      unsigned long n = atol(num_buf);
      nums[nums_i++] = n;
      printf("n: %ld\n", n);
    }
    unsigned long part = nums[0];
    switch (m[y_len - 1][x]) {
    case '+':
        printf("adding\n");
      for (int i = 1; i < nums_i ; i++) {
        part += nums[i];
      }
      sum += part;
      printf("res: %ld\n", part);
      nums_i = 0;
      break;
    case '*':
        printf("multiplying\n");
      for (int i = 1; i < nums_i ; i++) {
        part *= nums[i];
      }
      sum += part;
      printf("res: %ld\n", part);
      nums_i = 0;
      break;
    }
  }
  printf("%ld\n", sum);
  fclose(f);
  if (l) {
    free(l);
  }
}
