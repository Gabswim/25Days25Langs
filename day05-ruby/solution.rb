def test(actual, expected)
    if actual == expected
        puts "✅ Test pass"
    else
        puts "❌ Test fail, Expected: #{expected}, Actual: #{actual}"
    end
end

def is_update_safe(update, rules)
    
    for i in 0...update.length - 1 do
        val = update[i]
        rule_for_val = rules[val]
        other_values = update[i + 1...update.length]
        other_values.each do |other_value|
            if(rule_for_val.include?(other_value))
                return false
            end
        end
    end
    true
end

def clean_update(update, rules)
    result = update.dup
    
    array_length = result.length
    for first_pos in 0...array_length
        for second_pos in 0...array_length
            current_value = result[first_pos]
            next_value = result[second_pos]
            must_come_before = rules[current_value]
            
            if must_come_before.include?(next_value)
                result[first_pos], result[second_pos] = next_value, current_value
            end
        end
    end
    
    result
 end

def sol_1(fileName)
    file = File.open(fileName)
    lines = file.readlines.map(&:chomp)

    page_ordering_rules = Hash.new { |h, key| h[key] = [] }
    updates = []
    lines.each do |line|
        if line.length == 5
            x = Integer(line[0, 2])
            y = Integer(line[3, 4])
            page_ordering_rules[y] << x
        elsif !line.empty?
            update = line.split(",").map { |num| Integer(num) }
            updates.push(update)
        end
    end

    safe_updates = updates.select { |update| is_update_safe(update, page_ordering_rules) }
    sum = safe_updates.map do |update|
        middle_val = update[update.length / 2]
        middle_val
    end.sum

    puts "sum #{sum}"
    print "safe_updates length  #{safe_updates.length}\n"
    puts "not_safe_updates length #{updates.length - safe_updates.length}"
    #puts "page_ordering_rules #{page_ordering_rules}"
    #puts "safe_updates #{safe_updates}"  

    return sum
end

def sol_2(fileName)
    file = File.open(fileName)
    lines = file.readlines.map(&:chomp)

    page_ordering_rules = Hash.new { |h, key| h[key] = [] }
    page_ordering_rules2 = Hash.new { |h, key| h[key] = [] }
    updates = []
    lines.each do |line|
        if line.length == 5
            x = Integer(line[0, 2])
            y = Integer(line[3, 4])
            page_ordering_rules[y] << x
            page_ordering_rules2[x] << y
        elsif !line.empty?
            update = line.split(",").map { |num| Integer(num) }
            updates.push(update)
        end
    end

    not_safe_updates = updates.select { |update| !is_update_safe(update, page_ordering_rules) }
    puts "not_safe_updates length  #{not_safe_updates.length}"
    clean_updates = not_safe_updates.map do |update|
        clean_update(update, page_ordering_rules2)
    end
    sum = clean_updates.map do |update|
        middle_val = update[update.length / 2]
        middle_val
    end.sum

    #puts "sum #{sum}"
    #puts "page_ordering_rules #{page_ordering_rules}"
    #puts "page_ordering_rules2 #{page_ordering_rules2}"
    #puts "not_safe_updates #{not_safe_updates}"  
    #puts "clean_updates #{clean_updates}"  

    return sum
end

test(sol_1("test-input.txt"), 143)
test(sol_1("input.txt"), 5732)

test(sol_2("test-input.txt"), 123)
test(sol_2("input.txt"), 4716)