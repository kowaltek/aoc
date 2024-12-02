#include <ctype.h>
#include <limits.h>
#include <math.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
  char from[16];
  bool is_high;
} Prev_pulse;

typedef struct module {
  char type;
  char name[16];
  bool is_on;
  size_t next_len;
  char next[16][16];
  size_t prev_len;
  size_t prev_size;
  Prev_pulse *prev;
} Module;

void print_module(Module m) {
  printf("%c%s, is_on: %d, next modules: ", m.type, m.name, m.is_on);
  for (int i = 0; i < m.next_len; i++) {
    printf("%s, ", m.next[i]);
  }
  printf("previous: ");
  for (int i = 0; i < m.prev_len; i++) {
    printf("(%s, was high %d), ", m.prev[i].from, m.prev[i].is_high);
  }
  printf("\n");
}

void add_to_prev(Module *ms, size_t ms_len, char *dest, char *from) {
  for (int i = 0; i < ms_len; i++) {
    if (strcmp(ms[i].name, dest) == 0) {
      ms[i].prev[ms[i].prev_len] = (Prev_pulse){.is_high = false};
      strcpy(ms[i].prev[ms[i].prev_len].from, from);
      ms[i].prev_len++;
      if (ms[i].prev_len == ms[i].prev_size) {
        ms[i].prev_size *= 2;
        ms[i].prev = realloc(ms[i].prev, sizeof(Prev_pulse) * ms[i].prev_size);
      }
      break;
    }
  }
}

typedef struct {
  bool is_high;
  char next[16];
  char from[16];
} Pulse;

void add_pulse(Pulse **ps, size_t *ps_len, size_t *ps_size, bool is_high,
               char *next, char *from) {
  (*ps)[*ps_len] = (Pulse){.is_high = is_high};
  strcpy((*ps)[*ps_len].next, next);
  strcpy((*ps)[*ps_len].from, from);
  (*ps_len)++;
  if (*ps_len == *ps_size) {
    *ps_size *= 2;
    *ps = realloc(*ps, sizeof(Pulse) * *ps_size);
  }
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
  char *line = NULL;

  size_t ms_len = 0;
  size_t ms_size = 8;
  Module *ms = malloc(sizeof(Module) * ms_size);

  while (getline(&l, &n, f) != -1) {
    line = l;

    Module m;
    m.prev_len = 0;
    m.prev_size = 2;
    m.prev = malloc(sizeof(Prev_pulse) * m.prev_size);
    m.next_len = 0;
    if (*line == '%' || *line == '&') {
      m.type = *line++;
      if (m.type == '%') {
        m.is_on = false;
      }
    } else {
      m.type = 'b';
    }

    char buf[10];
    size_t buf_len = 0;
    char c;
    while ((c = *line++) != ' ') {
      buf[buf_len++] = c;
    }
    buf[buf_len] = 0;
    strcpy(m.name, buf);

    while (*line++ != '>')
      ;

    while (*line != '\n' && *line != '\0') {
      line++;
      buf_len = 0;
      while ((c = *line++) != ',' && c != '\n') {
        buf[buf_len++] = c;
      }
      buf[buf_len] = 0;
      strcpy(m.next[m.next_len++], buf);
      while (*line != ' ' && *line != '\n' && *line != '\0') {
        line++;
      }
    }

    ms[ms_len++] = m;
    if (ms_len == ms_size) {
      ms_size *= 2;
      ms = realloc(ms, sizeof(Module) * ms_size);
    }
  }

  for (int i = 0; i < ms_len; i++) {
    for (int j = 0; j < ms[i].next_len; j++) {
      add_to_prev(ms, ms_len, ms[i].next[j], ms[i].name);
    }
  }

  for (int i = 0; i < ms_len; i++) {
    print_module(ms[i]);
  }

  uint low_ps = 0;
  uint high_ps = 0;

  size_t ps_len = 0;
  size_t ps_size = 8;
  Pulse *ps = malloc(sizeof(Pulse) * ps_size);

  size_t ps_tmp_len = 0;
  size_t ps_tmp_size = 8;
  Pulse *ps_tmp = malloc(sizeof(Pulse) * ps_tmp_size);

  int u = 0;

  for (int t = 0; t < 1000; t++) {
    ps[ps_len++] = (Pulse){.is_high = false, .next = "broadcaster"};
    low_ps++;
    printf("----------------------------------\n");
    while (ps_len > 0) {
      /* while (u < 20) { */
      printf("----------------------------------\n");
      ps_tmp_len = 0;
      for (int i = 0; i < ps_len; i++) {
        printf("pulse: is_high: %d, next: %s, from: %s\n", ps[i].is_high,
               ps[i].next, ps[i].from);
        Pulse p = ps[i];
        Module *m = NULL;
        for (int j = 0; j < ms_len; j++) {
          if (strcmp(ms[j].name, p.next) == 0) {
            m = &ms[j];
            break;
          }
        }
        if (m == NULL) {
          continue;
        }
        print_module(*m);
        if (m->type == '%' && !p.is_high) {
          if (m->is_on) {
            m->is_on = false;
            for (int j = 0; j < m->next_len; j++) {
              add_pulse(&ps_tmp, &ps_tmp_len, &ps_tmp_size, false, m->next[j],
                        m->name);
              low_ps++;
            }
          } else {
            m->is_on = true;
            for (int j = 0; j < m->next_len; j++) {
              add_pulse(&ps_tmp, &ps_tmp_len, &ps_tmp_size, true, m->next[j],
                        m->name);
              high_ps++;
            }
          }
        } else if (m->type == '&') {
          bool all_high = true;
          for (int j = 0; j < m->prev_len; j++) {
            if (strcmp(p.from, m->prev[j].from) == 0) {
              m->prev[j].is_high = p.is_high;
            }
            if (!m->prev[j].is_high) {
              all_high = false;
            }
          }
          if (all_high) {
            for (int j = 0; j < m->next_len; j++) {
              add_pulse(&ps_tmp, &ps_tmp_len, &ps_tmp_size, false, m->next[j],
                        m->name);
              low_ps++;
            }
          } else {
            for (int j = 0; j < m->next_len; j++) {
              add_pulse(&ps_tmp, &ps_tmp_len, &ps_tmp_size, true, m->next[j],
                        m->name);
              high_ps++;
            }
          }
        } else if (m->type == 'b') {
          for (int j = 0; j < m->next_len; j++) {
            add_pulse(&ps_tmp, &ps_tmp_len, &ps_tmp_size, false, m->next[j],
                      m->name);
            low_ps++;
          }
        }
      }

      ps_len = 0;
      for (int i = 0; i < ps_tmp_len; i++) {
        /* printf("pulse_tmp: is_high: %d, next: %s, from: %s\n",
         * ps_tmp[i].is_high, */
        /*        ps_tmp[i].next, ps_tmp[i].from); */
        ps[ps_len].is_high = ps_tmp[i].is_high;
        strcpy(ps[ps_len].next, ps_tmp[i].next);
        strcpy(ps[ps_len].from, ps_tmp[i].from);
        ps_len++;
        if (ps_len == ps_size) {
          ps_size *= 2;
          ps = realloc(ps, sizeof(Pulse) * ps_size);
        }
      }
      /* u++; */
    }
  }

  res = low_ps * high_ps;

  printf("low: %d, high: %d\n", low_ps, high_ps);
  printf("\n");
  printf("\n");
  printf("%lld", res);
  printf("\n");
  fclose(f);
  if (l) {
    free(l);
  }
}
