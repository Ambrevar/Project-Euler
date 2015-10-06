-- We use decreasing pandigital permutations as input set.
-- 9-digits numbers cannot work since 1+2+...+9 = 45, so it is dividable by 3.
-- Same thing for 8-digits numbers.
-- We loop from 7-digit numbers to 1-digit numbers.

local factorial = {1, 2, 6, 24, 120, 720, 5040, 40320, 362880}

-- See problem 3.
local function isprime(n)
	if n == 2 or n == 3 or n == 5 then
		return true
	end

	if n <= 1 or n%2 == 0 or n%3 == 0 or n%5 == 0 then
		return false
	end

	local i = 7
	local step = 4
	while i*i <= n do
		if n%i == 0 then
			return false
		end
		i = i+step
		step = 6-step
	end
	return true
end

-- n: n-th permutation, d: number of digits. We divide the permutation in slides
-- of (d-1)!. This is slower than the recursive permutation generator, but here
-- we access increasing values.
local function permutate(n, d)
	if n <= 0 or d <= 0 or d >= 10 or n > factorial[d] then
		return 0
	end

	local result={}
	local source={1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i = d, 2, -1 do
		-- There are 'i' slices os size (i-1)! in 'i!'.
		local index = math.ceil(i * n /factorial[i])
		result[d+1-i] = source[index]
		table.remove(source, index)

		-- Remove the quantity in n that made it belong to the i-th slice.
		n = n - (index-1)*factorial[i-1]
	end
	result[d] = source[n]

	local sum = 0
	for i = 1, d do
		sum = sum*10 + result[i]
	end
	return sum
end

for d = 7, 1, -1 do
	for i = factorial[d], 1, -1 do
		local result = permutate(i, d)
		if isprime(result) then
			print(result)
			return
		end
	end
end
