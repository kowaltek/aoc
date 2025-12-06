#include <limits.h>
#include <math.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/param.h>

typedef struct {
  unsigned long min;
  unsigned long max;
} range;

int main() {
  FILE *f;
  if ((f = fopen("input.txt", "r")) == NULL) {
    perror("can't open input.txt");
    exit(1);
  }
  char *l = NULL;
  size_t n;
  range ranges[256];
  int r_len = 0;
  unsigned long sum = 0;
  while (getline(&l, &n, f) != -1) {
    if (l[0] == '\n') {
      break;
    }
    char *min_str = strtok(l, "-");
    unsigned long min = atol(min_str);
    char *max_str = strtok(NULL, "\n");
    unsigned long max = atol(max_str);
    range r = {min, max};
    ranges[r_len] = r;
    r_len++;
  }
  for (int i = 0; i < r_len-1; i++) {
    for (int j = i+1; j < r_len; j++) {
      if ((ranges[i].min <= ranges[j].min && ranges[i].max >= ranges[j].min) ||
          (ranges[i].min <= ranges[j].max && ranges[i].max >= ranges[j].max) ||
          (ranges[i].min >= ranges[j].min && ranges[i].max <= ranges[j].max)) {
        ranges[i].min = MIN(ranges[i].min, ranges[j].min);
        ranges[i].max = MAX(ranges[i].max, ranges[j].max);
        ranges[j] = ranges[r_len - 1];
        j = i+1;
        r_len--;
      }
    }
  }
  for (int i = 0; i < r_len; i++) {
    sum += ranges[i].max - ranges[i].min + 1;
  }
  printf("%ld\n", sum);
  fclose(f);
  if (l) {
    free(l);
  }
}
