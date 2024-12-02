#include <ctype.h>
#include <limits.h>
#include <math.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
  uint len;
  uint dir;
  uint color;
} Inst;

typedef struct {
  long y;
  long x;
} Point;

typedef struct {
  long *xs;
  size_t len;
  size_t size;
} Row;

bool find(Row r, long x) {
  for (int i = 0; i < r.len; i++) {
    if (r.xs[i] == x) {
      return true;
    }
  }
  return false;
}

int cmpfunc(const void *a, const void *b) { return (*(long *)a - *(long *)b); }

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

  size_t insts_len = 0;
  size_t insts_size = 8;
  Inst *insts = malloc(sizeof(Inst) * insts_size);

  while (getline(&l, &n, f) != -1) {
    line = l;

    Inst inst;
    char c;
    char buf[7];
    size_t buf_i = 0;

    c = *line++;
    inst.dir = c == 'R' ? 0 : c == 'D' ? 1 : c == 'L' ? 2 : 3;

    line++;
    while (isdigit((c = *line++))) {
      buf[buf_i++] = c;
    }
    buf[buf_i] = 0;
    inst.len = atoi(buf);

    line += 2;
    buf_i = 0;
    while ((c = *line++) != ')') {
      buf[buf_i++] = c;
    }
    inst.dir = buf[buf_i - 1] == '0'   ? 0
               : buf[buf_i - 1] == '1' ? 1
               : buf[buf_i - 1] == '2' ? 2
                                       : 3;
    buf[buf_i - 1] = 0;
    inst.len = strtol(buf, NULL, 16);

    insts[insts_len++] = inst;
    if (insts_len == insts_size) {
      insts_size *= 2;
      insts = realloc(insts, sizeof(Inst) * insts_size);
    }
  }

  size_t p_len = 0;
  size_t p_size = 16;
  Point *p = malloc(sizeof(Point) * p_size);
  Point start = (Point){.x = 0, .y = 0};
  p[p_len++] = start;

  long long circ = 1;

  for (int i = 0; i < insts_len; i++) {
    if (insts[i].dir == 0) {
      p[p_len] =
          (Point){.x = p[p_len - 1].x + insts[i].len, .y = p[p_len - 1].y};
    } else if (insts[i].dir == 1) {
      p[p_len] =
          (Point){.x = p[p_len - 1].x, .y = p[p_len - 1].y + insts[i].len};
    } else if (insts[i].dir == 2) {
      p[p_len] =
          (Point){.x = p[p_len - 1].x - insts[i].len, .y = p[p_len - 1].y};
    } else if (insts[i].dir == 3) {
      p[p_len] =
          (Point){.x = p[p_len - 1].x, .y = p[p_len - 1].y - insts[i].len};
    }
    circ += insts[i].len;
    p_len++;
    if (p_len == p_size) {
      p_size *= 2;
      p = realloc(p, sizeof(Point) * p_size);
    }
  }

  for (int i = 0; i < p_len; i++) {
    printf("x: %ld, y: %ld\n", p[i].x, p[i].y);
    res += p[i].x * p[(i + 1) % p_len].y - p[(i + 1) % p_len].x * p[i].y;
  }

  res += circ;

  res = res % 2 + res / 2;

  printf("\n");
  printf("\n");
  printf("%lld", res);
  printf("\n");
  fclose(f);
  if (l) {
    free(l);
  }
}
