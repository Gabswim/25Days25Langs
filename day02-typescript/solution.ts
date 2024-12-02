import { readFileSync } from "node:fs";

const isAllIncreasingOrAllDecreasing = (level: number[]): boolean => {
  let allIncreasing = true;
  let allDecreasing = true;

  for (let i = 1; i < level.length; i++) {
    const diff = level[i] - level[i - 1];
    const absDiff = Math.abs(diff);

    if (absDiff > 3 || diff === 0) return false;
    if (diff > 0) allDecreasing = false;
    if (diff < 0) allIncreasing = false;
  }

  return allIncreasing || allDecreasing;
};

const readFileAndMap = async (fileName: string): Promise<number[][]> => {
  return readFileSync(fileName)
    .toString()
    .split("\n")
    .map((level) => level.split(" ").map(Number)) as number[][];
};

const sol1 = async (fileName: string): Promise<number> => {
  let numberOfSafeReport = 0;

  const levels = await readFileAndMap(fileName);

  for (const level of levels) {
    if (isAllIncreasingOrAllDecreasing(level)) {
      numberOfSafeReport++;
    }
  }

  return numberOfSafeReport;
};

const sol2 = async (fileName: string): Promise<number> => {
  let numberOfSafeReport = 0;
  const levels = await readFileAndMap(fileName);

  for await (const level of levels) {
    if (isAllIncreasingOrAllDecreasing(level)) {
      numberOfSafeReport++;
      continue;
    }

    for (let i = 0; i < level.length; i++) {
      const levelCopy = [...level];
      levelCopy.splice(i, 1);
      if (isAllIncreasingOrAllDecreasing(levelCopy)) {
        numberOfSafeReport++;
        break;
      }
    }
  }

  return numberOfSafeReport;
};



const test = <T>(actual: T, expected: T) => {
  if (actual === expected) console.log("✅ Test pass");
  else console.log(`❌ Test fail, Expected: ${expected}, Actual: ${actual}`);
};

test(await sol1("test-input.txt"), 2);
test(await sol1("input.txt"), 379);

test(await sol2("test-input.txt"), 4);
test(await sol2("input.txt"), 430);
