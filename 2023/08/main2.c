#include <ctype.h>
#include <math.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define NODESLEN 26 * 26 * 26

typedef struct Node {
  int left;
  int right;
} node;

node node_list[NODESLEN];

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

  size_t instructions_size = 64;
  uint *instructions = malloc(sizeof(uint) * instructions_size);
  size_t instructions_len = 0;

  size_t starting_points_size = 64;
  size_t *starting_points = malloc(sizeof(size_t) * instructions_size);
  size_t starting_points_len = 0;

  while (getline(&l, &n, f) != -1) {
    line = l;
    printf("%s", line);

    if (strchr(line, '=') == NULL && n > 5) {
      char c;
      while ((c = *line++) == 'R' || c == 'L') {
        instructions[instructions_len++] = c == 'R' ? 1 : 0;
        if (instructions_len == instructions_size) {
          instructions_size *= 2;
          instructions =
              realloc(instructions, sizeof(uint) * instructions_size);
        }
      }
    }

    if (strchr(line, '=') != NULL && n > 2) {
      char c;
      int index = 0;
      uint offset = 2;
      while ((c = *line++) != ' ') {
        index += (c - 65) * (offset == 2 ? 26 * 26 : offset == 1 ? 26 : 1);
        if (offset == 0) {
          if (c == 'A') {
            starting_points[starting_points_len++] = index;
            if (starting_points_len == starting_points_size) {
              starting_points_size *= 2;
              starting_points =
                  realloc(starting_points, sizeof(uint) * starting_points_size);
            }
          }
        }
        offset--;
      }
      while (*line++ != '(')
        ;
      int left = 0;
      offset = 2;
      while ((c = *line++) != ',') {
        left += (c - 65) * (offset == 2 ? 26 * 26 : offset == 1 ? 26 : 1);
        offset--;
      }
      while (*line++ != ' ')
        ;
      int right = 0;
      offset = 2;
      while ((c = *line++) != ')') {
        right += (c - 65) * (offset == 2 ? 26 * 26 : offset == 1 ? 26 : 1);
        offset--;
      }
      node_list[index] = (node){left, right};
    }
  }

  size_t instruction_i = 0;
  unsigned long long res_list[6];
  for (int i = 0; i < 6; i++) {
    res_list[i] = 0;
  }
  while (true) {
    for (int i = 0; i < starting_points_len; i++) {
      if (starting_points[i] % 26 == 25) {
        res_list[i] = res_list[i] != 0 ? res_list[i] : res;
      }
      starting_points[i] = instructions[instruction_i] == 1
                               ? node_list[starting_points[i]].right
                               : node_list[starting_points[i]].left;
    }
    res++;
    int ended = 0;
    for (int i = 0; i < 6; i++) {
      ended = res_list[i] != 0 ? ended + 1 : ended;
    }
    if (ended == starting_points_len) {
      break;
    }
    instruction_i = (instruction_i + 1) % instructions_len;
  }
  for (int i = 0; i < 5; i++) {
    printf("%lld, %lld\n", res_list[i], res_list[i + 1]);
    int gcd_i = gcd(res_list[i], res_list[i + 1]);
    printf("gcd: %d\n", gcd_i);
    res_list[i + 1] = res_list[i + 1] * (res_list[i] / gcd_i);
  }
  printf("\n");
  printf("%lld", res_list[5]);
  printf("\n");
  fclose(f);
  if (l) {
    free(l);
  }
}
