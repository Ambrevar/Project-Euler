--[[ We check prime numbers in decreasing order if they are pandigital. use
decreasing pandigital permutations as input set.
9-digits numbers cannot work since 1+2+...+9 = 45, so it is dividable by 3.
Same thing for 8-digits numbers.
We loop from 7-digit primes down.

The sieve is too big, thus this is much too slow.
--]]--

local limit = math.ceil(math.sqrt(9999999))

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

local function pandigital(n)
	local hit = {}
	while n > 0 do
		local m = n % 10
		if hit[m] then
			return false
		else
			hit[m] = true
			n = (n-m)/10
		end
	end
	for i = 1, 9 do
		if not hit[i] then
			for j = i+1, 9 do
				if hit[j] then
					return false
				end
			end
		end
	end
	return true
end

local	_, primes = make_sieve(limit)
for p = #primes, 1, -1 do
	if pandigital(primes[p]) then
		print(primes[p])
		return
	end
end
