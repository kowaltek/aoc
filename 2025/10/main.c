#include <ctype.h>
#include <limits.h>
#include <math.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/param.h>

int find_shortest(uint lights, uint buttons[128], int buttons_len);
unsigned make_key(int lights[128], int lights_len, int pressed[128],
                  int pressed_len);
uint hash(int *arr1, size_t len1, int *arr2, size_t len2) {
  uint res = 0;
  for (int i = 0; i < len1; i++) {
    res += arr1[i];
    res = res * 117 % UINT_MAX;
  }
  for (int i = 0; i < len2; i++) {
    res += arr2[i];
    res = res * 177 % UINT_MAX;
  }
  return res;
}

int main() {
  FILE *f;
  if ((f = fopen("input.txt", "r")) == NULL) {
    perror("can't open input.txt");
    exit(1);
  }
  int sum = 0;
  char *l = NULL;
  size_t n;

  while (getline(&l, &n, f) != -1) {
    char *saveptr_main = NULL;
    char *saveptr_buttons = NULL;
    char *saveptr_switch = NULL;
    l++;
    uint lights = 0;
    char *lights_str = strtok_r(l, "]", &saveptr_main);

    for (int i = 0; i < strlen(lights_str); i++) {
      if (lights_str[i] == '#') {
        lights += 1 << i;
      }
    }

    uint buttons[128] = {0};
    int buttons_len = 0;
    char *buttons_str = strtok_r(NULL, "{", &saveptr_main);

    for (char *button_str = strtok_r(buttons_str, " ", &saveptr_buttons);
         button_str != NULL;
         button_str = strtok_r(NULL, " ", &saveptr_buttons)) {
      button_str++;
      button_str[strlen(button_str) - 1] = 0;
      for (char *switch_str = strtok_r(button_str, ",", &saveptr_switch);
           switch_str != NULL;
           switch_str = strtok_r(NULL, ",", &saveptr_switch)) {
        n = atoi(switch_str);
        buttons[buttons_len] += 1 << n;
      }
      buttons_len++;
    }

    static int cache[UINT_MAX] = {0};

    int part = buttons_len;
    uint combinations = pow(2, buttons_len);
    for (uint i = 0; i < combinations; i++) {
      int cr = 0;
      int cl = lights;
      for (int j = 0; j < buttons_len; j++) {
        if ((i >> j)%2 == 1) {
          cl ^= buttons[j];
          cr++;
        }
      }
      if (cl == 0 && cr < part) {
        part = cr;
      }
      if (part ==1) {
          break;
      }
    }

    // int part = find_shortest(lights, buttons, buttons_len);
    //
    printf("%d\n", part);
    sum += part;

    l = NULL;
  }
  printf("%d\n", sum);
  fclose(f);
}

int find_shortest(uint lights, uint buttons[128], int buttons_len) {
  if (lights == 0) {
    return 0;
  }

  int res = -1;
  for (int i = 0; i < buttons_len; i++) {
    lights ^= buttons[i];

    uint pressed = buttons[i];
    buttons[i] = buttons[buttons_len - 1];
    buttons_len--;

    int part = find_shortest(lights, buttons, buttons_len) + 1;

    buttons_len++;
    buttons[buttons_len - 1] = buttons[i];
    buttons[i] = pressed;

    lights ^= buttons[i];

    if (part < res || res == -1) {
      res = part;
    }
    if (res == 1) {
      break;
    }
  }
  return res;
}

unsigned make_key(int lights[128], int lights_len, int pressed[128],
                  int pressed_len) {
  unsigned key = 0;
  for (int i = 0; i < lights_len; i++) {
    key |= lights[i] << i;
  }
  int root = 1;
  return key;
}
