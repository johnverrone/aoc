import { part1, part2 } from '../utils/print';
import { readFileAsText } from '../utils/readFile';
import { sum } from '../utils/sum';

const contents = await readFileAsText(import.meta.dir);
const rows = contents.split('\n');

type Pull = {
  red: number;
  green: number;
  blue: number;
};

type Game = {
  id: number;
  turns: Pull[];
};

// part 1
const games: Game[] = rows.map((r) => {
  const [info, p] = r.split(':');
  const gameId = Number(info.replace('Game ', ''));
  const pulls: Pull[] = p.split(';').map((t, idx) => {
    const cubes = t.split(',').reduce<Pull>(
      (acc, curr) => {
        const match = curr.match(/(\d+)\s(\w*)/) ?? [];
        const num = Number(match[1]);
        const color = match[2];
        return {
          ...acc,
          [color]: num,
        };
      },
      { red: 0, green: 0, blue: 0 }
    );
    return cubes;
  });
  return { id: +gameId, turns: pulls };
});

const possibleGameIds: number[] = [];
const redCubes = 12;
const greenCubes = 13;
const blueCubes = 14;
games.forEach((g) => {
  let possible = true;
  g.turns.forEach((t) => {
    if (t.red > redCubes) {
      possible = false;
      return;
    }
    if (t.green > greenCubes) {
      possible = false;
      return;
    }
    if (t.blue > blueCubes) {
      possible = false;
      return;
    }
  });
  if (possible) {
    possibleGameIds.push(g.id);
  }
});

part1(sum(possibleGameIds));

// part 2
let powerSum = 0;
games.forEach((g) => {
  let minRed = 0;
  let minGreen = 0;
  let minBlue = 0;
  g.turns.forEach((t) => {
    if (t.red > minRed) {
      minRed = t.red;
    }
    if (t.green > minGreen) {
      minGreen = t.green;
    }
    if (t.blue > minBlue) {
      minBlue = t.blue;
    }
  });
  const power = minRed * minGreen * minBlue;
  powerSum += power;
});

part2(powerSum);
