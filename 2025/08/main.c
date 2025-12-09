#include <ctype.h>
#include <limits.h>
#include <math.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
  int x, y, z;
} point;

typedef struct {
  point p1, p2;
  float dist;
} pair;

typedef struct {
  point ps[2048];
  int len;
} circuit;

float dist(point p1, point p2);
int cmp(const void *a, const void *b);
int cmp_c(const void *a, const void *b);
bool cmp_points(point p1, point p2);

int main() {
  FILE *f;
  if ((f = fopen("input.txt", "r")) == NULL) {
    perror("can't open input.txt");
    exit(1);
  }
  char *l = NULL;
  size_t n;
  static point points[2048];
  size_t p_len = 0;
  while (getline(&l, &n, f) != -1) {
    point p = {0};
    char *x = strtok(l, ",");
    p.x = atoi(x);
    char *y = strtok(NULL, ",");
    p.y = atoi(y);
    char *z = strtok(NULL, "\n");
    p.z = atoi(z);
    points[p_len++] = p;
  }
  static pair pairs[2048*2048];
  size_t pairs_len = 0;
  for (int i = 0; i < p_len - 1; i++) {
    for (int j = i + 1; j < p_len; j++) {
      pair p = {points[i], points[j], dist(points[i], points[j])};
      pairs[pairs_len++] = p;
      if (pairs_len == 2048 * 2048) {
        qsort(pairs, pairs_len, sizeof(pair), cmp);
        pairs_len = 2048;
      }
    }
  }

  qsort(pairs, pairs_len, sizeof(pair), cmp);

  for (int i = 0; i < pairs_len; i++) {
    pair p = pairs[i];
    printf("p1: %d,%d,%d; p2: %d,%d,%d; dist: %f\n", p.p1.x, p.p1.y, p.p1.z,
           p.p2.x, p.p2.y, p.p2.z, p.dist);
  }

  pair ans;
  static circuit circuits[2048] = {0};
  int c_len = 0;
  for (int i = 0; i < pairs_len; i++) {
    pair p = pairs[i];
    printf("\np1: %d,%d,%d; p2: %d,%d,%d; dist: %f\n", p.p1.x, p.p1.y,
    p.p1.z,
           p.p2.x, p.p2.y, p.p2.z, p.dist);
    bool match = false;
    for (int j = 0; j < c_len; j++) {
      if (match) {
        break;
      }
      circuit c = circuits[j];
      for (int k = 0; k < c.len; k++) {
        if (cmp_points(p.p1, c.ps[k])) {
          match = true;
          bool other_in = false;
          for (int l = 0; l < c.len; l++) {
            if (cmp_points(p.p2, c.ps[l])) {
              other_in = true;
              break;
            }
          }
          if (!other_in) {
            c.ps[c.len++] = p.p2;
            circuits[j] = c;
          }
          break;
        }
        if (cmp_points(p.p2, c.ps[k])) {
          match = true;
          bool other_in = false;
          for (int l = 0; l < c.len; l++) {
            if (cmp_points(p.p1, c.ps[l])) {
              other_in = true;
              break;
            }
          }
          if (!other_in) {
            c.ps[c.len++] = p.p1;
            circuits[j] = c;
          }
          break;
        }
      }
    }
    if (!match) {
      circuit c = {{p.p1, p.p2}, 2};
      circuits[c_len++] = c;
    } else {
      for (int j = 0; j < c_len - 1; j++) {
        circuit c1 = circuits[j];
        for (int k = j + 1; k < c_len; k++) {
          circuit c2 = circuits[k];
          bool merge = false;
          point duplicate;
          for (int l = 0; l < c1.len; l++) {
            if (merge) {
              break;
            }
            for (int m = 0; m < c2.len; m++) {
              if (cmp_points(c1.ps[l], c2.ps[m])) {
                duplicate = c1.ps[l];
                merge = true;
                break;
              }
            }
          }
          if (merge) {
            for (int l = 0; l < c2.len; l++) {
              if (cmp_points(c2.ps[l], duplicate)) {
                continue;
              }
              c1.ps[c1.len++] = c2.ps[l];
            }
            circuits[k] = circuits[c_len - 1];
            circuits[j] = c1;
            c_len--;
            k = j + 1;
          }
        }
      }
    }
    qsort(circuits, c_len, sizeof(circuit), cmp_c);
    if (circuits[0].len == 1000) {
      printf("\np1: %d,%d,%d; p2: %d,%d,%d; dist: %f\n", p.p1.x, p.p1.y, p.p1.z,
             p.p2.x, p.p2.y, p.p2.z, p.dist);
      ans = pairs[i];
      break;
    }
    for (int i = 0; i < c_len; i++) {
        circuit c = circuits[i];
        printf("circuit len: %d\n", c.len);
        printf("circuit members: ");
        for (int j = 0; j < c.len; j++) {
            printf("[%d,%d,%d],", c.ps[j].x, c.ps[j].y, c.ps[j].z);
        }
        printf("\n");
    }
  }

  printf("\np1: %d,%d,%d; p2: %d,%d,%d; dist: %f\n", ans.p1.x,
         ans.p1.y, ans.p1.z, ans.p2.x, ans.p2.y,
         ans.p2.z, ans.dist);

  qsort(circuits, c_len, sizeof(circuit), cmp_c);

  for (int i = 0; i < c_len; i++) {
    circuit c = circuits[i];
    printf("circuit len: %d\n", c.len);
    printf("circuit members: ");
    for (int j = 0; j < c.len; j++) {
      printf("[%d,%d,%d],", c.ps[j].x, c.ps[j].y, c.ps[j].z);
    }
    printf("\n");
  }

  unsigned long sum = 1;
  for (int i = 0; i < 3; i++) {
    sum *= circuits[i].len;
  }
  printf("%ld\n", sum);
  fclose(f);
  if (l) {
    free(l);
  }
}

float dist(point p1, point p2) {
  return sqrt(pow(p1.x - p2.x, 2) + pow(p1.y - p2.y, 2) + pow(p1.z - p2.z, 2));
}

int cmp(const void *a, const void *b) {
  pair *p1 = (pair *)a;
  pair *p2 = (pair *)b;
  return p1->dist == p2->dist ? 0 : p1->dist > p2->dist ? 1 : -1;
}

int cmp_c(const void *a, const void *b) {
  circuit *c1 = (circuit *)a;
  circuit *c2 = (circuit *)b;
  return c1->len == c2->len ? 0 : c1->len > c2->len ? -1 : 1;
}

bool cmp_points(point p1, point p2) {
  return p1.x == p2.x && p1.y == p2.y && p1.z == p2.z;
}
