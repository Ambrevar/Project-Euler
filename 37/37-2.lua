--[[ We generate numbers from the *_digitset variables. Since there are
restrictions on what the leftmost/inner/rightmost digits can be, this allows
us to restrict the input set.

This is slightly faster than 37-1.
--]]

-- We assume that no prime numbers will be above 1000000.
local limit = 1000
local sieve

-- Digit sets.
local inner_digitset = {1, 3, 7, 9}
local left_digitset = {2, 3, 5, 7}
local right_digitset = {3, 7}


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

local function truncatable(t, inner_digits, left_digit, right_digit)
	-- Left-right check
	local result=left_digit
	for _, i in ipairs (t) do
		result = result*10 + inner_digits[i]
		if sieve[result] then
			return nil
		end
	end

	result=result*10 + right_digit
	if sieve[result] then
		return nil
	end

	-- Right-left check
	result=right_digit
	local order = 10
	for i = #t, 1, -1 do
		result = result + order * inner_digits[t[i]]
		order = 10 * order
		if sieve[result] then
			return nil
		end
	end

	result = result + order * left_digit
	if sieve[result] then
		return nil
	end
	return result
end

sieve	= make_sieve(limit)

-- Loop here over innersize until 11 numbers are found.
local inner_count = 0
local count = 0
local count_max = 11
local sum = 0

while true do
	local inner = {}

	-- Left
	local left = 1
	while left <= #left_digitset do
		-- Right
		local right = 1
		while right <= #right_digitset do

			-- Reset 'inner'.
			for i = 1, inner_count do
				inner[i] = 1
			end
			local inner_index
			while count < count_max do
				inner_index = inner_count

				local number = truncatable(inner, inner_digitset, left_digitset[left], right_digitset[right])
				if number ~= nil then
					count = count + 1
					sum = sum + number
				end

				if inner_index >=1 then
					while inner[inner_index] == #inner_digitset do
						inner[inner_index] = 1
						inner_index = inner_index - 1
					end

					if inner_index >=1 then
						inner[inner_index] = inner[inner_index] + 1
					end
				end
				if inner_index == 0 then
					break
				end
			end

			if count >= count_max then
				print(sum)
				os.exit()
			end
			
			-- Right
			right = right + 1
		end

		-- Left
		left = left + 1
	end
	inner_count = inner_count + 1
end
