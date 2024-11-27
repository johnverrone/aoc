import { part1, part2 } from '../utils/print';
import { readFileAsText } from '../utils/readFile';
import { sum } from '../utils/sum';

const contents = await readFileAsText(import.meta.dir);
const rows = contents.split('\n');

// part 1
const numbers: [number, number, number][] = [];
const symbolLocations: [number, number][] = [];
rows.forEach((row, r) => {
  const nums = row.matchAll(/\d+/g);
  for (const n of nums) {
    if (n.index !== undefined) {
      numbers.push([Number(n[0]), r, n.index]);
    }
  }

  const syms = row.matchAll(/[^A-Za-z0-9_.]/g);
  for (const match of syms) {
    if (match.index !== undefined) {
      symbolLocations.push([r, match.index]);
    }
  }
});

const partNumbers: number[] = [];
for (const n of numbers) {
  const [num, x, y] = n;
  for (let r = x - 1; r <= x + 1; r++) {
    for (let c = y - 1; c < y + num.toString().length + 1; c++) {
      if (symbolLocations.some((item) => item[0] === r && item[1] === c)) {
        partNumbers.push(num);
      }
    }
  }
}

part1(sum(partNumbers));

// part 2
const potentialGearLocations: Map<string, number[]> = new Map();
for (const n of numbers) {
  const [num, x, y] = n;
  for (let r = x - 1; r <= x + 1; r++) {
    for (let c = y - 1; c < y + num.toString().length + 1; c++) {
      if (rows[r]?.charAt(c) === '*') {
        const key = `${r}${c}`;
        potentialGearLocations.set(key, [
          ...(potentialGearLocations.get(key) ?? []),
          num,
        ]);
      }
    }
  }
}

const gearRatios: number[] = [];
for (const [_, gears] of potentialGearLocations) {
  // not a gear, doesn't touch two parts
  if (gears.length !== 2) continue;
  gearRatios.push(gears[0] * gears[1]);
}

part2(sum(gearRatios));
