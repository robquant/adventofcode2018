import strutils, sequtils, math, tables

let
  strNums = readFile("input.txt").strip().splitLines()  
  nums = strNums.map(parseInt)       

let
  sumNums = sum(nums)
  
echo sumNums


var freqs = initTable[int, bool]()
var freq = 0
var found = false
while true:
  for f in nums:
    freq += f
    if freqs.hasKey(freq):
      echo freq
      found = true
      break
    freqs[freq] = true
  if found:
    break