#include <stdio.h>
#include <string.h>
#include <stdlib.h>

const char *RANGE = "165432-707912";
const size_t RANGE_SIZE = sizeof(RANGE);

int is_valid(int);

int main() {
  char start[7];
  char end[7];
  strncpy(start, &RANGE[0], 6);
  start[6] = '\0';
  strncpy(end, &RANGE[7], 7);

  int start_int = atoi(start);
  int end_int = atoi(end);

  int count = 0;
  for (int i = start_int; i <= end_int; i++) {
    count += is_valid(i);
  }

  printf("passwords: %d\n", count);
  return 0;
}

int is_valid(int val) {
  int two_digits = -1;
  int decreases = 0;

  int pprev = -1;
  int prev = val % 10;
  val /= 10;
  int curr = val % 10;
  while (val > 0) {
    if (curr > prev) {
      decreases = 1;
      return 0;
    }
    if (curr == prev && two_digits < 0) {
      two_digits = curr;
    }
    if (curr == prev && curr == pprev && curr == two_digits) {
      two_digits = -1;
    }
    val /= 10;
    pprev = prev;
    prev = curr;
    curr = val % 10;
  }
  if (two_digits > 0 && decreases == 0) {
    return 1;
  }
  return 0;
}
