#include <ctype.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char *digit_word_list[] = {"one", "two",   "three", "four", "five",
                           "six", "seven", "eight", "nine"};
int digit_word_list_l = 9;

int find_first_digit_word_before(char *s, char *p);
int find_last_digit_word(char *s);

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
    char *c = l;
    char *last_p = NULL;
    char d1, d2;
    bool first_found = false;
    while (*c != '\0') {
      if (isdigit(*c)) {
        if (!first_found) {
          d1 = *c;
          int word_d = find_first_digit_word_before(l, c);
          if (word_d != -1) {
            d1 = word_d + 0x30;
          }
          first_found = true;
        }
        d2 = *c;
        last_p = c;
      }
      c++;
    }
    int word_d = find_last_digit_word(last_p);
    if (word_d != -1) {
      d2 = word_d + 0x30;
    }
    uint num = (d1 - 0x30) * 10 + d2 - 0x30;
    printf("%d\n", num);
    sum += num;
  }
  printf("%d\n", sum);
  fclose(f);
  if (l) {
    free(l);
  }
}

int find_first_digit_word_before(char *s, char *p) {
  char *first = NULL;
  int d = -1;
  for (int i = 0; i < digit_word_list_l; i++) {
    char *res = strstr(s, digit_word_list[i]);
    if (res != NULL && (first == NULL || first > res)) {
      first = res;
      d = i + 1;
    }
  }
  return first < p ? d : -1;
}

int find_last_digit_word(char *s) {
  char *first = NULL;
  int d = -1;
  for (int i = 0; i < digit_word_list_l; i++) {
    char *res = strstr(s, digit_word_list[i]);
    char *tmpStr = s;
    char *tmpRes = res;
    while (tmpRes != NULL) {
      res = tmpRes;
      tmpStr = tmpStr + strlen(digit_word_list[i]);
      tmpRes = strstr(tmpStr, digit_word_list[i]);
    }
    if (res != NULL && (first == NULL || first < res)) {
      first = res;
      d = i + 1;
    }
  }
  return d;
}
