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
  size_t x;
  size_t y;
} Pos;

void tilt_north(char **m, size_t y_len, size_t x_len, Pos *obstacles,
                size_t obstacles_len) {
  // from top
  for (int x = 0; x < x_len; x++) {
    size_t last_empty = 0;
    for (int y = 0; y < y_len; y++) {
      if (m[y][x] == '#') {
        break;
      }
      if (m[y][x] == 'O') {
        if (y != last_empty) {
          m[last_empty++][x] = 'O';
          m[y][x] = '.';
        } else {
          last_empty++;
        }
      }
    }
  }

  // from square boulders
  for (int i = 0; i < obstacles_len; i++) {
    Pos obs = obstacles[i];
    size_t last_empty = obs.y + 1;
    for (int y = obs.y + 1; y < y_len; y++) {
      if (m[y][obs.x] == '#') {
        break;
      }
      if (m[y][obs.x] == 'O') {
        if (y != last_empty) {
          m[last_empty++][obs.x] = 'O';
          m[y][obs.x] = '.';
        } else {
          last_empty++;
        }
      }
    }
  }
}

void tilt_east(char **m, size_t y_len, size_t x_len, Pos *obstacles,
               size_t obstacles_len) {
  // from right
  for (int y = 0; y < y_len; y++) {
    size_t last_empty = x_len - 1;
    for (int x = x_len - 1; x >= 0; x--) {
      if (m[y][x] == '#') {
        break;
      }
      if (m[y][x] == 'O') {
        if (x != last_empty) {
          m[y][last_empty--] = 'O';
          m[y][x] = '.';
        } else {
          last_empty--;
        }
      }
    }
  }

  // from square boulders
  for (int i = 0; i < obstacles_len; i++) {
    Pos obs = obstacles[i];
    size_t last_empty = obs.x - 1;
    for (int x = obs.x - 1; x >= 0; x--) {
      if (m[obs.y][x] == '#') {
        break;
      }
      if (m[obs.y][x] == 'O') {
        if (x != last_empty) {
          m[obs.y][last_empty--] = 'O';
          m[obs.y][x] = '.';
        } else {
          last_empty--;
        }
      }
    }
  }
}

void tilt_south(char **m, size_t y_len, size_t x_len, Pos *obstacles,
                size_t obstacles_len) {
  // from down
  for (int x = 0; x < x_len; x++) {
    size_t last_empty = y_len - 1;
    for (int y = y_len - 1; y >= 0; y--) {
      if (m[y][x] == '#') {
        break;
      }
      if (m[y][x] == 'O') {
        if (y != last_empty) {
          m[last_empty--][x] = 'O';
          m[y][x] = '.';
        } else {
          last_empty--;
        }
      }
    }
  }

  // from square boulders
  for (int i = 0; i < obstacles_len; i++) {
    Pos obs = obstacles[i];
    size_t last_empty = obs.y - 1;
    for (int y = obs.y - 1; y >= 0; y--) {
      if (m[y][obs.x] == '#') {
        break;
      }
      if (m[y][obs.x] == 'O') {
        if (y != last_empty) {
          m[last_empty--][obs.x] = 'O';
          m[y][obs.x] = '.';
        } else {
          last_empty--;
        }
      }
    }
  }
}

