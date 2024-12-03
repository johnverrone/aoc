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
  //printf("currently checking: %d\n", val);
  int two_digits = 0;
  int decreases = 0;

  int pprev = -1;
  int prev = val % 10;
  val /= 10;
  int curr = val % 10;
  while (val > 0) {
    //printf("\tchecking val: %d and curr: %d and prev: %d and pprev: %d\n", val, curr, prev, pprev);
    if (curr > prev) {
      //printf("\t\tdecreases %d -> %d\n", curr, prev);
      decreases = 1;
    }
    if (curr == prev) {
      //printf("\t\ttwo digits %d and %d\n", curr, prev);
      if (pprev >= 0 && curr != pprev) {
        //printf("\t\tthree digits %d, %d, and %d\n", curr, prev, pprev);
        two_digits = 1;
      }
    }
    val /= 10;
    pprev = prev;
    prev = curr;
    curr = val % 10;
  }
  //printf("two_digits: %d, decreases: %d\n", two_digits, decreases);
  if (two_digits == 1 && decreases == 0) {
    //printf("valid password\n");
    return 1;
  }
  return 0;
}
