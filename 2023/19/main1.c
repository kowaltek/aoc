#include <ctype.h>
#include <limits.h>
#include <math.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
  int cat;
  int cmp;
  int val;
  char dst[8];
  size_t dst_len;
} Rule;

typedef struct {
  char name[8];
  size_t name_len;
  Rule rules[8];
  size_t rules_len;
} Workflow;

Workflow find_next(char *name, Workflow *ws, size_t w_len) {
  for (int i = 0; i < w_len; i++) {
    if (strcmp(ws[i].name, name) == 0) {
      return ws[i];
    }
  }
  printf("Can not find workflow: %s\n", name);
  exit(1);
}

bool accept(int p[4], Workflow *ws, size_t w_len, char *next) {
  Workflow w = find_next(next, ws, w_len);
  /* printf("next: %s\n", next); */
  int i = 0;
  while (w.rules[i].cat != 4 &&
         (p[w.rules[i].cat] - w.rules[i].val) * w.rules[i].cmp < 0) {
    i++;
  }

  if (strcmp(w.rules[i].dst, "A") == 0) {
    return true;
  }

  if (strcmp(w.rules[i].dst, "R") == 0) {
    return false;
  }

  return accept(p, ws, w_len, w.rules[i].dst);
}

int main() {
  FILE *f;
  if ((f = fopen("input.txt", "r")) == NULL) {
    perror("can't open input.txt");
    exit(1);
  }

  int res = 0;
  char *l = NULL;
  size_t n = 0;
  char *line = NULL;

  size_t w_len = 0;
  size_t w_size = 8;
  Workflow *ws = malloc(sizeof(Workflow) * w_size);

  size_t p_len = 0;
  size_t p_size = 8;
  int *ps = malloc(sizeof(int) * p_size * 4);

  bool in_workflows = true;

  while (getline(&l, &n, f) != -1) {
    line = l;

    if (strlen(line) < 3) {
      in_workflows = false;
      continue;
    }

    if (in_workflows) {
      Workflow w;
      char c;
      char buf[8];
      size_t buf_len = 0;

      while ((c = *line++) != '{') {
        buf[buf_len++] = c;
      }
      buf[buf_len] = 0;
      w.name_len = buf_len;
      strcpy(w.name, buf);

      w.rules_len = 0;
      buf_len = 0;
      while ((c = *line++) != '}' && c != '\n') {
        Rule r;
        if (c == 'A' || c == 'R') {
          r.cat = 4;
          r.dst_len = 1;
          r.dst[0] = c;
          r.dst[1] = 0;
          w.rules[w.rules_len++] = r;
          break;
        }

        if (*line != '<' && *line != '>') {
          r.cat = 4;
          buf_len = 0;
          line--;
          while ((c = *line++) != ',' && c != '}') {
            buf[buf_len++] = c;
          }
          buf[buf_len] = 0;
          r.dst_len = buf_len;
          strcpy(r.dst, buf);
          w.rules[w.rules_len++] = r;
          break;
        }

        r.cat = c == 'x' ? 0 : c == 'm' ? 1 : c == 'a' ? 2 : 3;

        c = *line++;
        r.cmp = c == '<' ? -1 : 1;

        buf_len = 0;
        while ((c = *line++) != ':') {
          buf[buf_len++] = c;
        }
        buf[buf_len] = 0;
        r.val = atoi(buf);

        buf_len = 0;
        while ((c = *line++) != ',' && c != '}') {
          buf[buf_len++] = c;
        }
        buf[buf_len] = 0;
        r.dst_len = buf_len;
        strcpy(r.dst, buf);

        /* printf("rule: cat: %d, cmp: %d, val: %d, dst: %s\n", r.cat, r.cmp, */
        /*        r.val, r.dst); */

        w.rules[w.rules_len++] = r;
      }
      ws[w_len++] = w;
      if (w_len == w_size) {
        w_size *= 2;
        ws = realloc(ws, sizeof(Workflow) * w_size);
      }
    } else {
      char c;
      char buf[8];
      size_t buf_len = 0;
      line++;

      for (int i = 0; i < 4; i++) {
        line += 2;
        buf_len = 0;
        while ((c = *line++) != ',' && c != '}') {
          buf[buf_len++] = c;
        }
        buf[buf_len] = 0;
        ps[p_len * 4 + i] = atoi(buf);
      }

      p_len++;
      if (p_len == p_size) {
        p_size *= 2;
        ps = realloc(ps, sizeof(int) * p_size * 4);
      }
    }
  }

  for (int i = 0; i < w_len; i++) {
    printf("workflow name: %s\n", ws[i].name);
    for (int j = 0; j < ws[i].rules_len; j++) {
      Rule r = ws[i].rules[j];
      printf("rule: cat: %d, cmp: %d, val: %d, dst: %s\n", r.cat, r.cmp, r.val,
             r.dst);
    }
  }

  for (int i = 0; i < p_len; i++) {
    printf("part: x: %d, m: %d, a: %d, s: %d\n", ps[i * 4], ps[i * 4 + 1],
           ps[i * 4 + 2], ps[i * 4 + 3]);
    int p[4] = {ps[i * 4], ps[i * 4 + 1], ps[i * 4 + 2], ps[i * 4 + 3]};
    /* printf("in\n"); */
    if (accept(p, ws, w_len, "in")) {
      printf("accepted\n");
      res += p[0] + p[1] + p[2] + p[3];
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