void tilt_west(char **m, size_t y_len, size_t x_len, Pos *obstacles,
               size_t obstacles_len) {
  // from right
  for (int y = 0; y < y_len; y++) {
    size_t last_empty = 0;
    for (int x = 0; x < x_len; x++) {
      if (m[y][x] == '#') {
        break;
      }
      if (m[y][x] == 'O') {
        if (x != last_empty) {
          m[y][last_empty++] = 'O';
          m[y][x] = '.';
        } else {
          last_empty++;
        }
      }
    }
  }

  // from square boulders
  for (int i = 0; i < obstacles_len; i++) {
    Pos obs = obstacles[i];
    size_t last_empty = obs.x + 1;
    for (int x = obs.x + 1; x < x_len; x++) {
      if (m[obs.y][x] == '#') {
        break;
      }
      if (m[obs.y][x] == 'O') {
        if (x != last_empty) {
          m[obs.y][last_empty++] = 'O';
          m[obs.y][x] = '.';
        } else {
          last_empty++;
        }
      }
    }
  }
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

  size_t m_len = 0;
  size_t m_size = 32;
  char **m = malloc(sizeof(char *) * m_size);
  size_t s_size = 0;

  size_t obstacles_len = 0;
  size_t obstacles_size = 32;
  Pos *obstacles = malloc(sizeof(Pos) * obstacles_size);

  while (getline(&l, &n, f) != -1) {
    line = l;

    if (strchr(line, '.') != NULL || strchr(line, '#') != NULL ||
        strchr(line, 'O')) {
      // add to stored map
      s_size = strlen(line);
      char *s = malloc(sizeof(char) * s_size);
      strcpy(s, line);
      s[--s_size] = 0;
      m[m_len++] = s;
      if (m_len == m_size) {
        m_size *= 2;
        m = realloc(m, sizeof(char *) * m_size);
      }
      for (size_t i = 0; i < strlen(s); i++) {
        if (s[i] == '#') {
          obstacles[obstacles_len++] = (Pos){.x = i, .y = m_len - 1};
          if (obstacles_len == obstacles_size) {
            obstacles_size *= 2;
            obstacles = realloc(obstacles, sizeof(Pos) * obstacles_size);
          }
        }
      }
    }

    res += 0;
  }

  size_t uber_m_size = 1000;
  char ***uber_m = malloc(sizeof(char ***) * uber_m_size);

  for (int i = 0; i < uber_m_size; i++) {
    uber_m[i] = malloc(sizeof(char *) * m_len);
    for (int j = 0; j < m_len; j++) {
      uber_m[i][j] = malloc(sizeof(char) * s_size);
    }
  }

  for (int j = 0; j < m_len; j++) {
    strcpy(uber_m[0][j], m[j]);
  }
  size_t uber_m_len = 1;

  size_t uber_i = 0;
  for (size_t i = 0; i < 1000000000; i++) {
    tilt_north(m, m_len, s_size, obstacles, obstacles_len);

    /* printf("\n"); */
    /* printf("after tilt_north\n"); */
    /* printf("\n"); */
    /* for (int y = 0; y < m_len; y++) { */
    /*   for (int x = 0; x < s_size; x++) { */
    /*     if (m[y][x] == 'O') { */
    /*       res += m_len - y; */
    /*     } */
    /*     printf("%c", m[y][x]); */
    /*   } */
    /*   printf("\n"); */
    /* } */

    tilt_west(m, m_len, s_size, obstacles, obstacles_len);

    /* printf("\n"); */
    /* printf("after tilt_west\n"); */
    /* printf("\n"); */
    /* for (int y = 0; y < m_len; y++) { */
    /*   for (int x = 0; x < s_size; x++) { */
    /*     if (m[y][x] == 'O') { */
    /*       res += m_len - y; */
    /*     } */
    /*     printf("%c", m[y][x]); */
    /*   } */
    /*   printf("\n"); */
    /* } */

    tilt_south(m, m_len, s_size, obstacles, obstacles_len);

    /* printf("\n"); */
    /* printf("after tilt_south\n"); */
    /* printf("\n"); */
    /* for (int y = 0; y < m_len; y++) { */
    /*   for (int x = 0; x < s_size; x++) { */
    /*     if (m[y][x] == 'O') { */
    /*       res += m_len - y; */
    /*     } */
    /*     printf("%c", m[y][x]); */
    /*   } */
    /*   printf("\n"); */
    /* } */

    tilt_east(m, m_len, s_size, obstacles, obstacles_len);

    /* printf("\n"); */
    /* printf("after tilt_east\n"); */
    /* printf("\n"); */
    /* for (int y = 0; y < m_len; y++) { */
    /*   for (int x = 0; x < s_size; x++) { */
    /*     if (m[y][x] == 'O') { */
    /*       res += m_len - y; */
    /*     } */
    /*     printf("%c", m[y][x]); */
    /*   } */
    /*   printf("\n"); */
    /* } */

    for (int j = 0; j < m_len; j++) {
      strcpy(uber_m[uber_m_len][j], m[j]);
    }

    // TODO
    // check for periodic repeating
    long long match = -1;
    for (int k = 0; k < uber_m_len; k++) {
      int matching = 0;
      for (int j = 0; j < m_len; j++) {
        if (strncmp(m[j], uber_m[k][j], s_size) == 0) {
          matching++;
        }
      }
      if (matching == m_len) {
        for (int y = 0; y < m_len; y++) {
          printf("%s    %s\n", uber_m[k][y], m[y]);
        }
        printf("\n");
        match = k;
        break;
      }
    }

    uber_m_len++;
    if (match != -1) {
      printf("match: %lld, i: %zu\n", match, i + 1);
      printf("period: %llu\n", (i + 1) - match);
      uber_i = ((1000000000 - match) % ((i + 1) - match));
      uber_i += match;
      printf("uber_i: %zu\n", uber_i);
      break;
    }
  }

  m = uber_m[uber_i];

  /* for (int k = 0; k < uber_m_len; k++) { */
  /*   printf("index: %d\n", k); */
  /*   for (int y = 0; y < m_len; y++) { */
  /*     for (int x = 0; x < s_size; x++) { */
  /*       printf("%c", uber_m[k][y][x]); */
  /*     } */
  /*     printf("\n"); */
  /*   } */
  /*   int tmp = 0; */
  /*   for (int y = 0; y < m_len; y++) { */
  /*     for (int x = 0; x < s_size; x++) { */
  /*       if (uber_m[k][y][x] == 'O') { */
  /*         tmp += m_len - y; */
  /*       } */
  /*     } */
  /*   } */
  /*   printf("\n"); */
  /*   printf("res: %d\n", tmp); */
  /*   printf("\n"); */
  /* } */

  for (int y = 0; y < m_len; y++) {
    for (int x = 0; x < s_size; x++) {
      if (m[y][x] == 'O') {
        res += m_len - y;
      }
      printf("%c", m[y][x]);
    }
    printf("\n");
  }

  for (int i = 0; i < m_len; i++) {
    free(m[i]);
  }

  m_len = 0;

  printf("\n");
  printf("%lld", res);
  printf("\n");
  fclose(f);
  if (l) {
    free(l);
  }
}
