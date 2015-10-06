--[[ We do not need to generate the irrational fraction: we can query the digit
d_n directly by noting that:

* If 1 <= n <= 9, d_n = n.
* If 10 <= n <= 99, there are twice more digits. We fetch the generating integer that is around the n-th digit.

	offset := (n-9)/2 - floor((n-9)/2)
	gen_number := 9 + floor((n-9)/2) + ceil(offset)
	d_n = (2*offset)-th digit, or last digit if offset is 0

Examples:

* d_20:
(20-9)/2 = 5 + 1/2
offset = 1/2
gen_number = 9 + 5 + ceil(1/2) = 15
d_20 = 1st digit of 15 = 1

* d_21
(21-9)/2 = 6
offset = 0
gen_number = 15
d_21 = last digit of 15 = 5

* d_200
(200 -1*9 - 2*90)/3 = 3 + 2/3
offset = 2/3
gen_number = 99 + 3 = 102
d_200 = 2nd digit of 103 = 0
--]]--

local function digiter(n)
	local order=1
	local power=1

	-- Query the generating number.
	while n > 0 do
		n = n - power*9*order
		power = power+1
		order = order*10
	end
	-- We want the value before 0.
	power = power-1
	order = order/10
	n = n + power*9*order
	local gen_number = math.floor(order-1 + n/power)

	-- Find d_n.
	local offset = n % power
	if  offset == 0 then
		return gen_number % 10
	else
		gen_number = gen_number + 1
		for i = 1, offset-1 do
			order = order/10
		end
		return math.floor(gen_number/order) % 10
	end
end

local input = 1
local result = 1
local limit = 7
for i = 1, limit do
	result = result * digiter(input)
	input = input*10
end

print(result)
