#include <limits.h>
#include <math.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main() {
  FILE *f;
  if ((f = fopen("input.txt", "r")) == NULL) {
    perror("can't open input.txt");
    exit(1);
  }
  char *l = NULL;
  size_t n = 0;
  unsigned long sum = 0;
  int pos = 50;
  if (getline(&l, &n, f) == -1) {
    perror("can't get line");
    exit(1);
  }
  char *tok;
  tok = strtok(l, "-");
  while (1) {
    if (tok == NULL) {
      break;
    }
    unsigned long start = atol(tok);
    tok = strtok(NULL, ",");
    if (tok == NULL) {
      break;
    }
    unsigned long end = atol(tok);
    for (; start <= end; start++) {
      unsigned long n = start;
      int num_len = 0;
      while (n > 0) {
        num_len++;
        n /= 10;
      }
      for (int i = 1; i <= num_len / 2; i++) {
        if (num_len % i != 0) {
          continue;
        }
        unsigned long size = (unsigned long)pow(10, i);
        unsigned long currSize = 1;
        unsigned long n = start % size;
        bool match = true;
        while (currSize <= start) {
          unsigned long next = start / currSize ;
          next %= size;
          currSize *= size;
          if (next != n) {
            match = false;
            break;
          }
        }
        if (match) {
          printf("found: %ld\n", start);
          sum += start;
          break;
        }
      }
    }
    tok = strtok(NULL, "-");
  }
  printf("%ld\n", sum);
  fclose(f);
  if (l) {
    free(l);
  }
}
