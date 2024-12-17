pub struct Solution;

impl Solution {
    pub fn rotate_the_box(box_grid: Vec<Vec<char>>) -> Vec<Vec<char>> {
        let (n, m) = (box_grid.len(), box_grid[0].len());
        let mut result_grid: Vec<Vec<char>> = vec![vec!['.'; n]; m];

        for (i, row) in box_grid.iter().enumerate() {
            let mut last = m - 1;
            for (j, ch) in row.iter().enumerate().rev() {
                match *ch {
                    '#' => {
                        result_grid[last][n - 1 - i] = '#';
                        last = last.saturating_sub(1);
                    }
                    '*' => {
                        result_grid[j][n - 1 - i] = '*';
                        last = j - 1;
                    }
                    _ => continue,
                }
            }
        }

        result_grid
    }
}

fn print_grid(grid: &Vec<Vec<char>>) {
    for row in grid {
        for &col in row {
            print!("{} ", col);
        }
        println!();
    }
    println!();
}

pub fn test() {
    let grid1 = vec![vec!['#', '.', '#']];
    print_grid(&Solution::rotate_the_box(grid1));

    let grid2 = vec![vec!['#', '.', '*', '.'], vec!['#', '#', '*', '.']];
    print_grid(&Solution::rotate_the_box(grid2));

    let grid3 = vec![
        vec!['#', '#', '*', '.', '*', '.'],
        vec!['#', '#', '#', '*', '.', '.'],
        vec!['#', '#', '#', '.', '#', '.'],
    ];
    print_grid(&Solution::rotate_the_box(grid3));
}
