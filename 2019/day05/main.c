#include <stdio.h>
#include <string.h>
#include <stdlib.h>

const char sample[] = "1,9,10,3,2,3,11,0,99,30,40,50";
const size_t ssize = sizeof(sample);

typedef struct IntCode {
  int code[1000];
  int size;
} IntCode;

IntCode load_program(char *in, size_t size) {
  char p[size];
  strncpy(p, in, size);
  char *pt;
  const char *delim = ",";

  IntCode ic = {};

  pt = strtok(p, delim);
  for (int i = 0; pt != NULL; i++) {
    int n = atoi(pt);
    ic.code[i] = n;
    ic.size++;
    pt = strtok(NULL, delim);
  }

  return ic;
}


int main() {
  IntCode ic = load_program((char *)sample, ssize);
  for (int i=0; i < ic.size; i++) {
    printf("%d\n", ic.code[i]);
  }
  return 0;
}
/* const input = `1,12,2,3,1,1,2,3,1,3,4,3,1,5,0,3,2,6,1,19,2,19,9,23,1,23,5,27,2,6,27,31,1,31,5,35,1,35,5,39,2,39,6,43,2,43,10,47,1,47,6,51,1,51,6,55,2,55,6,59,1,10,59,63,1,5,63,67,2,10,67,71,1,6,71,75,1,5,75,79,1,10,79,83,2,83,10,87,1,87,9,91,1,91,10,95,2,6,95,99,1,5,99,103,1,103,13,107,1,107,10,111,2,9,111,115,1,115,6,119,2,13,119,123,1,123,6,127,1,5,127,131,2,6,131,135,2,6,135,139,1,139,5,143,1,143,10,147,1,147,2,151,1,151,13,0,99,2,0,14,0` */

/* func main() { */
/* 	strCodes := strings.Split(input, ",") */
/* 	var codes = make([]int, len(strCodes)) */
/* 	for i, code := range strCodes { */
/* 		codes[i] = mustInt(code) */
/* 	} */
/* 	var fresh = make([]int, len(codes)) */
/* 	copy(fresh, codes) */

/* 	// part 1 */
/* 	// fmt.Println(runMachine(fresh)) */

/* 	// part 2 */
/* 	for noun := 0; noun < 100; noun++ { */
/* 		for verb := 0; verb < 100; verb++ { */
/* 			copy(fresh, codes) */
/* 			fresh[1] = noun */
/* 			fresh[2] = verb */
/* 			r := runMachine(fresh) */
/* 			// fmt.Println(r) */
/* 			if r == 19690720 { */
/* 				fmt.Println(100*noun + verb) */
/* 				return */
/* 			} */
/* 		} */
/* 	} */
/* } */

/* func mustInt(s string) int { */
/* 	if v, err := strconv.Atoi(s); err != nil { */
/* 		panic(fmt.Sprint(s, err)) */
/* 	} else { */
/* 		return v */
/* 	} */
/* } */

/* func runMachine(codes []int) int { */
/* 	pc := 0 */

/* 	for { */
/* 		op := codes[pc] */
/* 		if op == 99 { */
/* 			return codes[0] */
/* 		} */
/* 		if op == 1 { */
/* 			fmt.Println("add", codes[pc+1], codes[pc+2], codes[pc+3]) */
/* 			codes[codes[pc+3]] = codes[codes[pc+1]] + codes[codes[pc+2]] */
/* 		} */
/* 		if op == 2 { */
/* 			fmt.Println("mul", codes[pc+1], codes[pc+2], codes[pc+3]) */
/* 			codes[codes[pc+3]] = codes[codes[pc+1]] * codes[codes[pc+2]] */
/* 		} */
/* 		pc += 4 */
/* 	} */
/* } */
