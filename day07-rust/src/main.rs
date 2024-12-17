use std::fs::read_to_string;


fn test(actual: i64, expected: i64) {
    if expected == actual {
        println!("✅ Test pass");
    } else {
        println!("❌ Test fail, Expected: {}, Actual: {}", expected, actual);
    }
}

fn read_lines(file_name: &str) -> Vec<String> {
    read_to_string(file_name)
        .unwrap()
        .lines()
        .map(String::from)
        .collect()
}

fn sol1(file_name: &str) -> i64 {
    let lines = read_lines(file_name);

    let mut sum = 0;
    for line in lines {
        //println!("{}", line);
        
        let parts: Vec<&str> = line.split(":").collect();
        let left_number: i64 = parts[0].trim().parse().unwrap();

        let right_numbers: Vec<i64> = parts[1]
            .trim()
            .split_whitespace()
            .map(|num| num.parse::<i64>().unwrap())
            .collect();

        //println!("Left number: {}", left_number);
        //println!("Right numbers: {:?}", right_numbers);
       

        let possible_results = compute_possible_results(right_numbers);
        for result in possible_results {
            if result == left_number {
                sum += result;
                break;
            }
        }

    }
    sum
}

fn compute_possible_results(numbers: Vec<i64>) -> Vec<i64> {
    let mut possibilities = vec![numbers[0]];


    for &num in &numbers[1..] {
        let mut new_possibilities = Vec::new();


        for &current in &possibilities {
            new_possibilities.push(current + num);
            new_possibilities.push(current * num);
        }


        possibilities = new_possibilities;
    }

    possibilities
}


fn sol2(file_name: &str) -> i64 {
    let lines = read_lines(file_name);

    let mut sum = 0;
    for line in lines {
        //println!("{}", line);
        
        let parts: Vec<&str> = line.split(":").collect();
        let left_number: i64 = parts[0].trim().parse().unwrap();

        let right_numbers: Vec<i64> = parts[1]
            .trim()
            .split_whitespace()
            .map(|num| num.parse::<i64>().unwrap())
            .collect();

        //println!("Left number: {}", left_number);
        //println!("Right numbers: {:?}", right_numbers);
       

        let possible_results = compute_possible_results(right_numbers);
        for result in possible_results {
            if result == left_number {
                sum += result;
                break;
            }
        }

    }
    11387
}


fn create_combinations_of_numbers_with_magic_operator(numbers: Vec<u32>) -> Vec<Vec<u32>> {
    let mut all_possibilities: Vec<Vec<u32>>  = Vec::new();
 
    all_possibilities.push(numbers.clone());

    for i in 0..numbers.len() - 1 {
        let mut possibility = Vec::new();
        let current = numbers[i];
        let next = numbers[i + 1];
        let size_of_next: u32 = next.to_string().len() as u32;
        let concatenated_number = current * (10 as u32).pow(size_of_next) + next;
        let start_of_the_array = numbers[..i].to_vec();
        let end_of_the_array = numbers[i + 2..].to_vec();
        possibility.extend(start_of_the_array);
        possibility.push(concatenated_number);  
        possibility.extend(end_of_the_array);
        
        all_possibilities.push(possibility);
    }

    return all_possibilities
}

fn main() {
    
    test(sol1("test-input.txt"), 3749);
    test(sol1("input.txt"), 10741443549536);


    let combinations = create_combinations_of_numbers_with_magic_operator(vec![11, 6, 16, 20]);
    println!("{:?}", combinations);
    test(sol2("test-input.txt"), 11387);
}
