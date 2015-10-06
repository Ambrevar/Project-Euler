--[[ We loop over the composite numbers to make sure the first hit is the minimum.
For every composite, two solutions:

* We can loop over all inferior primes and make sure that sqrt((k-p)/2) is an integer.

* Or we can loop over all inferior double squares and make sure k-ds is prime.

The double-square loop is faster for several reasons:

* sqrt() is an expensive operation.

* Since we need the sieve in any case, primality check comes for free.

* double squares grow much faster than prime numbers.

--]]

local limit = 100

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

local sieve = make_sieve(limit)
local sqrt2 = math.sqrt(2)
local double_squares = {}
for i = 1, limit/sqrt2 do
	double_squares[i] = 2*i*i
end
local start = 3

while true do
	for k = start, limit*limit-1, 2 do
		if sieve[k] then
			-- k is a composite number.
			local goldbach = false
			for _, ds in ipairs(double_squares) do
				if ds >= k then
					break
				end
				if not sieve[k-ds] then
					goldbach = true
					break
				end
			end
			if not goldbach then
				print(k)
				os.exit()
			end
		end
	end

	-- Sieve was not big enough: enlarge it and start again.
	start = limit*limit+1
	limit = 2*limit
	sieve = make_sieve(limit)
	for i = #double_squares+1, limit/sqrt2 do
		double_squares[i] = 2*i*i
	end
end
