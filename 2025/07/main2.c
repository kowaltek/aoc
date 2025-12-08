#include <ctype.h>
#include <limits.h>
#include <math.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
  int x;
  int y;
} point;

unsigned long long traverse(unsigned long long cache[], point pos, point splitters[],
                            int depth, int width);

int main() {
  FILE *f;
  if ((f = fopen("input.txt", "r")) == NULL) {
    perror("can't open input.txt");
    exit(1);
  }
  char *l = NULL;
  size_t n;
  unsigned long long sum = 0;
  point splitters[2048];
  int s_len = 0;
  point start;
  int y = 0;
  int width;
  while (getline(&l, &n, f) != -1) {
    width = n - 1;
    for (int i = 0; i < n && l[i] != '\n'; i++) {
      point p = {i, y};
      switch (l[i]) {
      case 'S':
        start = p;
        break;
      case '^':
        splitters[s_len++] = p;
        break;
      }
    }
    y++;
  }
  unsigned long long cache[y * 256];
  for (int i = 0; i < y * 256; i++) {
    cache[i] = 0;
  }
  sum = traverse(cache, start, splitters, y, 141);
  printf("%llu\n", sum);
  fclose(f);
  if (l) {
    free(l);
  }
}

unsigned long long traverse(unsigned long long cache[], point pos, point splitters[],
                            int depth, int width) {
  if (pos.y == depth) {
    return 1;
  }
  if (cache[pos.y * 256 + pos.x] != 0) {
    return cache[pos.y * 256 + pos.x];
  }
  for (int i = 0; i < 2048; i++) {
    if (splitters[i].x == pos.x && splitters[i].y == pos.y) {
      point p1 = pos;
      point p2 = pos;
      p1.x = p1.x > 0 ? p1.x - 1 : p1.x;
      p1.y++;
      p2.x = p2.x < width - 1 ? p2.x + 1 : p2.x;
      p2.y++;
      unsigned long long tmp = traverse(cache, p1, splitters, depth, width) +
                               traverse(cache, p2, splitters, depth, width);
      cache[pos.y * 256 + pos.x] = tmp;
      return tmp;
    }
  }
  pos.y++;
  unsigned long long tmp = traverse(cache, pos, splitters, depth, width);
  cache[(pos.y-1) * 256 + pos.x] = tmp;
  return tmp;
}
