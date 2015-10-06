-- This is the dynamic version of 31-2: it does not depend on the coin types.
-- It is slower due to bookkeeping.

local values = {200, 100, 50, 20 ,10, 5, 2, 1}

function decomp_count(n)
	local count = 0
	local t = {n}
	local coin = 1
	while t[1] >= 0 do
		if coin < #values then
			coin = coin + 1
			t[coin] = n
		else
			if n >= values[coin] then
				n = n - values[coin]
			else
				count = count + 1
				coin = coin - 1
				n = t[coin]
				while n < values[coin] and coin > 1 do
					coin = coin - 1
					n = t[coin]
				end
				if coin < 1 then
					return count
				end
				n = n - values[coin]
				t[coin] = n
			end
			
		end
	end

	return count
end		

print(decomp_count(200))
