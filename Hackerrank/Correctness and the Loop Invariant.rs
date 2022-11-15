use std::env;
use std::fs::File;
use std::io::{self, BufRead, Write};

fn user_sort(unsorted: &[i32]) -> Vec<i32> {
    let mut result: Vec<i32> = vec![unsorted[0]];

    for i in 1..unsorted.len() {
        result.push(unsorted[i]);

        if result[i] < result[i - 1] {
            let mut index = i;

            while index != 0
            && result[index] < result[index - 1] {
                let temp = result[index - 1];

                result[index - 1] = result[index];
                result[index]     = temp;                

                index -= 1;
            }
        } else {
            continue;
        }
    }

    result
}

fn main() {
    let stdin = io::stdin();
    let mut stdin_iterator = stdin.lock().lines();

    let mut fptr = File::create(env::var("OUTPUT_PATH")
        .unwrap()).unwrap();

    let _n = stdin_iterator.next().unwrap()
        .unwrap().trim()
        .parse::<i32>().unwrap();

    let arr: Vec<i32> = stdin_iterator.next().unwrap().unwrap()
        .trim_end()
        .split(' ')
        .map(|s| s.to_string().parse::<i32>().unwrap())
        .collect();

    let result = user_sort(&arr);

    for ele in result {
        write!(&mut fptr, "{} ", ele).ok();
    }
}