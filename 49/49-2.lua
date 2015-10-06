--[[ Same as 49-1 but store prime permutations together in ascending order. To
group permutations together we use a hash map with the increasing digits as key.
Performance is similar.
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

local permutations = {}

for p = 1000, 9999 do
	if not sieve[p] then
		-- p is prime.

		local digits = {0, 0, 0, 0, 0, 0, 0, 0, 0}
		digits[0] = 0
		local prime = p
		while p > 0 do
			local d = p % 10
			digits[d] = digits[d] + 1
			p = (p-d) / 10
		end
		p = prime
		local key = table.concat(digits, '', 0)

		if not permutations[key] then
			permutations[key] = {p}
		else
			permutations[key][#permutations[key]+1] = p
		end

		-- Test if we found a sequence other than the one where the higher number is 8147.
		if #permutations[key] >= 3 and permutations[key][#permutations[key]] ~= 8147 then
			-- Permutations are already sorted.
			-- Use last prime in the table since we just added it.
			for i = #permutations[key]-1, 2, -1 do
				for j = i-1, 1, -1 do
					if permutations[key][#permutations[key]] - permutations[key][i] == permutations[key][i] - permutations[key][j] then
						print(permutations[key][j] .. permutations[key][i] .. permutations[key][#permutations[key]])
					end
				end
			end
		end
	end
end
