--[[ We generate 4-digit primes.
For each of these prime, generate the permutations that are 4 digits and primes.
Remove these primes from the outer loop.
Sort the primes so that we can quickly find an arithmetic sequence.
--]]

local function make_sieve(limit)
	local bound
	if limit > 1 then
		bound = limit * limit
	elseif limit == 1 then
		bound = 2
	else
		bound = 1
	end

	local sieve = {}
	sieve[0] = true
	if limit == 0 then
		return sieve
	end
	sieve[1] = true
	for i = 2, limit do
		if not sieve[i] then
			for j = i*i, bound-1, i do
				sieve[j] = true
			end
		end
	end

	return sieve
end

local sieve = make_sieve(100)

local function permute(array, size, result)
	if size == 1 then
		local n = 0
		for _, v in ipairs(array) do
			n = n * 10 + v
		end
		if n > 1000 and not sieve[n] then
			-- We store result as a key to ignore duplicates.
			result[n] = true
		end
		return
	end

	for k = 1, size do
		array[k], array[size] = array[size], array[k]
		permutate(array, size-1, result)
		array[k], array[size] = array[size], array[k]
	end
end

local function permutations(n)
	-- Store the digits of n in a table.
	local digits = {}
	while n > 0 do
		local digit = n%10
		digits[#digits+1] = digit
		n = (n - digit) / 10
	end
	
	local p = {}
	permute(digits, #digits, p)

	return p
end


local function distance(array)
	-- Store values as indices.
	local hit = {}
	for _, p in pairs(array) do
		hit[p] = true
	end

	-- Sort array so that we can loop over growing values.
	table.sort(array)
	for middle = 2, #array-1 do
		for low = 1, middle-1 do
			local high = array[middle] + (array[middle] - array[low])
			if hit[high] then
				return array[low], array[middle], high
			end
		end
	end
	return nil
end

for i = 1000, 9999 do
	if not sieve[i] then
		-- i is prime.
		local p = permutations(i)

		local prime_set = {}
		for k, _ in pairs(p) do
		-- Remove prime from sieve so that we don't parse them in future loops.
			sieve[k] = true
			prime_set[#prime_set+1] = k
		end

		if #prime_set >= 3 then
			local low, middle, high = distance(prime_set)
			if low and low ~= 1487 then
				print(low .. middle .. high)
				break
			end
		end
	end
end

