--[[ We loop over the decreasing primes. As soon as the sum is over the sieve
size we can early out since further sums will be even higher.
--]]

local limit = 1000

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

local sieve, primes = make_sieve(limit)
local bound = limit*limit-1

-- We look for sequences.
local start = #primes
local length = 2
local result = 0

while start >= length do
	local sum = 0

	for i = start, 1, -1 do
		sum = sum + primes[i]
		if sum > bound then
			-- Early out.
			break
		end
		if not sieve[sum] and start-i > length then
			length = start-i
			result = sum
		end
	end

	start = start - 1
end

print(result)
