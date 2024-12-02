#include <ctype.h>
#include <math.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

long long extrapolate_arr(long long *arr, size_t len) {
  long long *arr2 = malloc(sizeof(long long) * len);
  long long prev = arr[0];
  for (size_t i = 1; i < len; i++) {
    arr2[i - 1] = arr[i] - prev;
    prev = arr[i];
  }
  bool is_zero = true;
  for (size_t i = 0; i < len - 1; i++) {
    if (arr2[i] != 0) {
      is_zero = false;
      break;
    }
  }
  return is_zero ? arr[0] : arr[0] - extrapolate_arr(arr2, len - 1);
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

  size_t arr_size = 64;
  long long *arr = malloc(sizeof(long long) * arr_size);

  while (getline(&l, &n, f) != -1) {
    line = l;

    char c;
    char buf[16];
    size_t buf_i = 0;
    size_t arr_len = 0;
    while ((c = *line++) != '\0') {
      if (isdigit(c) || c == '-') {
        buf[buf_i++] = c;
      } else if (buf_i > 0) {
        buf[buf_i] = '\0';
        buf_i = 0;
        arr[arr_len++] = atoll(buf);
        if (arr_len == arr_size) {
          arr_size *= 2;
          arr = realloc(arr, sizeof(arr[0]) * arr_size);
        }
      }
    }

    res += extrapolate_arr(arr, arr_len);
  }

  printf("\n");
  printf("%lld", res);
  printf("\n");
  fclose(f);
  if (l) {
    free(l);
  }
}
