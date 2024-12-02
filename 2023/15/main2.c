#include <ctype.h>
#include <math.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void trim(char *s) {
  char c;
  while ((c = *s++) != '\0') {
    if (c == '\n' || c == ' ') {
      *(s - 1) = '\0';
      break;
    }
  }
}

typedef struct {
  char *symbol;
  int s_len;
  int focal_len;
} Lens;

typedef struct {
  Lens *lenses;
  int len;
  int size;
} Box;

int hash(char *s, size_t len) {
  int res = 0;
  for (int i = 0; i < len; i++) {
    res += s[i];
    res = res * 17 % 256;
  }
  return res;
}

void remove_lens(Box boxes[256], Lens lens) {
  int box_i = hash(lens.symbol, lens.s_len - 1);
  Box *curr_box = &boxes[box_i];
  for (int i = 0; i < curr_box->len; i++) {
    Lens *curr_lens = &curr_box->lenses[i];
    if (strcmp(curr_lens->symbol, lens.symbol) == 0) {
      free(curr_lens->symbol);
      curr_box->len--;
      while (i < curr_box->len) {
        curr_box->lenses[i] = curr_box->lenses[i + 1];
        i++;
      }
    }
  }
}

void add_lens(Box boxes[256], Lens lens) {
  int box_i = hash(lens.symbol, lens.s_len - 1);
  Box *curr_box = &boxes[box_i];
  for (int i = 0; i < curr_box->len; i++) {
    Lens *curr_lens = &curr_box->lenses[i];
    if (strcmp(curr_lens->symbol, lens.symbol) == 0) {
      curr_lens->focal_len = lens.focal_len;
      return;
    }
  }
  curr_box->len++;
  if (curr_box->len == curr_box->size) {
    curr_box->size *= 2;
    curr_box->lenses = realloc(curr_box->lenses, sizeof(Lens) * curr_box->size);
  }
  curr_box->lenses[curr_box->len - 1] = lens;
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

  Box boxes[256];
  for (int i = 0; i < 256; i++) {
    boxes[i].len = 0;
    boxes[i].size = 16;
    boxes[i].lenses = malloc(sizeof(Lens) * boxes[i].size);
  }

  while (getline(&l, &n, f) != -1) {
    line = l;

    while (*line != 0 && *line != '\n') {
      int len = 0;
      char buf[10];
      while (line[len] != '-' && line[len] != '=') {
        buf[len] = line[len];
        len++;
      }
      buf[len++] = 0;
      Lens lens = (Lens){.s_len = len};
      lens.symbol = malloc(sizeof(char) * len);
      strcpy(lens.symbol, buf);

      if (line[len - 1] == '-') {
        remove_lens(boxes, lens);
      } else {
        lens.focal_len = line[len] - 48;
        len++;
        add_lens(boxes, lens);
      }

      line += len + 1;
    }
  }

  for (int i = 0; i < 256; i++) {
    for (int j = 0; j < boxes[i].len; j++) {
        printf("i: %d, j: %d\n", i, j);
        printf("focal_len: %d\n", boxes[i].lenses[j].focal_len);
      res += (i + 1) * (j + 1) * boxes[i].lenses[j].focal_len;
    }
  }

  printf("\n");
  printf("%lld", res);
  printf("\n");
  fclose(f);
  if (l) {
    free(l);
  }
}
