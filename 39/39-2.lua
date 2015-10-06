--[[ Other way around: We generate all triangles and fill a table with count.
Pretty much like the sieve.

We name the hypotenuse h.
We loop over the long side l and the short s.
l cannot be more than (limit-1)/2 since h > l.

s <= l by definition.
But when l is high enough that

  l + l + sqrt(l^2 + l^2) >= limit

i.e.

  l >= limit/(2+sqrt(2))

then s can be further limited.

  s + l + sqrt(s^2 + l^2) <= limit
  s^2 <= limit * (limit - 2*l) / 2 / (limit - l)

We could split the loop in 2 at limit/(2+sqrt(2)), but this does not save cycles.
--]]


local limit = 1000
local perimeters = {}

for long = 1, math.floor((limit-1)/2) do
	local long2 = long*long
	local short_max = limit*(limit-2*long)/2/(limit-long)
	if short_max > long then
		short_max = long
	end

	for short = 1, short_max do
		local h = math.sqrt(long2 + short*short)
		if h == math.floor(h) then
			p = h + long + short
			if perimeters[p] then
				perimeters[p] = perimeters[p] + 1
			else
				perimeters[p] = 1
			end
		end
	end
end

local result = 0
local count = 0
for k, v in pairs(perimeters) do
	if v > count then
		count = v
		result = k
	end
end

print(result)
