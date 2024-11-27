import { part1, part2 } from '../utils/print';
import { readFileAsText } from '../utils/readFile';
import { sum } from '../utils/sum';

const contents = await readFileAsText(import.meta.dir);
const rows = contents.split('\n');

// part 1
const nums = rows.map((r) => {
  const first = r.replace(/[a-z]/g, '').at(0);
  const last = r.replace(/[a-z]/g, '').at(-1);
  return Number(`${first}${last}`);
});

const s = sum(nums);

part1(s);

// part 2
const stringToNum: { [key: string]: string } = {
  one: '1',
  two: '2',
  three: '3',
  four: '4',
  five: '5',
  six: '6',
  seven: '7',
  eight: '8',
  nine: '9',
};

const nums2 = rows.map((r) => {
  const first = r
    .replace(/one|two|three|four|five|six|seven|eight|nine/g, (a) => {
      return stringToNum[a] ?? a;
    })
    .replace(/[a-z]/g, '')
    .at(0);

  const last = r
    .split('')
    .reverse()
    .join('')
    .replace(/eno|owt|eerht|ruof|evif|xis|neves|thgie|enin/g, (a) => {
      return stringToNum[a.split('').reverse().join('')] ?? a;
    })
    .split('')
    .reverse()
    .join('')
    .replace(/[a-z]/g, '')
    .at(-1);
  return Number(`${first}${last}`);
});

const s2 = sum(nums2);

part2(s2);
