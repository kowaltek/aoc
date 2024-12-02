#include <ctype.h>
#include <math.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
  int dir;
  int y;
  int x;
  bool to_remove;
} Beam;

typedef struct {
  int dirs[4];
  int y;
  int x;
} Tile;

void trim(char *s) {
  char c;
  while ((c = *s++) != '\0') {
    if (c == '\n' || c == ' ') {
      *(s - 1) = '\0';
      break;
    }
  }
}

// returns new beams array length
size_t remove_beams(Beam *beams, size_t beams_len) {
  for (int i = 0; i < beams_len - 1; i++) {
    if (beams[i].to_remove) {
      for (int j = i; j < beams_len - 1; j++) {
        beams[j] = beams[j + 1];
      }
      beams_len--;
    }
  }
  if (beams[beams_len - 1].to_remove) {
    beams_len--;
  }
  return beams_len;
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

  // visited points - need to record the direction from which were visited
  size_t visited_tiles_len = 0;
  size_t visited_tiles_size = 16;
  Tile *visited_tiles;
  visited_tiles = malloc(sizeof(Tile) * visited_tiles_size);

  size_t m_len = 0;
  size_t m_size = 16;
  char **m;
  m = malloc(sizeof(char *) * m_size);

  size_t row_size;

  while (getline(&l, &n, f) != -1) {
    line = l;

    row_size = strlen(line);
    char *row = malloc(sizeof(char) * (row_size + 1));
    strcpy(row, line);
    trim(row);
    m[m_len++] = row;
    if (m_len == m_size) {
      m_size *= 2;
      m = realloc(m, sizeof(char *) * m_size);
    }
    row_size = strlen(row);
  }

  for (int y = 0; y < m_len; y++) {
    for (int x = 0; x < row_size; x++) {
      printf("%c", m[y][x]);
    }
    printf("\n");
  }
  printf("\n");

  size_t beams_len = 1;
  size_t beams_size = 16;
  Beam *beams = malloc(sizeof(Beam) * beams_size);
  beams[0] = (Beam){.dir = 0, .y = 0, .x = -1, .to_remove = false};

  while (beams_len > 0) {
    int initial_beams_len = beams_len;
    printf("\n");
    for (int i = 0; i < initial_beams_len; i++) {
      printf("beam[%d].x: %d, beam[%d].y: %d, beam[%d].dir: %d\n", i,
             beams[i].x, i, beams[i].y, i, beams[i].dir);
      int dir = beams[i].dir;
      if (dir == 0) {
        beams[i].x++;
      } else if (dir == 1) {
        beams[i].y++;
      } else if (dir == 2) {
        beams[i].x--;
      } else if (dir == 3) {
        beams[i].y--;
      }

      for (int visited_tiles_i = 0; visited_tiles_i < visited_tiles_len;
           visited_tiles_i++) {
        Tile t = visited_tiles[visited_tiles_i];
        if (beams[i].x == t.x && beams[i].y == t.y &&
            t.dirs[beams[i].dir] > 0) {
          beams[i].to_remove = true;
          break;
        }
      }

      bool already_in_visited = false;
      for (int visited_tiles_i = 0; visited_tiles_i < visited_tiles_len;
           visited_tiles_i++) {
        if (visited_tiles[visited_tiles_i].x == beams[i].x &&
            visited_tiles[visited_tiles_i].y == beams[i].y) {
          already_in_visited = true;
          visited_tiles[visited_tiles_i].dirs[beams[i].dir] = 1;
        }
      }

      if (!already_in_visited) {
        visited_tiles[visited_tiles_len].x = beams[i].x;
        visited_tiles[visited_tiles_len].y = beams[i].y;
        for (int dir = 0; dir < 4; dir++) {
          visited_tiles[visited_tiles_len].dirs[dir] = -1;
        }
        visited_tiles[visited_tiles_len++].dirs[beams[i].dir] = 1;
        if (visited_tiles_len == visited_tiles_size) {
          visited_tiles_size *= 2;
          visited_tiles =
              realloc(visited_tiles, sizeof(Tile) * visited_tiles_size);
        }
      }

      if (beams[i].y < 0 || beams[i].y >= m_len || beams[i].x < 0 ||
          beams[i].x >= row_size) {
        beams[i].to_remove = true;
      } else {
        char curr_c = m[beams[i].y][beams[i].x];
        if (curr_c == '\\') {
          if (dir == 0) {
            beams[i].dir = 1;
          } else if (dir == 1) {
            beams[i].dir = 0;
          } else if (dir == 2) {
            beams[i].dir = 3;
          } else if (dir == 3) {
            beams[i].dir = 2;
          }
        } else if (curr_c == '/') {
          if (dir == 0) {
            beams[i].dir = 3;
          } else if (dir == 1) {
            beams[i].dir = 2;
          } else if (dir == 2) {
            beams[i].dir = 1;
          } else if (dir == 3) {
            beams[i].dir = 0;
          }
        } else if (curr_c == '-' && (dir == 1 || dir == 3)) {
          beams[i].dir = 0;
          beams[beams_len].x = beams[i].x;
          beams[beams_len].y = beams[i].y;
          beams[beams_len].to_remove = false;
          beams[beams_len++].dir = 2;
          if (beams_len == beams_size) {
            beams_size *= 2;
            beams = realloc(beams, sizeof(Beam) * beams_size);
          }
        } else if (curr_c == '|' && (dir == 0 || dir == 2)) {
          beams[i].dir = 1;
          beams[beams_len].x = beams[i].x;
          beams[beams_len].y = beams[i].y;
          beams[beams_len].to_remove = false;
          beams[beams_len++].dir = 3;
          if (beams_len == beams_size) {
            beams_size *= 2;
            beams = realloc(beams, sizeof(Beam) * beams_size);
          }
        }
      }
    }

    beams_len = remove_beams(beams, beams_len);
    /* printf("beams_len: %zu\n", beams_len); */
  }

  for (int i = 0; i < visited_tiles_len; i++) {
    if (visited_tiles[i].x < row_size && visited_tiles[i].x >= 0 &&
        visited_tiles[i].y < m_len && visited_tiles[i].y >= 0) {
      res++;
    }
  }

  for (int y = 0; y < m_len; y++) {
    for (int x = 0; x < row_size; x++) {
      bool visited = false;
      for (int t = 0; t < visited_tiles_len; t++) {
        if (visited_tiles[t].x == x && visited_tiles[t].y == y) {
          visited = true;
          break;
        }
      }
      if (visited && m[y][x] == '.') {
        printf("#");
      } else {
        printf("%c", m[y][x]);
      }
    }
    printf("\n");
  }

  printf("\n");
  printf("\n");
  printf("%lld", res);
  printf("\n");
  fclose(f);
  if (l) {
    free(l);
  }
}
