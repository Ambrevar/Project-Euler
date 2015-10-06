-- Recursive version: short, dynamic but very slow.

local values = {200, 100, 50, 20 ,10, 5, 2, 1}
local bound = 200

function decomp_count(n, m)
	local max = m or n
	if n == 0 then return 1 end
	local result = 0
	for i = 1, #values do
		local v=values[i]
		if v <= max then
			local t = decomp_count(n-v, math.min(v,n-v))
			result = result + t
		end
	end
	return result
end

print(decomp_count(bound))
