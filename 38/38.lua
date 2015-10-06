--[[ Analysis:
We write the solution (k, n) where concat(k * (1,2,...,n)) is 9-pandigital.

(k,n)=(1,9) is solution.
So n<=9.

Better: 9*(1,2,3,4,5) = 918273645,
so if k>9, n<5.
k should begin with '9'.

Also k<=9876.

The is no 0, so k cannot be a multiple of 5.
Indeed, if k%5 == 0, then k*2 % 5 = 0.

Using the same reasoning, if n>=5, then k cannot be even.

We write that k has D digits.
k*2, ..., k*9 have D or D+1 digits.
So 9=D*X + (D+1)*Y
If n=2, k has 4 digits and 5000 <= k <= 9876.
If n=3, k has 3 digits and 100 <= k <= 333.
If n=4, k has 2 digits and 25 <= k <= 33.
--]]

local function pandigital(k, n)
	local tab = {0, 0, 0, 0, 0, 0, 0, 0, 0}
	local result = 0
	local power = 1

	for i=n, 1, -1 do
		local product = i * k
		result = result + power * product

		while product ~= 0 do
			local index = product % 10
			if tab[index] ~= 0 then
				return 0
			else
				tab[index] = 1
			end
			product = math.floor(product/10)
			power = 10 * power
		end

	end
	return result
end

-- Bounds
local result = 0

-- n=2.
for k = 9000, 9876 do
	if k%5 ~= 0 then
		local p = pandigital(k, 2)
		if p ~= 0  and p > result then
			result = p
		end
	end
end

print(result)
