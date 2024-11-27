# Advent of Code 2023

Using [bun](https://bun.sh/) as a TypeScript runtime this year.

To run a puzzle, ensure `bun` is installed and run `bun dayX/index.ts`

### Day 01

Started a little tougher on part 2 than I was hoping for 😂. I really wanted to use regex for this (it felt right), but it got a little tricky when the written out versions of the numbers could overlap. Eg. `eightwo` needs to be `8` when searching from the front and `2` when searching from the back. I ended up doing a lot of ugly string reversing but we got the answer and I'm going to leave the code as-is. Waking up early and going straight to the computer. Feels like the holidays.

### Day 02

A very verbose approach but it was fun to create an appropriate data structure for a `Game`. This made part 2 pretty easy as well. Didn't solve either part in O(n) but whatevs. Was nice to get both answers right on the first try. The power of organizing the data right. Also learned a little more about how JavaScript regex matching works.

### Day 03

Did not do this one on actual day 3, but it was fun! Tested a bit of regex skills, tried to be smart about runtime complexity while parsing 2d arrays. React has tended to make me program in a way more functional way (no variables, use functional array methods like `.map`) but it was refreshing to go back to using some `for` loops and pushing to a solution array.
