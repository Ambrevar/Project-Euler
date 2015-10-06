--[[
We compute the pentogonal numbers in ascending order.
Thus 	we can use an array to check if smaller numbers are pentagonal.

The performance difference between array checking and direct computation is
shallow.

The preimage 'n' of a pentagonal number is (1+sqrt(24*(p[n]+p[i])+1))/6.

For every pentagonal 'p[n]' we loop i down from n-1 to 1 to check for p[n]-p[i]
and p[n]+p[i]. Looping down makes sure that if we find a distance, it will be
the shortest for 'n'. Thus we can early out.

The first result is *not necessarily* the right one, as opposed to what many
claim on the forum. Indeed, there might be n'>n and j>i such that p[n']-p[j] <
	p[n]-p[i] and such that (n',j) satisfies the condition. For instance,

	p[10]-p[9] == 28 < p[9]-p[7] == 47.

So we keep looping over n until the distance p[n]-p[n-1] is strictly more than
the last result.
--]]

local hit = {true}
local p = {1}
local n = 2
local result

-- Find a first solution.
while not result do
	p[n] = n*(3*n-1)/2
	hit[p[n]] = true

	for i = n-1, 1, -1 do
		if hit[p[n]-p[i]] then
			preimage = (1+math.sqrt(24*(p[n]+p[i])+1))/6
			if math.ceil(preimage) == preimage then
				result = p[n]-p[i]
				break
			end
		end
	end

	n=n+1
end

-- Find the optimal solution.
while true do
	p[n] = n*(3*n-1)/2
	hit[p[n]] = true

	for i = n-1, 1, -1 do
		if p[n]-p[i] < result then
			if hit[p[n]-p[i]] then
				preimage = (1+math.sqrt(24*(p[n]+p[i])+1))/6
				if math.ceil(preimage) == preimage then
					result = p[n]-p[i]
					break
				end
			end
		else
			if i == n-i then
				print(result)
				os.exit()
			end
			break
		end
	end

	n=n+1
end
