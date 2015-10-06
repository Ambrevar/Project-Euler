-- Only check even values of perimeter, since a right triangle perimeter cannot
-- be odd.

-- Note: Much faster with Luajit.

local function decomp(n)
	local half = math.floor((n-1)/2)
	local count = 0
	for i = 1, half do
		for j = math.floor((n-i)/2), n-i-1 do
			local k = n-i-j
			if i*i == j*j + k*k then
				count = count+1
			end
		end
	end
	return count
end

local result = 0
local count = 0
-- The smallest perimeter with integral lengths is 12.
for i = 12, 1000, 2 do
	local d = decomp(i)
	if d > count then
		result = i
		count = d
	end
end

print(result)
