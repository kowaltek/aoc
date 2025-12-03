#include <ctype.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
  uint red;
  uint green;
  uint blue;
} grab;

grab max_grab = {12, 13, 14};

uint get_id(char *l);
bool gb_is_greater(grab g1, grab g2);
bool grabs_possible(char *l);
grab get_grab(char *start, char *end);

int main() {
  FILE *f;
  if ((f = fopen("input.txt", "r")) == NULL) {
    perror("can't open input.txt");
    exit(1);
  }
  char *l = NULL;
  size_t n = 0;
  uint sum = 0;
  while (getline(&l, &n, f) != -1) {
    uint id = get_id(l);
    if (grabs_possible(l)) {
      sum += id;
    }
  }
  printf("%d\n", sum);
  fclose(f);
  if (l) {
    free(l);
  }
}

uint get_id(char *l) {
  char buf[4];
  char *d = buf;
  // move past 'Game ' and get the game id
  l += 4;
  while (*++l != ':') {
    *d++ = *l;
  }
  *d = '\0';
  return atoi(buf);
}

bool gb_is_possible(grab g) {
  if (g.red > max_grab.red || g.green > max_grab.green ||
      g.blue > max_grab.blue) {
    return false;
  }
  return true;
}

bool grabs_possible(char *l) {
  while (*l++ != ':')
    ;
  while (*++l != '\0' && *l != '\n') {
    if (*l != ' ') {
      char *start = l;
      while (*++l != ';' && *l != '\0' && *l != '\n')
        ;
      char *end = l;
      grab g = get_grab(start, end);
      if (!gb_is_possible(g)) {
        return false;
      }
    }
  }
  return true;
}

grab get_grab(char *start, char *end) {
  grab g = {0, 0, 0};
  while (start < end) {
    // at this point we are sure *start points to first digit
    // get the num
    char num_buf[4];
    char *c = num_buf;
    *c++ = *start;
    while (*++start != ' ') {
      *c++ = *start;
    }
    *c = '\0';
    uint num = atoi(num_buf);
    // get the word
    char char_buf[6];
    c = char_buf;
    while (*++start != ',' && *start != ';' && *start != '\0' &&
           *start != '\n') {
      *c++ = *start;
    }
    start++;
    *c = '\0';
    if (strcmp(char_buf, "red") == 0) {
      g.red = num;
    } else if (strcmp(char_buf, "green") == 0) {
      g.green = num;
    } else if (strcmp(char_buf, "blue") == 0) {
      g.blue = num;
    }
  }
  return g;
}
