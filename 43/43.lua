-- Generate all 3-digit permutations of {0, ..., 9} that are divisible by 17.
-- Prepend an unused digit so that the result is divisible by 13.
-- Prepend an unused digit so that the result is divisible by 11.
-- And so on until 2.
-- Sum the result.

-- WARNING: mask starts at index 0. Value is an array so that we don't remove
-- the leading 0.

-- Structure example: {
-- value = {2, 8, 9}
-- mask = {false, false, true, false, false, false, false, true, true}
--}

local function permutate(pandigits, array, size, digit_count)
	if size == #array - digit_count then
		local v = array[#array-2]*100 + array[#array-1]*10 + array[#array]
		if v % 17 == 0 then
			pandigits[#pandigits+1] = {hit = {}, value = {array[#array-2], array[#array-1], array[#array]}}
				
			pandigits[#pandigits].hit[array[#array-2]] = true
			pandigits[#pandigits].hit[array[#array-1]] = true
			pandigits[#pandigits].hit[array[#array]] = true
		end
		return
	else
		for i = 1, size do
			array[size], array[i] = array[i], array[size]
			permutate(pandigits, array, size-1, digit_count)
			array[size], array[i] = array[i], array[size]
		end
	end
end

local function tabtonum(t)
	local i = 1
	local result = 0
	while t[i] ~= nil do
		result = result*10 + t[i]
		i = i+1
	end
	return result
end

local t = {1,2,3,4,5,6,7,8,9,0}
local pandigits = {}
permutate(pandigits, t, #t, 3)

local primes = {13, 11, 7, 5, 3, 2}
for _, prime in ipairs(primes) do
	-- Store higher order pandigital numbers.
	pandigits_up = {}

	for _, pandigital in ipairs(pandigits) do
		for digit = 0, 9 do

			if not pandigital.hit[digit] then
				local v = 100*digit + 10*pandigital.value[1] + pandigital.value[2]

				if v % prime == 0 then
					pandigits_up[#pandigits_up+1] = {hit={}, value={}}

					-- Copy old arrays.
					for d = 0, 9 do
						pandigits_up[#pandigits_up].hit[d] = pandigital.hit[d]
					end
					for i, d in ipairs(pandigital.value) do
						pandigits_up[#pandigits_up].value[i+1] = d
					end

					-- Update.
					pandigits_up[#pandigits_up].hit[digit] = true
					pandigits_up[#pandigits_up].value[1] = digit
				end
			end
		end
	end

	-- Forget current order and only keep higher order.
	pandigits = pandigits_up
end

local result = 0
for _, pandigital in pairs(pandigits) do

	-- Complete numbers with last digits put in front.
	local j = 0
	while pandigital.hit[j] do
		j = j+1
	end
	if j ~= 0 then
		result = result + j*1000000000 + tabtonum(pandigital.value)
	end
end

print(result)
