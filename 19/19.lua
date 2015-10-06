--[[
Nothing in particular here.
We encode monday as 1, saturday as 6, sunday as 0.
--]]

months = {
	function () return 31 end,
	function (year)
		if year % 400 == 0
			or (year % 4 == 0 and year % 100 ~= 0) then
			return 29
		else
			return 28
		end
	end,

	function () return 31 end,
	function () return 30 end,
	function () return 31 end,
	function () return 30 end,
	function () return 31 end,
	function () return 31 end,
	function () return 30 end,
	function () return 31 end,
	function () return 30 end,
	function () return 31 end,
}

-- Year 1900 has 365 days (not a leap year).
-- January 1st 1900 is a monday.
-- So January 1st 1901 is a:
day = ( 1 + 365 ) % 7
sunday_count = 0
for year = 1901, 2000 do
	for month = 1, 12 do
		if (day == 0) then
			sunday_count = sunday_count + 1
		end
		day = ( day + months[month](year) ) % 7
	end
end
	
print (sunday_count)
