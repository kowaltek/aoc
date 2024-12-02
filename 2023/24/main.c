#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define LOW_BOUND 200000000000000
#define HIGH_BOUND 400000000000000

#define Max(a, b) ((a) > (b) ? (a) : (b))

#define Min(a, b) ((a) > (b) ? (b) : (a))

#define Abs(a) ((a) < 0 ? -(a) : (a))

typedef struct {
  double x;
  double y;
  double z;
} Point;

typedef struct {
  double x;
  double y;
  double z;
} Velocity;

typedef struct {
  Point p1;
  Point p2;
  Velocity v;
  double a;
  double b;
} Hail;

int main() {
  FILE *f;
  if ((f = fopen("input.txt", "r")) == NULL) {
    perror("can't open input.txt");
    exit(1);
  }

  unsigned long res = 0;
  char *l = NULL;
  char *line = NULL;
  size_t n = 0;

  size_t h_len = 0;
  size_t h_size = 16;
  Hail *h = malloc(sizeof(Hail) * h_size);

  while (getline(&l, &n, f) != -1) {
    line = l;
    char c;
    char buf[32];
    size_t buf_len = 0;

    while ((c = *line++) != ',') {
      buf[buf_len++] = c;
    }
    buf[buf_len] = 0;
    h[h_len].p1.x = atof(buf);

    buf_len = 0;
    while ((c = *line++) != ',') {
      buf[buf_len++] = c;
    }
    buf[buf_len] = 0;
    h[h_len].p1.y = atof(buf);

    buf_len = 0;
    while ((c = *line++) != '@') {
      buf[buf_len++] = c;
    }
    buf[buf_len] = 0;
    h[h_len].p1.z = atof(buf);

    buf_len = 0;
    while ((c = *line++) != ',') {
      buf[buf_len++] = c;
    }
    buf[buf_len] = 0;
    h[h_len].v.x = atof(buf);

    buf_len = 0;
    while ((c = *line++) != ',') {
      buf[buf_len++] = c;
    }
    buf[buf_len] = 0;
    h[h_len].v.y = atof(buf);

    buf_len = 0;
    while ((c = *line++) != '\n' && c != 0) {
      buf[buf_len++] = c;
    }
    buf[buf_len] = 0;
    h[h_len].v.z = atof(buf);

    h[h_len].b = h[h_len].p1.y - h[h_len].p1.x / h[h_len].v.x * h[h_len].v.y;
    h[h_len].a = h[h_len].v.y / h[h_len].v.x;

    if (h[h_len].v.x == 0) {
      h[h_len].p2.x = h[h_len].p1.x;
      h[h_len].p2.y = h[h_len].v.y == 0  ? h[h_len].p1.y
                      : h[h_len].v.y > 0 ? HIGH_BOUND
                                         : LOW_BOUND;
    } else if (h[h_len].v.y == 0) {
      h[h_len].p2.y = h[h_len].p1.y;
      h[h_len].p2.x = h[h_len].v.x == 0  ? h[h_len].p1.x
                      : h[h_len].v.x > 0 ? HIGH_BOUND
                                         : LOW_BOUND;
    } else {
      if (h[h_len].v.x > 0) {
        h[h_len].p2.x = HIGH_BOUND;
      } else {
        h[h_len].p2.x = LOW_BOUND;
      }
      h[h_len].p2.y = h[h_len].p2.x * h[h_len].a + h[h_len].b;
    }

    h_len++;
    if (h_len == h_size) {
      h_size *= 2;
      h = realloc(h, sizeof(Hail) * h_size);
    }
  }
  for (int i = 0; i < 20; i++) {
    printf("Hail:\nPoint 1: x: %f, y: %f, z: %f\nPoint 2: x: %f, y: "
           "%f, z: %f\nVelocity: x: %f, y: %f, z: %f\nA: %f, B: "
           "%f\n\n",
           h[i].p1.x, h[i].p1.y, h[i].p1.z, h[i].p2.x, h[i].p2.y, h[i].p2.z,
           h[i].v.x, h[i].v.y, h[i].v.z, h[i].a, h[i].b);
  }

  for (int i = 0; i < h_len; i++) {
    /* printf("Hail:\nPoint 1: x: %f, y: %f, z: %f\nPoint 2: x: %f, y: " */
    /*        "%f, z: %f\nVelocity: x: %f, y: %f, z: %f\nA: %f, B: " */
    /*        "%f\n\n", */
    /*        h[i].p1.x, h[i].p1.y, h[i].p1.z, h[i].p2.x, h[i].p2.y, h[i].p2.z,
     */
    /*        h[i].v.x, h[i].v.y, h[i].v.z, h[i].a, h[i].b); */

    for (int j = i + 1; j < h_len; j++) {
      if (h[i].a == h[j].a && h[i].b == h[j].b) {
        /* printf("%d, %d\n", i, j); */
        res++;
      } else if (h[i].a != h[j].a) {
        double min_x1 = Min(h[i].p1.x, h[i].p2.x);
        double max_x1 = Max(h[i].p1.x, h[i].p2.x);
        double min_x2 = Min(h[j].p1.x, h[j].p2.x);
        double max_x2 = Max(h[j].p1.x, h[j].p2.x);
        double x = (h[j].b - h[i].b) / (h[i].a - h[j].a);
        double y = h[i].a * x + h[i].b;
        /* printf("x: %f\n", x); */
        if (x >= Max(min_x1, min_x2) && x <= Min(max_x1, max_x2) &&
            x >= LOW_BOUND && x <= HIGH_BOUND && y >= LOW_BOUND &&
            y <= HIGH_BOUND) {
          /* printf("%d, %d\n", i, j); */
          res++;
        }
      }
    }
  }

  printf("\n");
  printf("\n");
  printf("%lu\n", res);
  printf("\n");
  fclose(f);
  if (l) {
    free(l);
  }
}
