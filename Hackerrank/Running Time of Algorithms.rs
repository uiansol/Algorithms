use std::env;
use std::fs::File;
use std::io::{self, BufRead, Write};

/*
 * Complete the 'runningTime' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

fn runningTime(arr: &[i32]) -> i32 {
    let mut vector = arr.clone().to_vec();
    let mut result = 0_i32;
    for i in 1..arr.len() {
        let mut index = i;
        while index != 0 && vector[index] < vector[index - 1] {
            let temp = vector[index - 1];
            vector[index - 1] = vector[index];
            vector[index] = temp;
            index -= 1;
            result += 1;
            }
        }        
    result
}

fn main() {
    let stdin = io::stdin();
    let mut stdin_iterator = stdin.lock().lines();

    let mut fptr = File::create(env::var("OUTPUT_PATH").unwrap()).unwrap();

    let n = stdin_iterator.next().unwrap().unwrap().trim().parse::<i32>().unwrap();

    let arr: Vec<i32> = stdin_iterator.next().unwrap().unwrap()
        .trim_end()
        .split(' ')
        .map(|s| s.to_string().parse::<i32>().unwrap())
        .collect();

    let result = runningTime(&arr);

    writeln!(&mut fptr, "{}", result).ok();
}
