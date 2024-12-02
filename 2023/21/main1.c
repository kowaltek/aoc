#include <ctype.h>
#include <limits.h>
#include <math.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define X_SIZE 131
#define Y_SIZE 131
#define STEPS 64

unsigned long long gcd(unsigned long long a, unsigned long long b) {
  while (a != b) {
    if (a > b) {
      a -= b;
    } else {
      b -= a;
    }
  }
  return a;
}

typedef struct {
  int y;
  int x;
  int dist;
} Node;

// returns new size of the nodes list
// expects n_len to already be increased to accomodate new node
size_t add_to_nodes(Node **n, size_t n_len, size_t n_size, Node to_add) {
  if (n_len == n_size) {
    n_size *= 2;
    *n = realloc(*n, sizeof(Node) * n_size);
  }

  (*n)[n_len - 1] = to_add;

  return n_size;
}

int main() {
  FILE *f;
  if ((f = fopen("input.txt", "r")) == NULL) {
    perror("can't open input.txt");
    exit(1);
  }

  unsigned long long res = 0;
  char *l = NULL;
  size_t n = 0;

  char m[Y_SIZE][X_SIZE];
  size_t y = 0;
  Node start = (Node){.y = -1, .x = -1, .dist = 0};

  while (getline(&l, &n, f) != -1) {
    strncpy(m[y], l, X_SIZE);
    for (int x = 0; x < X_SIZE; x++) {
      if (m[y][x] == 'S') {
        start.y = y;
        start.x = x;
      }
    }
    y++;
  }

  for (int i = 0; i < Y_SIZE; i++) {
    for (int j = 0; j < X_SIZE; j++) {
      printf("%c", m[i][j]);
    }
    printf("\n");
  }

  size_t n_len = 0;
  size_t n_size = 8;
  Node *ns = malloc(sizeof(Node) * n_size);
  size_t n_tmp_len = 0;
  size_t n_tmp_size = 8;
  Node *ns_tmp = malloc(sizeof(Node) * n_tmp_size);
  ns[n_len++] = start;
  m[start.y][start.x] = '.';

  while (n_len > 0) {
    printf("n_len: %zu\n", n_len);
    n_tmp_len = 0;
    for (int i = 0; i < n_len; i++) {
      Node curr = ns[i];

      if (curr.dist % 2 == 0) {
        res++;
        m[curr.y][curr.x] = 'x';
      }

      if (curr.dist < STEPS) {
        if (curr.x + 1 < X_SIZE && m[curr.y][curr.x + 1] == '.') {
          n_tmp_size = add_to_nodes(
              &ns_tmp, ++n_tmp_len, n_tmp_size,
              (Node){.y = curr.y, .x = curr.x + 1, .dist = curr.dist + 1});
          m[curr.y][curr.x + 1] = 'v';
        }
        if (curr.x - 1 >= 0 && m[curr.y][curr.x - 1] == '.') {
          n_tmp_size = add_to_nodes(
              &ns_tmp, ++n_tmp_len, n_tmp_size,
              (Node){.y = curr.y, .x = curr.x - 1, .dist = curr.dist + 1});
          m[curr.y][curr.x - 1] = 'v';
        }
        if (curr.y + 1 < Y_SIZE && m[curr.y + 1][curr.x] == '.') {
          n_tmp_size = add_to_nodes(
              &ns_tmp, ++n_tmp_len, n_tmp_size,
              (Node){.y = curr.y + 1, .x = curr.x, .dist = curr.dist + 1});
          m[curr.y + 1][curr.x] = 'v';
        }
        if (curr.y - 1 >= 0 && m[curr.y - 1][curr.x] == '.') {
          n_tmp_size = add_to_nodes(
              &ns_tmp, ++n_tmp_len, n_tmp_size,
              (Node){.y = curr.y - 1, .x = curr.x, .dist = curr.dist + 1});
          m[curr.y - 1][curr.x] = 'v';
        }
      }
    }

    n_len = 0;
    for (int i = 0; i < n_tmp_len; i++) {
      ns[n_len++] = ns_tmp[i];
      if (n_len == n_size) {
        n_size *= 2;
        ns = realloc(ns, sizeof(Node) * n_size);
      }
    }
  }

  for (int i = 0; i < Y_SIZE; i++) {
    for (int j = 0; j < X_SIZE; j++) {
      printf("%c", m[i][j]);
    }
    printf("\n");
  }

  printf("\n");
  printf("\n");
  printf("%lld\n", res);
  printf("\n");
  fclose(f);
  if (l) {
    free(l);
  }
}
