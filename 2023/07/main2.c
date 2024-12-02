#include <ctype.h>
#include <math.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define HANDSLEN 1000

typedef struct {
  uint strength;
  char cards[6];
  uint bet;
} hand;

hand parse_hand(char *buf) {
  char tmp[6];
  strcpy(tmp, buf);
  uint run;
  uint longest_run = 1;
  hand h;
  uint no1 = 0;
  for (int i = 0; i < 5; i++) {
    if (tmp[i] == '1') {
      no1++;
    }
  }
  for (int i = 0; i < 5; i++) {
    run = 1;
    for (int j = i + 1; j < 5; j++) {
      if (tmp[i] == tmp[j] && tmp[i] != '1') {
        run++;
      }
    }
    if (run > longest_run) {
      longest_run = run;
    }
  }
  longest_run += no1;

  for (int i = 0; i < 5; i++) {
    for (int j = i + 1; j < 5; j++) {
      if (tmp[i] == tmp[j]) {
        tmp[i] = '\0';
        break;
      }
    }
  }
  uint distinct = 0;
  for (int i = 0; i < 5; i++) {
    if (tmp[i] != '\0' && tmp[i] != '1') {
      distinct++;
    }
  }
  if (distinct == 1 || distinct == 0) {
    h.strength = 6;
  } else if (distinct == 2) {
    if (longest_run == 4) {
      h.strength = 5;
    } else {
      h.strength = 4;
    }
  } else if (distinct == 3) {
    if (longest_run == 3) {
      h.strength = 3;
    } else {
      h.strength = 2;
    }
  } else if (distinct == 4) {
    h.strength = 1;
  } else {
    h.strength = 0;
  }
  strcpy(h.cards, buf);
  return h;
}

void print_hand(hand h) {
  printf("hand strength -> %d; cards -> %s; bet -> %d\n", h.strength, h.cards,
         h.bet);
}

int compare_hands(const void *a, const void *b) {
  hand h1 = *(hand *)a;
  hand h2 = *(hand *)b;
  if (h1.strength != h2.strength) {
    return h1.strength < h2.strength ? -1 : 1;
  }
  for (int i = 0; i < 5; i++) {
    if (h1.cards[i] != h2.cards[i]) {
      if (isdigit(h1.cards[i]) && isdigit(h2.cards[i])) {
        return h1.cards[i] < h2.cards[i] ? -1 : 1;
      }
      if (isdigit(h1.cards[i]) || isdigit(h2.cards[i])) {
        return isdigit(h1.cards[i]) ? -1 : 1;
      }
      return h1.cards[i] > h2.cards[i] ? -1 : 1;
    }
  }
  return 0;
}

hand hands[HANDSLEN];
size_t hands_i = 0;

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
  while (getline(&l, &n, f) != -1) {
    line = l;
    char buf[6];
    size_t buf_i = 0;
    while (*line != ' ') {
      if (*line == 'J') {
        *line = '1';
      }
      buf[buf_i++] = *line++;
    }
    buf[buf_i] = '\0';
    hand h = parse_hand(buf);
    buf_i = 0;
    line++;
    while (*line != '\n' && *line != '\0')
      buf[buf_i++] = *line++;
    ;
    buf[buf_i] = '\0';
    h.bet = atoi(buf);
    hands[hands_i++] = h;
  }
  qsort(hands, HANDSLEN, sizeof(hand), compare_hands);
  for (int i = 0; i < HANDSLEN; i++) {
    print_hand(hands[i]);
    res += (i + 1) * hands[i].bet;
  }
  printf("\n");
  printf("%lld", res);
  printf("\n");
  fclose(f);
  if (l) {
    free(l);
  }
}
