#include <ctype.h>
#include <math.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define SSIZE 1024

void trim(char *s) {
  char c;
  while ((c = *s++) != '\0') {
    if (c == '\n' || c == ' ') {
      *(s - 1) = '\0';
      break;
    }
  }
}

bool check_valid(char *s, size_t s_len, size_t *groups, size_t groups_len) {
  size_t strike = 0;
  size_t groups_i = 0;
  char prev = '#';
  for (int i = 0; i < s_len; i++) {
    if (s[i] == '#') {
      strike++;
    } else if (prev == '#' && i > 0) {
      if (groups_i == groups_len || strike != groups[groups_i++]) {
        return false;
      }
      strike = 0;
    }
    prev = s[i];
  }
  /* printf("groups_i: %ld\n", groups_i); */
  /* printf("groups_len: %ld\n", groups_len); */
  if (s[s_len - 1] == '#' &&
      (groups_i == groups_len || strike != groups[groups_i++])) {
    return false;
  }
  if (groups_i != groups_len) {
    return false;
  }
  return true;
}

size_t check_combinations(char *start, char *s, size_t s_len, size_t *groups,
                          size_t groups_len, size_t *unknown_indices) {
  bool base = true;
  for (int i = 0; i < s_len; i++) {
    if (unknown_indices[i] == 1) {
      base = false;
      break;
    }
  }

  if (base) {
    /* printf("checking: %s\n", s); */
    bool is_valid = check_valid(s, s_len, groups, groups_len);
    /* if (is_valid) { */
    /*   printf("is_valid: %s\n", s); */
    /* } */
    return is_valid;
  }

  size_t sum = 0;
  for (int i = 0; i < s_len; i++) {
    if (s + i >= start && unknown_indices[i] == 1) {
      unknown_indices[i] = 0;
      s[i] = '#';
      sum += check_combinations(s + i + 1, s, s_len, groups, groups_len,
                                unknown_indices);
      s[i] = '.';
      sum += check_combinations(s + i + 1, s, s_len, groups, groups_len,
                                unknown_indices);
      unknown_indices[i] = 1;
      break;
    }
  }
  return sum;
}

size_t make_cache_key(size_t s_i, size_t s_len, size_t group_i,
                      size_t groups_len, size_t curr_strike) {
  return s_i * s_len * (groups_len + 1) + group_i * s_len + curr_strike;
}

size_t check_combinations2(size_t s_i, char *s, size_t s_len, size_t groups_i,
                           size_t *groups, size_t groups_len,
                           long long int *cache, size_t curr_strike) {
  /* printf("%s\n", s); */
  /* printf("curr_strike: %d, group_i: %d, s_i: %d, groups_len: %d, s_len:
   * %d\n", */
  /* curr_strike, groups_i, s_i, groups_len, s_len); */
  if (groups_i > groups_len) {
    return 0;
  }
  if (s_i == s_len) {
    if (s[s_i - 1] == '#') {
      return curr_strike == groups[groups_i] && groups_i == groups_len - 1;
    }
    return groups_i == groups_len ? 1 : 0;
  }

  size_t cache_key =
      make_cache_key(s_i, s_len, groups_i, groups_len, curr_strike);
  if (cache[cache_key] != -1) {
    /* printf("cache hit for key: %d, value: %d\n", cache_key,
     * cache[cache_key]); */
    return cache[cache_key];
  }

  size_t tmp;
  if (s[s_i] == '.') {
    if (s_i > 0 && s[s_i - 1] == '#' && curr_strike != groups[groups_i]) {
      tmp = 0;
    } else if (s_i > 0 && s[s_i - 1] == '#') {
      tmp = check_combinations2(s_i + 1, s, s_len, groups_i + 1, groups,
                                groups_len, cache, 0);
    } else {
      tmp = check_combinations2(s_i + 1, s, s_len, groups_i, groups, groups_len,
                                cache, 0);
    }
    /* printf("save cache for key: %d, value: %d\n", cache_key, tmp); */
    cache[cache_key] = tmp;
    return tmp;
  }

  if (s[s_i] == '#') {
    tmp = check_combinations2(s_i + 1, s, s_len, groups_i, groups, groups_len,
                              cache, curr_strike + 1);
    /* printf("save cache for key: %d, value: %d\n", cache_key, tmp); */
    cache[cache_key] = tmp;
    return tmp;
  }

  // s[s_i] is '?' at this point
  size_t combined = 0;
  // case '.'
  s[s_i] = '.';
  if (s_i > 0 && s[s_i - 1] == '#' && curr_strike != groups[groups_i]) {
    combined += 0;
  } else if (s_i > 0 && s[s_i - 1] == '#') {
    combined += check_combinations2(s_i + 1, s, s_len, groups_i + 1, groups,
                                    groups_len, cache, 0);
  } else {
    combined += check_combinations2(s_i + 1, s, s_len, groups_i, groups,
                                    groups_len, cache, 0);
  }
  s[s_i] = '?';
  // case '#'
  s[s_i] = '#';
  combined += check_combinations2(s_i + 1, s, s_len, groups_i, groups,
                                  groups_len, cache, curr_strike + 1);
  s[s_i] = '?';
  /* printf("save cache for key: %d, value: %d\n", cache_key, combined); */
  cache[cache_key] = combined;
  return combined;
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

  char *s = malloc(sizeof(char) * SSIZE);
  size_t s_len = 0;

  size_t groups_size = SSIZE;
  size_t groups_len = 0;
  size_t *groups = malloc(sizeof(size_t) * groups_size);

  while (getline(&l, &n, f) != -1) {
    line = l;

    s_len = 0;
    groups_len = 0;
    size_t unknown_indices[SSIZE];

    for (int i = 0; i < SSIZE; i++) {
      unknown_indices[i] = 0;
    }

    while ((s[s_len++] = *line++) != ' ')
      ;

    s[s_len - 1] = '?';

    for (int i = 1; i < 5; i++) {
      for (int j = 0; j < s_len; j++) {
        s[i * s_len + j] = s[j];
      }
    }
    s_len *= 5;
    s[--s_len] = 0;
    printf("%s\n", s);

    size_t buf_i = 0;
    char buf[16];
    char c;

    while ((c = *line++) != '\0') {
      if (isdigit(c)) {
        buf[buf_i++] = c;
      } else if (buf_i > 0) {
        buf[buf_i] = 0;
        buf_i = 0;
        groups[groups_len++] = atoi(buf);
        if (groups_len == groups_size) {
          groups_size *= 2;
          groups = realloc(groups, sizeof(size_t) * groups_size);
        }
      }
    }

    for (int i = 0; i < groups_len; i++) {
      groups[4 * groups_len + i] = groups[3 * groups_len + i] =
          groups[2 * groups_len + i] = groups[groups_len + i] = groups[i];
    }
    groups_len *= 5;

    for (int i = 0; i < groups_len; i++) {
      printf("%ld, ", groups[i]);
    }
    printf("\n");

    size_t cache_len = s_len * s_len * (groups_len + 1);
    long long int *cache = malloc(sizeof(size_t) * cache_len);
    for (int i = 0; i < cache_len; i++) {
      cache[i] = -1;
    }
    /* size_t combinations = */
    /*     check_combinations(s, s, s_len, groups, groups_len, unknown_indices);
     */
    size_t combinations =
        check_combinations2(0, s, s_len, 0, groups, groups_len, cache, 0);
    printf("combinations: %ld\n", combinations);

    res += combinations;
  }

  printf("\n");
  printf("%lld", res);
  printf("\n");
  fclose(f);
  if (l) {
    free(l);
  }
}
