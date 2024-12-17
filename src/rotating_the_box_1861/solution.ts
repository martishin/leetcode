function rotateTheBox(boxGrid: string[][]): string[][] {
  const n = boxGrid.length;
  const m = boxGrid[0].length;

  const resultGrid: string[][] = Array.from({ length: m }, () => Array(n).fill("."));

  for (let i = 0; i < n; i++) {
    let last = m - 1;
    for (let j = m - 1; j >= 0; j--) {
      const ch = boxGrid[i][j];
      switch (ch) {
        case "#":
          resultGrid[last][n - 1 - i] = "#";
          last--;
          break;
        case "*":
          resultGrid[j][n - 1 - i] = "*";
          last = j - 1;
          break;
        default:
          break;
      }
    }
  }

  return resultGrid;
}

function printGrid(grid: string[][]): void {
  for (const row of grid) {
    console.log(row.join(" "));
  }
  console.log();
}

export function test() {
  const grid1: string[][] = [["#", ".", "#"]];
  printGrid(rotateTheBox(grid1));

  const grid2: string[][] = [
    ["#", ".", "*", "."],
    ["#", "#", "*", "."],
  ];
  printGrid(rotateTheBox(grid2));

  const grid3: string[][] = [
    ["#", "#", "*", ".", "*", "."],
    ["#", "#", "#", "*", ".", "."],
    ["#", "#", "#", ".", "#", "."],
  ];
  printGrid(rotateTheBox(grid3));
}
