-- Hexagonal numbers are also triangular. Thus we do not need to test for
-- triangular equality.

local h = 144
while true do
	h = h+1
	local hn = h*(2*h-1)
	local p = (1+math.sqrt(1 + 24 * hn))/6
	if p == math.floor(p) then
		print(p, hn)
		break
	end
end
