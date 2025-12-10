#include <ctype.h>
#include <limits.h>
#include <math.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/param.h>

typedef struct {
  long x, y;
} point;

typedef struct {
  point p1, p2;
  long long area;
} pair;

typedef struct {
  point ps[2048];
  int len;
} circuit;

long long area(point p1, point p2);
int cmp(const void *a, const void *b);

int main() {
  FILE *f;
  if ((f = fopen("input.txt", "r")) == NULL) {
    perror("can't open input.txt");
    exit(1);
  }
  char *l = NULL;
  size_t n;
  static point points[2048];
  static point connecting_points[2048 * 2048];
  size_t p_len = 0;
  size_t cp_len = 0;
  long long y_max = 0, x_max = 0;
  point prev;
  bool has_prev = false;
  while (getline(&l, &n, f) != -1) {
    point p = {0};
    char *x = strtok(l, ",");
    p.x = atol(x);
    char *y = strtok(NULL, "\n");
    p.y = atol(y);
    points[p_len++] = p;
    connecting_points[cp_len++] = p;
    if (has_prev) {
      if (p.x == prev.x) {
        long long max_y = MAX(p.y, prev.y);
        long long min_y = MIN(p.y, prev.y);
        for (long long y = min_y + 1; y < max_y; y++) {
          point c = {p.x, y};
          connecting_points[cp_len++] = c;
        }
      } else {
        long long max_x = MAX(p.x, prev.x);
        long long min_x = MIN(p.x, prev.x);
        for (long long x = min_x + 1; x < max_x; x++) {
          point c = {x, p.y};
          connecting_points[cp_len++] = c;
        }
      }
    }
    has_prev = true;
    prev = p;
    if (y_max < p.y) {
      y_max = p.y;
    }
    if (x_max < p.x) {
      x_max = p.x;
    }
  }
  if (points[0].x == prev.x) {
    long long max_y = MAX(points[0].y, prev.y);
    long long min_y = MIN(points[0].y, prev.y);
    for (long long y = min_y + 1; y < max_y; y++) {
      point c = {points[0].x, y};
      connecting_points[cp_len++] = c;
    }
  } else {
    long long max_x = MAX(points[0].x, prev.x);
    long long min_x = MIN(points[0].x, prev.x);
    for (long long x = min_x + 1; x < max_x; x++) {
      point c = {x, points[0].y};
      connecting_points[cp_len++] = c;
    }
  }

  static pair pairs[2048 * 2048];
  size_t pairs_len = 0;
  for (int i = 0; i < p_len - 1; i++) {
    for (int j = i + 1; j < p_len; j++) {
      pair p = {points[i], points[j], area(points[i], points[j])};
      bool valid = true;
      for (int k = 0; k < cp_len; k++) {
        long long max_x, min_x, max_y, min_y;
        max_x = MAX(points[i].x, points[j].x);
        min_x = MIN(points[i].x, points[j].x);
        max_y = MAX(points[i].y, points[j].y);
        min_y = MIN(points[i].y, points[j].y);
        if (connecting_points[k].x < max_x && connecting_points[k].x > min_x && connecting_points[k].y < max_y &&
            connecting_points[k].y > min_y) {
          valid = false;
          break;
        }
      }
      if (valid) {
        pairs[pairs_len++] = p;
      }
    }
  }

  qsort(pairs, pairs_len, sizeof(pair), cmp);

  // for (int i = 0; i < y_max + 2; i++) {
  //   for (int j = 0; j < x_max + 2; j++) {
  //     bool corner = false;
  //     for (int k = 0; k < p_len; k++) {
  //       if (points[k].y == i && points[k].x == j) {
  //         corner = true;
  //         break;
  //       }
  //     }
  //     if (corner) {
  //       printf("#");
  //     } else {
  //       printf(".");
  //     }
  //   }
  //   printf("\n");
  // }
  // printf("\n");

  for (int i = 0; i < 5; i++) {
    pair p = pairs[i];
    printf("p1: %lu,%lu; p2: %lu,%lu; dist: %llu\n", p.p1.x, p.p1.y, p.p2.x,
           p.p2.y, p.area);
  }

  long long sum = pairs[0].area;
  printf("%llu\n", sum);
  fclose(f);
  if (l) {
    free(l);
  }
}

long long area(point p1, point p2) {
  return (labs(p1.x - p2.x) + 1) * (labs(p1.y - p2.y) + 1);
}

int cmp(const void *a, const void *b) {
  pair *p1 = (pair *)a;
  pair *p2 = (pair *)b;
  return p1->area == p2->area ? 0 : p1->area > p2->area ? -1 : 1;
}
