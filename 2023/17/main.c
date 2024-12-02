#include <ctype.h>
#include <limits.h>
#include <math.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define X_SIZE 141
#define Y_SIZE 141

typedef struct {
  int y;
  int x;
  uint dist;
  uint dir;
  uint dir_len;
} Node;

// returns new size of the nodes list
// expects n_len to already be increased to accomodate new node
size_t add_to_nodes(Node **n, size_t n_len, size_t n_size, Node to_add) {
  if (n_len == n_size) {
    n_size *= 2;
    *n = realloc(*n, sizeof(Node) * n_size);
  }

  bool added = false;
  for (int i = n_len - 2; i >= 0; i--) {
    if ((*n)[i].dist < to_add.dist) {
      (*n)[i + 1] = (*n)[i];
    } else {
      (*n)[i + 1] = to_add;
      added = true;
      break;
    }
  }

  if (!added) {
    (*n)[0] = to_add;
  }

  return n_size;
}

uint traverse(int m[Y_SIZE][X_SIZE]) {
  size_t n_size = 16;
  size_t n_len = 0;
  Node *n = malloc(sizeof(Node) * n_size);
  Node initial_node = {.y = 0, .x = 0, .dist = 0, .dir = 0, .dir_len = 0};
  n[n_len++] = initial_node;
  initial_node = (Node){.y = 0, .x = 0, .dist = 0, .dir = 1, .dir_len = 0};
  n[n_len++] = initial_node;

  size_t v_size = Y_SIZE * X_SIZE * 4 * 10;
  uint v[Y_SIZE * X_SIZE * 4 * 10];
  for (size_t i = 0; i < v_size; i++) {
    v[i] = UINT_MAX;
  }

  while (n_len > 0) {
    if (n_len % 1000 == 0) {
      printf("n_len: %zu\n", n_len);
    }
    Node curr = n[--n_len];

    if (v[curr.y * X_SIZE * 4 * 10 + curr.x * 4 * 10 + curr.dir * 10 +
          curr.dir_len] != UINT_MAX) {
      continue;
    }

    // can turn
    if (curr.dir_len > 2) {
      if (curr.dir == 0 || curr.dir == 2) {
        if (curr.y + 1 < Y_SIZE) {
          n_size =
              add_to_nodes(&n, ++n_len, n_size,
                           (Node){.y = curr.y + 1,
                                  .x = curr.x,
                                  .dist = curr.dist + m[curr.y + 1][curr.x],
                                  .dir = 1,
                                  .dir_len = 0});
        }
        if (curr.y - 1 >= 0) {
          n_size =
              add_to_nodes(&n, ++n_len, n_size,
                           (Node){.y = curr.y - 1,
                                  .x = curr.x,
                                  .dist = curr.dist + m[curr.y - 1][curr.x],
                                  .dir = 3,
                                  .dir_len = 0});
        }
      }

      if (curr.dir == 1 || curr.dir == 3) {
        if (curr.x + 1 < X_SIZE) {
          n_size =
              add_to_nodes(&n, ++n_len, n_size,
                           (Node){.y = curr.y,
                                  .x = curr.x + 1,
                                  .dist = curr.dist + m[curr.y][curr.x + 1],
                                  .dir = 0,
                                  .dir_len = 0});
        }
        if (curr.x - 1 >= 0) {
          n_size =
              add_to_nodes(&n, ++n_len, n_size,
                           (Node){.y = curr.y,
                                  .x = curr.x - 1,
                                  .dist = curr.dist + m[curr.y][curr.x - 1],
                                  .dir = 2,
                                  .dir_len = 0});
        }
      }
    }

    // can go forward
    if (curr.dir_len < 9) {
      int new_x, new_y;
      new_x = curr.x;
      new_y = curr.y;

      if (curr.dir == 0) {
        new_x++;
      } else if (curr.dir == 1) {
        new_y++;
      } else if (curr.dir == 2) {
        new_x--;
      } else if (curr.dir == 3) {
        new_y--;
      }

      if (new_x >= 0 && new_x < X_SIZE && new_y >= 0 && new_y < Y_SIZE) {
        n_size = add_to_nodes(&n, ++n_len, n_size,
                              (Node){.y = new_y,
                                     .x = new_x,
                                     .dist = curr.dist + m[new_y][new_x],
                                     .dir = curr.dir,
                                     .dir_len = curr.dir_len + 1});
      }
    }

    v[curr.y * X_SIZE * 4 * 10 + curr.x * 4 * 10 + curr.dir * 10 +
      curr.dir_len] = curr.dist;
  }

  int res = INT_MAX;
  for (int dir = 0; dir < 4; dir++) {
    for (int dir_len = 0; dir_len < 10; dir_len++) {
      if (v[(Y_SIZE - 1) * X_SIZE * 4 * 10 + (X_SIZE - 1) * 4 * 10 + dir * 10 +
            dir_len] < res) {
        res = v[(Y_SIZE - 1) * X_SIZE * 4 * 10 + (X_SIZE - 1) * 4 * 10 +
                dir * 10 + dir_len];
      }
    }
  }

  return res;
}

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

  size_t y = 0;
  size_t x = 0;
  int m[Y_SIZE][X_SIZE];

  while (getline(&l, &n, f) != -1) {
    line = l;

    char buf[2];
    buf[1] = 0;
    while (isdigit(buf[0] = *line++)) {
      m[y][x++] = atoi(buf);
    }
    x = 0;
    y++;
  }

  res = traverse(m);

  printf("\n");
  for (int y = 0; y < Y_SIZE; y++) {
    for (int x = 0; x < X_SIZE; x++) {
      printf("%d", m[y][x]);
    }
    printf("\n");
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
