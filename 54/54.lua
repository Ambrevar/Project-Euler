--[[ No fancy algorithm, this is mostly about structuring and sorting.
Hand structure: {{card, quantity}...}.
Example: {{K,2}, {Q,2}, {A,1}}.
Hand sorting function sort by quantity first between {c1, q1} and {c2, q2}.
--]]

local card_values = {}
do
	local card_values_order = {'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}
	for i, j in ipairs(card_values_order) do
		card_values[j] = i
	end
end

local ranks_order = {'High', 'One', 'Two', 'Three', 'Straight', 'Flush', 'FullHouse', 'Four', 'StraightFlush', 'RoyalFlush'}
local ranks = {}
for i, j in ipairs(ranks_order) do
	ranks[j] = i
end

local function hand_sort(a, b)
	if a[2] > b[2] then
		return true
	elseif a[2] < b[2] then
		return false
	else
		if card_values[a[1]] > card_values[b[1]] then
			return true
		else
			return false
		end
	end
end

-- Input is a hand string from the text file.
local function check_hand (input)
	-- Create 'hand' structure from 'input' string.
	local hand = {}
	local handgen = {}
	for i = 0, 4 do
		local index = string.sub(input, i*3+1, i*3+1)
		if handgen[index] == nil then
			handgen[index] = 1
		else
			handgen[index] = handgen[index] + 1
		end
	end

	local index = 1
	for i, j in pairs(handgen) do
		hand[index] = {i, j}
		index = index + 1
	end

	table.sort(hand, hand_sort)
	
	-- Check if all cards are from the same suit.
	local is_same_suit = true
	local current_suit = string.sub(input, 2, 2)
	for i = 1, 4 do
		if current_suit ~=  string.sub(input, i*3+2, i*3+2) then
			is_same_suit = false
			break
		end
	end

	-- Straight.
	local is_straight = false
	if #hand == 5 then
		is_straight = true
		for i = 2, 5 do
			if card_values[hand[i][1]] ~= card_values[hand[i-1][1]]-1 then
				is_straight = false
				break
			end
		end
	end

	local rank

	-- Check if Flush, Straight Flush or Royal Flush.
	if is_same_suit then
		if is_straight then
			if hand[1][1] == 'A' then
				rank = 'RoyalFlush'
			else
				rank = 'StraightFlush'
			end
		else
			-- Flush > Straight
			rank = 'Flush'
		end

		-- All but straight and Flush-like.
	elseif not is_straight then
		if hand[1][2] == 4 then
			rank = 'Four'
		elseif hand[1][2] == 3 and hand[2][2] == 2 then
			rank = 'FullHouse'
		elseif hand[1][2] == 3 then
			rank = 'Three'
		elseif hand[1][2] == 2 and hand[2][2] == 2 then
			rank = 'Two'
		elseif hand[1][2] == 2 then
			rank = 'One'
		else
			rank='High'
		end
	else
		rank = 'Straight'
	end

	return hand, rank
end

-- True if input1 > input2.
local function compare(input1, input2)
	local hand1, rank1 = check_hand(input1)
	local hand2, rank2 = check_hand(input2)

	if ranks[rank1] > ranks[rank2] then
		return true
	elseif ranks[rank1] < ranks[rank2] then
		return false
	else
		-- Same rank
		for i = 1, #hand1 do
			if card_values[hand1[i][1]] > card_values[hand2[i][1]] then
				return true
			elseif card_values[hand1[i][1]] < card_values[hand2[i][1]] then
				return false
			end
		end
	end
end

-- Load file.
local fp = assert(io.open('p054_poker.txt', "r"))
local buffer = fp:read("*all")
assert(fp:close())

local player1_score = 0

for line in buffer:gmatch("[^\n]+\n?") do
	local _, _, input1, input2 = string.find(line, "(..............) (..............)\n?")

	if compare(input1, input2) then
		player1_score = player1_score +1
	end
end

print(player1_score)
