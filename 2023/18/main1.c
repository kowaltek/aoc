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
  int y;
  int x;
} Point;

int main() {
  FILE *f;
  if ((f = fopen("input.txt", "r")) == NULL) {
    perror("can't open input.txt");
    exit(1);
  }

  uint res = 0;
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
    buf[buf_i] = 0;
    inst.color = strtol(buf, NULL, 16);

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
  int min_x = 0;
  int min_y = 0;
  int max_x = 0;
  int max_y = 0;

  for (int i = 0; i < insts_len; i++) {
    printf("inst.dir: %d, inst.len: %d, inst.color: %d\n", insts[i].dir,
           insts[i].len, insts[i].color);
    for (int j = 0; j < insts[i].len; j++) {
      if (insts[i].dir == 0) {
        p[p_len] = (Point){.x = p[p_len - 1].x + 1, .y = p[p_len - 1].y};
      } else if (insts[i].dir == 1) {
        p[p_len] = (Point){.x = p[p_len - 1].x, .y = p[p_len - 1].y + 1};
      } else if (insts[i].dir == 2) {
        p[p_len] = (Point){.x = p[p_len - 1].x - 1, .y = p[p_len - 1].y};
      } else if (insts[i].dir == 3) {
        p[p_len] = (Point){.x = p[p_len - 1].x, .y = p[p_len - 1].y - 1};
      }
      p_len++;
      if (p_len == p_size) {
        p_size *= 2;
        p = realloc(p, sizeof(Point) * p_size);
      }
    }

    min_x = p[p_len - 1].x < min_x ? p[p_len - 1].x : min_x;
    min_y = p[p_len - 1].y < min_y ? p[p_len - 1].y : min_y;
    max_x = p[p_len - 1].x > max_x ? p[p_len - 1].x : max_x;
    max_y = p[p_len - 1].y > max_y ? p[p_len - 1].y : max_y;
  }

  int x_size = max_x - min_x + 1;
  int y_size = max_y - min_y + 1;
  uint *m = calloc(x_size * y_size, sizeof(uint));

  for (int i = 0; i < p_len; i++) {
    m[(p[i].y - min_y) * x_size + p[i].x - min_x] = 1;
  }

  for (int i = 0; i < p_len; i++) {
    printf("x: %d, y: %d\n", p[i].x, p[i].y);
  }

  for (int i = 0; i < y_size * x_size; i++) {
    if (i % x_size == 0) {
      printf("\n");
    }
    printf("%d", m[i]);
  }

  for (int row = x_size; row < y_size * x_size - x_size; row += x_size) {
    for (int col = x_size - 2; col >= 0; col--) {
      if (m[row + col] == 1) {
        continue;
      }
      int i = row + col;
      int intersections = 0;
      bool pending = false;
      bool from_top;
      while (i % x_size != 0) {
        if (m[i] == 1) {
          if (m[i - x_size] == 1 && m[i + x_size] == 1) {
            intersections++;
          } else if (m[i - x_size] == 1) {
            if (pending && from_top) {
              intersections++;
              pending = !pending;
            } else {
              from_top = false;
              pending = !pending;
            }
          } else if (m[i + x_size] == 1) {
            if (pending && !from_top) {
              intersections++;
              pending = !pending;
            } else {
              from_top = true;
              pending = !pending;
            }
          }

          /* if (row % x_size > 0 && m[i - x_size] == 1) { */
          /*   from_top = true; */
          /* } else { */
          /*   from_top = false; */
          /* } */
          /* while (m[i] == 1 && i % x_size > 1) { */
          /*   i--; */
          /* } */
          /* if (row % x_size > 0 && m[i - x_size] == 1 && !from_top) { */
          /*   intersections++; */
          /* } else if (row < x_size * (y_size - 1) && m[i + x_size] == 1 && */
          /*            from_top) { */
          /*   intersections++; */
          /* } */
        }
        i--;
      }

      if (m[i] == 1 && m[i + 1] == 0) {
        intersections++;
      } else if (m[i] == 1 && (m[i - x_size] == 1 || m[i + x_size] == 1)) {
        intersections++;
      }
      if (intersections % 2 == 1) {
        m[row + col] = 2;
      }
    }
  }

  printf("\n");

  for (int i = 0; i < y_size * x_size; i++) {
    if (i % x_size == 0) {
      printf("\n");
    }
    printf("%d", m[i]);
    if (m[i] == 2) {
      m[i] = 1;
    }
    res += m[i];
  }

  printf("\n");
  printf("\n");
  printf("%d", res);
  printf("\n");
  fclose(f);
  if (l) {
    free(l);
  }
}
