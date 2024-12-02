#include <ctype.h>
#include <math.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define LEN 1

// this works for both parts - simply change input and LEN

int main() {
  FILE *f;
  if ((f = fopen("input.txt", "r")) == NULL) {
    perror("can't open input.txt");
    exit(1);
  }
  long long times[LEN];
  size_t times_i = 0;
  long long distances[LEN];
  size_t distances_i = 0;
  long long res = 0;
  char *l = NULL;
  size_t n = 0;
  char *line = NULL;
  while (getline(&l, &n, f) != -1) {
    line = l;
    while (*line++ != ':')
      ;
    if (strstr(l, "Time") != NULL) {
      char c;
      char buf[8];
      size_t buf_i = 0;
      while ((c = *line++) != '\0') {
        if (isdigit(c)) {
          buf[buf_i++] = c;
        } else if (buf_i > 0) {
          buf[buf_i] = '\0';
          times[times_i++] = atoll(buf);
          buf_i = 0;
        }
      }
    }
    if (strstr(l, "Distance") != NULL) {
      char c;
      char buf[8];
      size_t buf_i = 0;
      while ((c = *line++) != '\0') {
        if (isdigit(c)) {
          buf[buf_i++] = c;
        } else if (buf_i > 0) {
          buf[buf_i] = '\0';
          distances[distances_i++] = atoll(buf);
          buf_i = 0;
        }
      }
    }
  }
  for (int i = 0; i < LEN; i++) {
    printf("times[i]: %lld\n", times[i]);
    printf("distances[i]: %lld\n", distances[i]);
    long long times_squared = times[i] * times[i];
    long double delta = sqrt(times_squared + 4 * (-distances[i]));
    long double times_delta = (-times[i]) + delta;
    long long floor_x = floor(times_delta / -2);
    long long ans = times[i] - 2 * floor_x - 1;
    printf("times_squared: %lld\n", times_squared);
    printf("delta: %Lf\n", delta);
    printf("times_delta: %Lf\n", times_delta);
    printf("floor_x: %lld\n", floor_x);
    printf("ans: %lld\n", ans);
    res = res == 0 ? ans : res * ans;
  }
  printf("\n");
  printf("%lld", res);
  printf("\n");
  fclose(f);
  if (l) {
    free(l);
  }
}
