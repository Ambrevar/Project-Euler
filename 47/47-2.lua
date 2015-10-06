--[[ Embed problem in sieve. Instead of storing booleans in the 'sieve' table, we
store the number of divisors.

Since we do not record divisors for every composite n, we can only be sure we
have found the right divisor count when i > n/2. That does not help much.

Recording divisors costs too much. It is slower to try to loop dynamically than
just reiterating with a higher limit. That might not always be true depending on
the count of numbers we are looking for.
--]]

-- We assume result is less than limit*limit.
local limit = 1000

local bound = limit*limit-1
local sieve = {}

for i = 2, limit do
	if not sieve[i] then
		-- i is prime.

		-- WARNING: Since we want the number of divisors and not only the primality,
		-- we need to start at j=2*i, and not j=i*i so that we set include all
		-- numbers between 2*i and i*i as multiple of i.
		for j = 2*i, bound, i do
			if not sieve[j] then sieve[j] = 1 else sieve[j] = sieve[j] + 1 end
		end
	end
end

for i = 1, bound do
	if sieve[i] == 4 and sieve[i+1] == 4 and sieve[i+2] == 4 and sieve[i+3] == 4 then
		print(i)
		break
	end
end
