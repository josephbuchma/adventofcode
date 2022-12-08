

const priority = (letter: string) => {
	const [a, z, A] = ["a", "z", "A"].map(c => c.charCodeAt(0))
	const code = letter.charCodeAt(0)
	if (code >= a && code <= z) {
		return code - a + 1
	}
	return  code - A + 27
}

const parseLine = (str: string) => str.split("").map(priority)

const splitAndSortCompartments = (items: number[]) => 
	[items.slice(0, items.length/2).sort(), items.slice(items.length/2).sort()] as const

const findRepeatedItems = ([c1, c2]: readonly [number[], number[]]) => {
	const found =  c1.reduce((acc, c1i) => {
		for (const c2i of c2) {
			if (c2i === c1i) {
				acc.add(c2i)
				break
			}
		}
		return acc;
	}, new Set<number>())
	return [...found.values()]
}

const sumNums = (nums: number[]) => nums.reduce((a, v) => a + v, 0)

const findSumOfRepeatedPriorities = (input: string) =>
	sumNums(input.split("\n").map(parseLine).map(splitAndSortCompartments).map(findRepeatedItems).map(sumNums))

const sampleInput = `
vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`.trim()

console.log("Sample input result: ", findSumOfRepeatedPriorities(sampleInput))

const realInput = await Deno.readTextFile("./input-part-1.txt");
console.log("Real input result: ", findSumOfRepeatedPriorities(realInput))

/* Part two */

const forEachChunk = <T>(chunkSize: number, arr: Array<T>, cb: (a: T[])=>void) => {
	for (let i=0; i< arr.length; i+=chunkSize) {
		cb(arr.slice(i, Math.min(i+chunkSize, arr.length)))
	}
}

const findSharedNumInSortedArrays = (arrays: Array<number[]>) => {
	let max = arrays[0][0];
	while (true) {
		let noMatch = false;
		let allSame = true;
		arrays = arrays.map(arr => {
			allSame = allSame && arr[0] === max;
			if (arr[0] === max) {
				return arr;
			}
			if (arr[0] > max) {
				max = arr[0]
				return arr;
			}
			if (arr.length === 1) {
				noMatch = true;
			}
			return arr.slice(1);
		})
		if (allSame) {
			return max;
		}
		if (noMatch) {
			return null;
		}
	}
}

const findSharedItems = (input: number[][]) => {
	let sum = 0;
	forEachChunk(3, input, (group) => {
		const badge = findSharedNumInSortedArrays(group);
		if (!badge) {
			throw new Error("no badge found")
		}
		if (badge) {
			sum += badge;
		}
	})
	return sum;
}

const realInputPt2 = await Deno.readTextFile('./input-part-2.txt');
const pt2Result =  findSharedItems(
	realInputPt2.split("\n")
		.map(parseLine)
		.map(l => l.sort((a, b) => a - b))
	);

console.log("Part 2 result:", pt2Result);