#include <ctype.h>
#include <limits.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

uint scores[209];
uint cards[209];

int main() {
  FILE *f;
  if ((f = fopen("input.txt", "r")) == NULL) {
    perror("can't open input.txt");
    exit(1);
  }
  char *l = NULL;
  char *line = NULL;
  size_t n = 0;
  uint sum = 0;
  int winning[32];
  size_t winning_len = 0;
  while (getline(&l, &n, f) != -1) {
    int score = 0;
    line = l;
    char c;
    while ((c = *line++) != ':')
      ;
    char buf[8];
    size_t buf_len = 0;
    winning_len = 0;
    while ((c = *line++) != '|') {
      if (isdigit(c)) {
        buf[buf_len++] = c;
      } else if (buf_len > 0) {
        buf[buf_len] = '\0';
        winning[winning_len++] = atoi(buf);
        buf_len = 0;
      }
    }
    buf_len = 0;
    while ((c = *line++) != '\0' && c != '\n') {
      if (isdigit(c)) {
        buf[buf_len++] = c;
      } else if (buf_len > 0) {
        buf[buf_len] = '\0';
        int have_num = atoi(buf);
        for (int i = 0; i < winning_len; i++) {
          if (have_num == winning[i]) {
            score = score == 0 ? 1 : score << 1;
            break;
          }
        }
        buf_len = 0;
      }
    }
    // there's no whitespace at the end of line,
    // so loop exits before adding number to array
    // need to do it manually
    buf[buf_len] = '\0';
    int have_num = atoi(buf);
    for (int i = 0; i < winning_len; i++) {
      if (have_num == winning[i]) {
        score = score == 0 ? 1 : score << 1;
        break;
      }
    }
    buf_len = 0;
    for (int i = 0; i < winning_len; i++) {
      printf("%d, ", winning[i]);
    }
    printf("\n");
    printf("score: %d, ", score);
    printf("\n");
    sum += score;
  }
  printf("%d\n", sum);
  fclose(f);
  if (l) {
    free(l);
  }
}
