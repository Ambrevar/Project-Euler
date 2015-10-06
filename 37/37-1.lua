-- Algorithm: We check for each prime if it is truncatable in both directions.

-- We assume that no prime numbers will be above 1000000.
local limit = 1000
local sieve, primes

-- See problem 3.
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
	local primes = {}
	sieve[0] = true
	if limit == 0 then
		return sieve, primes
	end
	sieve[1] = true
	for i = 2, limit do
		if not sieve[i] then
			primes[#primes+1] = i
			for j = i*i, bound-1, i do
				sieve[j] = true
			end
		end
	end

	-- If we would only care about the first limit-th values, this loop would not
	-- be necessary.
	for i = limit+1, bound-1 do
		if not sieve[i] then
			primes[#primes+1] = i
		end
	end

	return sieve, primes
end


local function truncatable(n)
	local left = n
	local right = 0
	local order = 1
	while left > 0 do
		if sieve[left] then
			return false
		end
		local m = left%10
		left = (left - m)/10
		right = right + order * m
		if sieve[right] then
			return false
		end
		order = order * 10
	end
	return true
end


sieve, primes = make_sieve(limit)

-- Loop here over innersize until 11 numbers are found.
local count = 0
-- We remove 2,3,5,7 from result.
local count_max=15
local sum=-17

for _, v in ipairs(primes) do
	if truncatable(v) then
		count = count + 1
		sum = sum + v
		if count == count_max then
			break
		end
	end
end

print(sum)
