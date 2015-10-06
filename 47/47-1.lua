--[[ Start at 2*3*5*7 = 210. We loop over composites. We can skip the numbers
for which the 3 next numbers are not all composites.

This is very slow.
--]]

-- We assume result is less than limit*limit.
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

local function divcount(n, primes)
	local primecount = 0

	for _, p in pairs(primes) do
		if n % p == 0 then
			primecount = primecount+1
			while n%p == 0 do
				n = n/p
			end
		end
		if n == 1 then break end
	end
	return primecount
end

local sieve, primes = make_sieve(limit)

for n=210, limit*limit-1 do
	if sieve[n] and sieve[n+1] and sieve[n+2] and sieve[n+3] then
		if divcount(n, primes) == 4 then
			n = n+1
			if divcount(n, primes) == 4 then
				n = n+1
				if divcount(n, primes) == 4 then
					n = n+1
					if divcount(n, primes) == 4 then
						print(n-3)
						break
					end
				end
			end
		end
	end
end
