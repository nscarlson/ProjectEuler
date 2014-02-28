contents = File.read('largest_product.txt')

i = 0
largest = "0";
largest_product = 1;

until i > contents.size - 5
	substring = contents[i..i+4].chars.sort.join

	if substring.to_i > largest.to_i
		largest = substring
		puts "New largest is: #{largest}"
	end

	i += 1
end

largest.chars.each do |i|
	largest_product = largest_product * i.to_i
end

puts "The largest product is: #{largest_product}"